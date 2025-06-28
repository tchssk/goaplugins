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
		if data.Type == expr.Empty {
			return
		}
		userType := data.Type
		obj := expr.AsObject(userType)
		if obj == nil {
			return
		}
		for _, nat := range *obj {
			if !mustGenerate(nat.Attribute.Meta) {
				continue
			}
			f.SectionTemplates = append(f.SectionTemplates, &codegen.SectionTemplate{
				Name:   "service-user-type-method",
				Source: methodT,
				Data:   buildMethodData(userType.Attribute(), nat, data.Name),
			})
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
	userType, ok := getDataType(method, isPayload)
	if !ok {
		return
	}
	obj := expr.AsObject(userType)
	if obj == nil {
		return
	}
	for _, nat := range *obj {
		if !mustGenerate(nat.Attribute.Meta) {
			continue
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
			Data:   buildMethodData(parent, nat, baseType),
		})
	}
}

func getDataType(method *expr.MethodExpr, isPayload bool) (expr.UserType, bool) {
	if isPayload {
		if method.Payload.Type == expr.Empty {
			return nil, false
		}
		dt, ok := method.Payload.Type.(expr.UserType)
		return dt, ok
	} else {
		if method.Result.Type == expr.Empty {
			return nil, false
		}
		dt, ok := method.Result.Type.(expr.UserType)
		return dt, ok
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

func buildMethodData(parent *expr.AttributeExpr, nat *expr.NamedAttributeExpr, baseType string) *MethodData {
	scope := codegen.NewNameScope()
	typ := scope.GoTypeName(nat.Attribute)
	if parent.IsPrimitivePointer(nat.Name, true) || !expr.IsPrimitive(nat.Attribute.Type) {
		typ = "*" + typ
	}
	return &MethodData{
		BaseType: baseType,
		Name:     codegen.GoifyAtt(nat.Attribute, nat.Name, true),
		Type:     typ,
	}
}

var methodT = `
func (p *{{ .BaseType }}) Get{{ .Name }}() {{ .Type }} {
	return p.{{ .Name }}
}
`
