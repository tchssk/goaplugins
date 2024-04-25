package goaversionremover

import (
	"strings"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

func init() {
	codegen.RegisterPlugin("goaversionremover", "gen", nil, Generate)
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	for _, f := range files {
		for _, section := range f.Section("source-header") {
			section.Source = strings.ReplaceAll(section.Source, " {{.ToolVersion}}", "")
		}
	}
	return files, nil
}
