package attributegetter_test

import (
	"bytes"
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tchssk/goaplugins/v3/attributegetter"
	"github.com/tchssk/goaplugins/v3/attributegetter/testdata"
	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/codegen/service"
	"goa.design/goa/v3/eval"
)

func TestService(t *testing.T) {
	cases := []struct {
		Name string
		DSL  func()
		Code string
	}{
		{"single service", testdata.SingleServiceDSL, testdata.SimpleCode},
		{"service with collection", testdata.ServiceWithCollectionDSL, testdata.ServiceWithCollectionCode},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := codegen.RunDSL(t, c.DSL)
			require.Len(t, root.Services, 1)
			services := service.NewServicesData(root)
			fs := service.Files("", root.Services[0], services, make(map[string][]string))
			require.NotNil(t, fs)
			_, err := attributegetter.Generate("", []eval.Root{root}, fs)
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
