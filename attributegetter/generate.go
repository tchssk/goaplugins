package attributegetter

import (
	"path/filepath"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/service"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

type (
	MethodData struct {
		Payload string
		Name    string
		Type    string
		Var     string
	}
)

func init() {
	codegen.RegisterPlugin("attributegetter", "gen", nil, Generate)
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	for _, f := range files {
		serviceAttributeGetter(f)
	}
	return files, nil
}

func serviceAttributeGetter(f *codegen.File) {
	if filepath.Base(f.Path) != "service.go" {
		return
	}
	var svc *expr.ServiceExpr
	for _, section := range f.Section("service") {
		data, ok := section.Data.(*service.Data)
		if !ok {
			return
		}
		svc = expr.Root.Service(data.Name)
	}
	if svc == nil {
		return
	}
	for _, section := range f.Section("service-payload") {
		data, ok := section.Data.(*service.MethodData)
		if !ok {
			continue
		}
		method := svc.Method(data.Name)
		if method == nil {
			continue
		}
		if method.Payload.Type == expr.Empty {
			continue
		}
		dt, ok := method.Payload.Type.(expr.UserType)
		if !ok {
			continue
		}
		fm := codegen.TemplateFuncs()
		obj := expr.AsObject(dt)
		for _, nat := range *obj {
			if !mustGenerate(nat.Attribute.Meta) {
				continue
			}
			f.SectionTemplates = append(f.SectionTemplates, &codegen.SectionTemplate{
				Name:    "service-payload-method",
				Source:  servicePayloadMethodT,
				Data:    buildMethodData(method, nat, data.Payload, codegen.NewNameScope()),
				FuncMap: fm,
			})
		}
	}
}

func mustGenerate(meta expr.MetaExpr) bool {
	if m, ok := meta["attributegetter:generate"]; ok {
		if len(m) > 0 && m[0] == "false" {
			return false
		}
	}
	return true
}

func buildMethodData(method *expr.MethodExpr, nat *expr.NamedAttributeExpr, payload string, scope *codegen.NameScope) *MethodData {
	name := codegen.GoifyAtt(nat.Attribute, nat.Name, true)
	typ := scope.GoTypeName(nat.Attribute)
	if method.Payload.IsPrimitivePointer(nat.Name, true) {
		typ = "*" + typ
	}
	return &MethodData{
		Payload: payload,
		Name:    name,
		Type:    typ,
		Var:     name,
	}
}

var servicePayloadMethodT = `
func (p *{{ .Payload }}) Get{{ .Name }}() {{ .Type }} {
	return p.{{ .Var }}
}
`
