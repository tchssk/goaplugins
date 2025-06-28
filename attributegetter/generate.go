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
		BaseType string
		Name     string
		Type     string
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
		appendSections("service-payload-method", f, svc, section, true)
	}
	for _, section := range f.Section("service-result") {
		appendSections("service-result-method", f, svc, section, false)
	}
	for _, section := range f.Section("service-user-type") {
		data, ok := section.Data.(*service.UserTypeData)
		if !ok {
			return
		}
		if err := codegen.WalkMappedAttr(expr.NewMappedAttributeExpr(data.Type.Attribute()), func(name, elem string, required bool, a *expr.AttributeExpr) error {
			if !mustGenerate(a.Meta) {
				return nil
			}
			f.SectionTemplates = append(f.SectionTemplates, &codegen.SectionTemplate{
				Name:   "service-user-type-method",
				Source: methodT,
				Data:   buildMethodData(data.Type.Attribute(), name, a, data.Name),
			})
			return nil
		}); err != nil {
			return
		}
	}
}

func appendSections(sectionName string, f *codegen.File, svc *expr.ServiceExpr, section *codegen.SectionTemplate, isPayload bool) {
	data, ok := section.Data.(*service.MethodData)
	if !ok {
		return
	}
	method := svc.Method(data.Name)
	if method == nil {
		return
	}
	var att *expr.AttributeExpr
	if isPayload {
		att = method.Payload
	} else {
		att = method.Result
	}
	if err := codegen.WalkMappedAttr(expr.NewMappedAttributeExpr(att), func(name, elem string, required bool, a *expr.AttributeExpr) error {
		if !mustGenerate(a.Meta) {
			return nil
		}
		var (
			parent   *expr.AttributeExpr
			baseType string
		)
		if isPayload {
			parent = method.Payload
			baseType = data.Payload
		} else {
			parent = method.Result
			baseType = data.Result
		}
		f.SectionTemplates = append(f.SectionTemplates, &codegen.SectionTemplate{
			Name:   sectionName,
			Source: methodT,
			Data:   buildMethodData(parent, name, a, baseType),
		})
		return nil
	}); err != nil {
		return
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

func buildMethodData(parent *expr.AttributeExpr, name string, att *expr.AttributeExpr, baseType string) *MethodData {
	scope := codegen.NewNameScope()
	typ := scope.GoTypeName(att)
	if parent.IsPrimitivePointer(name, true) || !expr.IsPrimitive(att.Type) {
		typ = "*" + typ
	}
	return &MethodData{
		BaseType: baseType,
		Name:     codegen.GoifyAtt(att, name, true),
		Type:     typ,
	}
}

var methodT = `
func (p *{{ .BaseType }}) Get{{ .Name }}() {{ .Type }} {
	return p.{{ .Name }}
}
`
