package clibuildpayload

import (
	"path/filepath"
	"strings"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

func init() {
	codegen.RegisterPlugin("codegendisabler-gen-http-client-cli-clibuildpayload", "gen", nil, Generate)
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	for _, f := range files {
		if !strings.HasPrefix(f.Path, filepath.Join("gen", "http")) || !strings.HasSuffix(f.Path, filepath.Join("client", "cli.go")) {
			continue
		}
		var sections []*codegen.SectionTemplate
		for _, section := range f.SectionTemplates {
			if section.Name != "cli-build-payload" {
				sections = append(sections, section)
			}
		}
		f.SectionTemplates = sections
	}
	return files, nil
}
