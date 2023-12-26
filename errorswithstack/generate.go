package errorswithstack

import (
	"path/filepath"
	"strings"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/service"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func init() {
	codegen.RegisterPlugin("errorswithstack", "gen", nil, Generate)
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	for _, f := range files {
		if filepath.Base(f.Path) != "service.go" {
			continue
		}
		var svc *expr.ServiceExpr
		for _, section := range f.Section("service") {
			data, ok := section.Data.(*service.Data)
			if !ok {
				continue
			}
			svc = expr.Root.Service(data.Name)
		}
		if svc == nil {
			continue
		}
		for _, section := range f.Section("source-header") {
			data, ok := section.Data.(map[string]any)
			if !ok {
				continue
			}
			imports, ok := data["Imports"].([]*codegen.ImportSpec)
			if !ok {
				continue
			}
			data["Imports"] = append(imports, codegen.SimpleImport("github.com/cockroachdb/errors/withstack"))
		}
		for _, section := range f.Section("error-init-func") {
			section.Source = strings.ReplaceAll(section.Source, "err,", "withstack.WithStackDepth(err, 1),")
		}
	}
	return files, nil
}
