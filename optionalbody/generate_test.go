package optionalbody_test

import (
	"bytes"
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tchssk/goaplugins/v3/optionalbody"
	"github.com/tchssk/goaplugins/v3/optionalbody/testdata"
	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/service"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
	httpcodegen "goa.design/goa/v3/http/codegen"
)

func TestService(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		Path string
		Code string
	}{
		{"method with optional body", testdata.SimpleDSL, "gen/service1/service.go", testdata.ServiceWithOptionalBodyCode},
		{"method without optional body", testdata.SimpleDSL, "gen/service2/service.go", testdata.ServiceWithoutOptionalBodyCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := codegen.RunDSL(t, c.DSL)
			require.Len(t, root.Services, 2)
			generation, err := codegen.NewGeneration("goa.design/goa/example", []eval.Root{root})
			require.NoError(t, err)
			plan, err := service.NewPlan(root, generation, expr.NewExampleGenerator(root.API.RandomizerFactory))
			require.NoError(t, err)
			require.NoError(t, generation.Freeze())
			require.NoError(t, plan.Link())
			fs, err := service.Files(plan)
			require.NoError(t, err)
			require.NotNil(t, fs)
			_, err = optionalbody.Update("", []eval.Root{root}, fs)
			require.NoError(t, err)
			buf := new(bytes.Buffer)
			for _, f := range fs {
				if f.Path != c.Path {
					continue
				}
				for _, s := range f.SectionTemplates[1:] {
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

func TestEncodeDecode(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		File int
		Code string
	}{
		{"method with optional body", testdata.SimpleDSL, 2, testdata.EncodeDecodeWithOptionalBodyCode},
		{"method without optional body", testdata.SimpleDSL, 3, testdata.EncodeDecodeWithoutOptionalBodyCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := codegen.RunDSL(t, c.DSL)
			generation, err := codegen.NewGeneration("goa.design/goa/example", []eval.Root{root})
			require.NoError(t, err)
			servicePlan, err := service.NewPlan(root, generation, expr.NewExampleGenerator(root.API.RandomizerFactory))
			require.NoError(t, err)
			plans, err := httpcodegen.NewPlans(generation, httpcodegen.PlanInput{Root: root, Service: servicePlan})
			require.NoError(t, err)
			require.NoError(t, generation.Freeze())
			require.NoError(t, servicePlan.Link())
			require.NoError(t, plans[0].Link())
			fs := plans[0].ServerFiles()
			require.Len(t, fs, 4)
			_, err = optionalbody.Update("", []eval.Root{root}, fs)
			require.NoError(t, err)
			buf := new(bytes.Buffer)
			for _, s := range fs[c.File].SectionTemplates[1:] {
				require.NoError(t, s.Write(buf))
			}
			bs, err := format.Source(buf.Bytes())
			require.NoError(t, err)
			code := string(bs)
			assert.Equal(t, c.Code, code)
		})
	}
}

func TestTypes(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		File int
		Code string
	}{
		{"method with optional service", testdata.SimpleDSL, 0, testdata.TypesWithOptionalBodyCode},
		{"method without optional service", testdata.SimpleDSL, 1, testdata.TypesWithoutOptionalBodyCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := codegen.RunDSL(t, c.DSL)
			generation, err := codegen.NewGeneration("goa.design/goa/example", []eval.Root{root})
			require.NoError(t, err)
			servicePlan, err := service.NewPlan(root, generation, expr.NewExampleGenerator(root.API.RandomizerFactory))
			require.NoError(t, err)
			plans, err := httpcodegen.NewPlans(generation, httpcodegen.PlanInput{Root: root, Service: servicePlan})
			require.NoError(t, err)
			require.NoError(t, generation.Freeze())
			require.NoError(t, servicePlan.Link())
			require.NoError(t, plans[0].Link())
			var files []*codegen.File
			fs := plans[0].ServerFiles()
			require.Len(t, fs, 4)
			files = append(files, fs...)
			fs = plans[0].ServerTypeFiles()
			require.Len(t, fs, 2)
			files = append(files, fs...)
			_, err = optionalbody.Update("", []eval.Root{root}, files)
			require.NoError(t, err)
			buf := new(bytes.Buffer)
			for _, s := range fs[c.File].SectionTemplates[1:] {
				require.NoError(t, s.Write(buf))
			}
			bs, err := format.Source(buf.Bytes())
			require.NoError(t, err)
			code := string(bs)
			assert.Equal(t, c.Code, code)
		})
	}
}
