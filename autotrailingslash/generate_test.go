package autotrailingslash_test

import (
	"bytes"
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tchssk/goaplugins/v3/autotrailingslash"
	"github.com/tchssk/goaplugins/v3/autotrailingslash/testdata"
	"goa.design/goa/v3/eval"
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
			root := httpcodegen.RunHTTPDSL(t, c.DSL)
			if err := autotrailingslash.Prepare("", []eval.Root{root}); err != nil {
				t.Fatal(err)
			}
			services := httpcodegen.CreateHTTPServices(root)
			fs := httpcodegen.ServerFiles("", services)
			if fs == nil {
				t.Fatalf("got nil file, expected not nil")
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
