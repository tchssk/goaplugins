package optionalbody

import (
	"path/filepath"
	"strings"

	"github.com/tchssk/goaplugins/v3/optionalbody/expr"
	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/service"
	"goa.design/goa/v3/eval"
	httpcodegen "goa.design/goa/v3/http/codegen"
)

const (
	serverPayloadInitNameSuffix        = "WithOptionalBody"
	serverPayloadInitDescriptionSuffix = " It allows an empty body."
)

const payloadInitT = `{{ printf "New%s initializes payload type %s." .Payload .Payload | comment }}
func New{{ .Payload }}(emptyBody bool) *{{ .Payload }} {
	return &{{ .Payload }}{
		emptyBody: emptyBody,
	}
}
`

const payloadHasEmptyBodyT = `{{ comment "HasEmptyBody reports whether the payload has an empty body." }}
func (p *{{ .Payload }}) HasEmptyBody() bool {
	return p.emptyBody
}
`

func init() {
	codegen.RegisterPlugin("optionalbody", "gen", nil, Update)
}

func Update(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	for _, f := range files {
		update(f)
	}
	return files, nil
}

func update(f *codegen.File) {
	switch filepath.Base(f.Path) {
	case "service.go":
		for _, section := range f.Section("service") {
			data, ok := section.Data.(*service.Data)
			if !ok {
				continue
			}
			s, ok := expr.Root.HTTPEndpoints[data.Name]
			if !ok {
				continue
			}
			for _, method := range data.Methods {
				if _, ok := s[method.Name]; !ok {
					continue
				}
				if strings.Contains(method.PayloadDef, "emptyBody bool") {
					continue
				}
				method.PayloadDef = addField(method.PayloadDef, "emptyBody bool")
				f.SectionTemplates = append(f.SectionTemplates, &codegen.SectionTemplate{
					Name:   "service-payload-init",
					Source: payloadInitT,
					Data:   method,
				})
				f.SectionTemplates = append(f.SectionTemplates, &codegen.SectionTemplate{
					Name:   "service-payload-hasemptybody",
					Source: payloadHasEmptyBodyT,
					Data:   method,
				})
			}
		}
	case "encode_decode.go":
		for _, section := range f.Section("request-decoder") {
			data, ok := section.Data.(*httpcodegen.EndpointData)
			if !ok {
				continue
			}
			s, ok := expr.Root.HTTPEndpoints[data.ServiceName]
			if !ok {
				continue
			}
			if _, ok := s[data.Method.Name]; !ok {
				continue
			}
			if !strings.HasSuffix(data.Payload.Request.PayloadInit.Description, serverPayloadInitDescriptionSuffix) {
				data.Payload.Request.PayloadInit.Description += serverPayloadInitDescriptionSuffix
			}
			section.Source = strings.Replace(section.Source,
				`		var (
			body {{ .Payload.Request.ServerBody.VarName }}
			err  error
		)`,
				`		var (
			body {{ .Payload.Request.ServerBody.VarName }}
			emptyBody bool
			err  error
		)`,
				-1,
			)
			section.Source = strings.Replace(section.Source,
				`	{{- if .Payload.Request.MustHaveBody }}
			if errors.Is(err, io.EOF) {
				return nil, goa.MissingPayloadError()
			}
	{{- else }}
			if errors.Is(err, io.EOF) {
				err = nil
			} else {
	{{- end }}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
	{{- if not .Payload.Request.MustHaveBody }}
			}
	{{- end }}`,
				`			if !errors.Is(err, io.EOF) {
				var gerr *goa.ServiceError
				if errors.As(err, &gerr) {
					return nil, gerr
				}
				return nil, goa.DecodePayloadError(err.Error())
			}
			emptyBody = true
			err = nil`,
				-1,
			)
			section.Source = strings.Replace(section.Source,
				`	{{- if .Payload.Request.ServerBody.ValidateRef }}
		{{ .Payload.Request.ServerBody.ValidateRef }}
		if err != nil {
			return nil, err
		}
	{{- end }}`,
				`	{{- if .Payload.Request.ServerBody.ValidateRef }}
		if !emptyBody {
			{{ .Payload.Request.ServerBody.ValidateRef }}
			if err != nil {
				return nil, err
			}
		}
	{{- end }}`,
				-1,
			)
			section.Source = strings.Replace(section.Source,
				`	payload := {{ .Payload.Request.PayloadInit.Name }}({{ range .Payload.Request.PayloadInit.ServerArgs }}{{ .Ref }}, {{ end }})`,
				`	payload := {{ .Payload.Request.PayloadInit.Name }}`+serverPayloadInitNameSuffix+`({{ range .Payload.Request.PayloadInit.ServerArgs }}{{ .Ref }}, {{ end }})
	if !emptyBody {
		payload = {{ .Payload.Request.PayloadInit.Name }}({{ range .Payload.Request.PayloadInit.ServerArgs }}{{ .Ref }}, {{ end }})
	}`,
				-1,
			)
		}
	case "types.go":
		for _, section := range f.Section("server-payload-init") {
			data, ok := section.Data.(*httpcodegen.InitData)
			if !ok {
				continue
			}
			if !strings.HasSuffix(data.Description, serverPayloadInitDescriptionSuffix) {
				continue
			}
			f.SectionTemplates = append(f.SectionTemplates, &codegen.SectionTemplate{
				Name:   "server-payload-init",
				Source: section.Source,
				Data: &httpcodegen.InitData{
					Name:                data.Name + serverPayloadInitNameSuffix,
					Description:         data.Description,
					ServerArgs:          data.ServerArgs,
					ClientArgs:          data.ClientArgs,
					CLIArgs:             data.CLIArgs,
					ReturnTypeName:      data.ReturnTypeName,
					ReturnTypeRef:       data.ReturnTypeRef,
					ReturnIsStruct:      data.ReturnIsStruct,
					ReturnTypeAttribute: data.ReturnTypeAttribute,
					ServerCode:          fixAssignments(data.ServerCode),
					ClientCode:          data.ClientCode,
				},
				FuncMap: section.FuncMap,
			})
		}
	}
}

func addField(s, field string) string {
	ss := strings.Split(s, "\n")
	if len(ss) < 3 {
		return s
	}
	return strings.Join(append(ss[:len(ss)-1], field, ss[len(ss)-1]), "\n")
}

func fixAssignments(s string) string {
	ss := strings.Split(s, "\n")
	if len(ss) < 3 {
		return s
	}
	return strings.Replace(strings.Replace(strings.Replace(ss[0], ":= &", ":= ", -1), ".", ".New", -1), "{", "(true)", -1)
}
