package goaversionremover_test

import (
	"bytes"
	"go/format"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tchssk/goaplugins/v3/goaversionremover"
	"github.com/tchssk/goaplugins/v3/goaversionremover/testdata"
	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/service"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func TestService(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"single service", testdata.SingleServiceDSL, testdata.SimpleCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := codegen.RunDSL(t, c.DSL)
			require.Len(t, root.Services, 1)
			generation, err := codegen.NewGeneration("goa.design/goa/example", []eval.Root{root})
			require.NoError(t, err)
			plan, err := service.NewPlan(root, generation, expr.NewExampleGenerator(root.API.RandomizerFactory))
			require.NoError(t, err)
			require.NoError(t, generation.Freeze())
			require.NoError(t, plan.Link())
			fs, err := service.Files(plan)
			require.NoError(t, err)
			require.NotNil(t, fs)
			_, err = goaversionremover.Generate("", []eval.Root{root}, fs)
			require.NoError(t, err)
			buf := new(bytes.Buffer)
			for _, f := range fs {
				if filepath.Base(f.Path) != "service.go" {
					continue
				}
				for _, s := range f.SectionTemplates {
					require.NoError(t, s.Write(buf))
				}
			}
			bs, err := format.Source(buf.Bytes())
			require.NoError(t, err)
			code := string(bs)
			assert.Equal(t, c.Code, code)
		})
	}
}
