package parseendpoint

import (
	"path/filepath"
	"strings"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

func init() {
	codegen.RegisterPlugin("codegendisabler-gen-http-cli-cli-parseendpoint", "gen", nil, Generate)
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	for _, f := range files {
		if !strings.HasPrefix(f.Path, filepath.Join("gen", "http", "cli")) || !strings.HasSuffix(f.Path, "cli.go") {
			continue
		}
		var sections []*codegen.SectionTemplate
		for _, section := range f.SectionTemplates {
			if section.Name != "parse-endpoint" {
				sections = append(sections, section)
			}
		}
		f.SectionTemplates = sections
	}
	return files, nil
}
