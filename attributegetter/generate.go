package attributegetter

import (
	"path/filepath"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/service"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

type (
	MethodPayloadData struct {
		Payload string
		Name    string
		Type    string
		Var     string
	}
	MethodResultData struct {
		Result string
		Name   string
		Type   string
		Var    string
	}
	MethodUserTypeData struct {
		UserType string
		Name     string
		Type     string
		Var      string
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
		appendSections(f, svc, section, true)
	}
	for _, section := range f.Section("service-result") {
		appendSections(f, svc, section, false)
	}
	for _, section := range f.Section("service-user-type") {
		data, ok := section.Data.(*service.UserTypeData)
		if !ok {
			return
		}
		if data.Type == expr.Empty {
			return
		}
		dt := data.Type
		fm := codegen.TemplateFuncs()
		obj := expr.AsObject(dt)
		if obj == nil {
			return
		}
		for _, nat := range *obj {
			if !mustGenerate(nat.Attribute.Meta) {
				continue
			}
			f.SectionTemplates = append(f.SectionTemplates, &codegen.SectionTemplate{
				Name:    "service-user-type-method",
				Source:  serviceUserTypeMethodT,
				Data:    buildMethodUserTypeData(dt.Attribute(), nat, dt.Name(), codegen.NewNameScope()),
				FuncMap: fm,
			})
		}
	}
}

func appendSections(f *codegen.File, svc *expr.ServiceExpr, section *codegen.SectionTemplate, isPayload bool) {
	data, ok := section.Data.(*service.MethodData)
	if !ok {
		return
	}
	method := svc.Method(data.Name)
	if method == nil {
		return
	}
	dt, ok := getDataType(method, isPayload)
	if !ok {
		return
	}
	fm := codegen.TemplateFuncs()
	obj := expr.AsObject(dt)
	if obj == nil {
		return
	}
	for _, nat := range *obj {
		if !mustGenerate(nat.Attribute.Meta) {
			continue
		}
		f.SectionTemplates = append(f.SectionTemplates, &codegen.SectionTemplate{
			Name:    getName(isPayload),
			Source:  getSource(isPayload),
			Data:    getData(method, data, nat, isPayload),
			FuncMap: fm,
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

func getName(isPayload bool) string {
	if isPayload {
		return "service-payload-method"
	} else {
		return "service-result-method"
	}
}

func getSource(isPayload bool) string {
	if isPayload {
		return servicePayloadMethodT
	} else {
		return serviceResultMethodT
	}
}

func getData(method *expr.MethodExpr, data *service.MethodData, nat *expr.NamedAttributeExpr, isPayload bool) any {
	if isPayload {
		return buildMethodPayloadData(method, nat, data.Payload, codegen.NewNameScope())
	} else {
		return buildMethodResultData(method, nat, data.Result, codegen.NewNameScope())
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

func buildMethodPayloadData(method *expr.MethodExpr, nat *expr.NamedAttributeExpr, payload string, scope *codegen.NameScope) *MethodPayloadData {
	name := codegen.GoifyAtt(nat.Attribute, nat.Name, true)
	typ := scope.GoTypeName(nat.Attribute)
	if method.Payload.IsPrimitivePointer(nat.Name, true) || !expr.IsPrimitive(nat.Attribute.Type) {
		typ = "*" + typ
	}
	return &MethodPayloadData{
		Payload: payload,
		Name:    name,
		Type:    typ,
		Var:     name,
	}
}

func buildMethodResultData(method *expr.MethodExpr, nat *expr.NamedAttributeExpr, result string, scope *codegen.NameScope) *MethodResultData {
	name := codegen.GoifyAtt(nat.Attribute, nat.Name, true)
	typ := scope.GoTypeName(nat.Attribute)
	if method.Result.IsPrimitivePointer(nat.Name, true) || !expr.IsPrimitive(nat.Attribute.Type) {
		typ = "*" + typ
	}
	return &MethodResultData{
		Result: result,
		Name:   name,
		Type:   typ,
		Var:    name,
	}
}

func buildMethodUserTypeData(parent *expr.AttributeExpr, nat *expr.NamedAttributeExpr, userType string, scope *codegen.NameScope) *MethodUserTypeData {
	name := codegen.GoifyAtt(nat.Attribute, nat.Name, true)
	typ := scope.GoTypeName(nat.Attribute)
	if parent.IsPrimitivePointer(nat.Name, true) {
		typ = "*" + typ
	}
	return &MethodUserTypeData{
		UserType: userType,
		Name:     name,
		Type:     typ,
		Var:      name,
	}
}

var servicePayloadMethodT = `
func (p *{{ .Payload }}) Get{{ .Name }}() {{ .Type }} {
	return p.{{ .Var }}
}
`

var serviceResultMethodT = `
func (p *{{ .Result }}) Get{{ .Name }}() {{ .Type }} {
	return p.{{ .Var }}
}
`

var serviceUserTypeMethodT = `
func (p *{{ .UserType }}) Get{{ .Name }}() {{ .Type }} {
	return p.{{ .Var }}
}
`
