package autotrailingslash_test

import (
	"bytes"
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tchssk/goaplugins/v3/autotrailingslash"
	"github.com/tchssk/goaplugins/v3/autotrailingslash/testdata"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
	httpcodegen "goa.design/goa/v3/http/codegen"
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
			root := expr.RunDSL(t, c.DSL)
			require.NoError(t, autotrailingslash.Prepare("", []eval.Root{root}))
			services := httpcodegen.CreateHTTPServices(root)
			fs := httpcodegen.ServerFiles("", services)
			require.NotNil(t, fs)
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
