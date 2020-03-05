package optionalbody

import (
	"path/filepath"
	"strings"

	"github.com/tchssk/goaplugins/optionalbody/expr"
	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/service"
	"goa.design/goa/v3/eval"
	httpcodegen "goa.design/goa/v3/http/codegen"
)

const (
	serverPayloadInitDescriptionSuffix = " It allows an empty body."
)

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
				if strings.Contains(method.PayloadDef, "EmptyBody bool") {
					continue
				}
				method.PayloadDef = addField(method.PayloadDef, "EmptyBody bool")
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
				section.Source = strings.Replace(section.Source,
					`			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())`,
					`			if err != io.EOF {
				return nil, goa.DecodePayloadError(err.Error())
			}`,
					-1,
				)
			}
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
			if section.FuncMap == nil {
				section.FuncMap = map[string]interface{}{}
			}
			section.FuncMap["fixAssignments"] = fixAssignments
			section.Source = strings.Replace(section.Source,
				`	{{- if .ServerCode }}
		{{ .ServerCode }}`,
				`	{{- if .ServerCode }}
		var v {{ .ReturnTypeRef }}
		if body == nil {
			{{ fixAssignments .ServerCode }}
		} else {
			{{ .ServerCode }}
		}`,
				-1,
			)
			data.ServerCode = strings.Replace(data.ServerCode, "v :=", "v =", -1)
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
	return strings.Join([]string{ss[0], "EmptyBody: true,", ss[len(ss)-1]}, "\n")
}
