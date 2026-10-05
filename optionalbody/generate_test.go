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
	httpcodegen "goa.design/goa/v3/http/codegen"
)

func TestService(t *testing.T) {
	cases := []struct {
		Name    string
		DSL     func()
		Service int
		Code    string
	}{
		{"method with optional body", testdata.SimpleDSL, 0, testdata.ServiceWithOptionalBodyCode},
		{"method without optional body", testdata.SimpleDSL, 1, testdata.ServiceWithoutOptionalBodyCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := codegen.RunDSL(t, c.DSL)
			require.Len(t, root.Services, 2)
			services := service.NewServicesData(root)
			fs := service.Files("", root.Services[c.Service], services, make(map[string][]string))
			require.NotNil(t, fs)
			_, err := optionalbody.Update("", []eval.Root{root}, fs)
			require.NoError(t, err)
			buf := new(bytes.Buffer)
			for _, f := range fs {
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
			services := httpcodegen.CreateHTTPServices(root)
			fs := httpcodegen.ServerFiles("", services)
			require.Len(t, fs, 4)
			_, err := optionalbody.Update("", []eval.Root{root}, fs)
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
			services := httpcodegen.CreateHTTPServices(root)
			var files []*codegen.File
			fs := httpcodegen.ServerFiles("", services)
			require.Len(t, fs, 4)
			files = append(files, fs...)
			fs = httpcodegen.ServerTypeFiles("", services)
			require.Len(t, fs, 2)
			files = append(files, fs...)
			_, err := optionalbody.Update("", []eval.Root{root}, files)
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
