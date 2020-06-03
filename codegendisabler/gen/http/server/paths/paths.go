package paths

import (
	"path/filepath"
	"strings"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
)

func init() {
	codegen.RegisterPlugin("codegendisabler-gen-http-server-paths", "gen", nil, Generate)
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	var fs []*codegen.File
	for _, f := range files {
		if !strings.HasPrefix(f.Path, filepath.Join("gen", "http")) || !strings.HasSuffix(f.Path, filepath.Join("server", "paths.go")) {
			fs = append(fs, f)
		}
	}
	return fs, nil
}
