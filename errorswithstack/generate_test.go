package errorswithstack_test

import (
	"bytes"
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tchssk/goaplugins/v3/errorswithstack"
	"github.com/tchssk/goaplugins/v3/errorswithstack/testdata"
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
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			root := codegen.RunDSL(t, c.DSL)
			if len(root.Services) != 1 {
				t.Fatalf("got %d services, expected 1", len(root.Services))
			}
			services := service.NewServicesData(root)
			fs := service.Files("", root.Services[0], services, make(map[string][]string))
			if fs == nil {
				t.Fatalf("got nil file, expected not nil")
			}
			if _, err := errorswithstack.Generate("", []eval.Root{root}, fs); err != nil {
				t.Fatal(err)
			}
			buf := new(bytes.Buffer)
			for _, f := range fs {
				for _, s := range f.SectionTemplates[1:] {
					if err := s.Write(buf); err != nil {
						t.Fatal(err)
					}
				}
			}
			bs, err := format.Source(buf.Bytes())
			if err != nil {
				t.Fatal(err)
			}
			code := string(bs)
			assert.Equal(t, c.Code, code)
		})
	}
}
