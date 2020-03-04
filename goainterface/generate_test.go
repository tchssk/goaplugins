package goainterface_test

import (
	"bytes"
	"fmt"
	"go/format"
	"testing"

	"github.com/tchssk/goaplugins/goainterface"
	"github.com/tchssk/goaplugins/goainterface/testdata"
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
			codegen.RunDSL(t, c.DSL)
			if len(expr.Root.Services) != 1 {
				t.Fatalf("got %d services, expected 1", len(expr.Root.Services))
			}
			fs := service.File("", expr.Root.Services[0])
			if fs == nil {
				t.Fatalf("got nil file, expected not nil")
			}
			goainterface.Generate("", []eval.Root{expr.Root}, []*codegen.File{fs})
			buf := new(bytes.Buffer)
			for _, s := range fs.SectionTemplates[1:] {
				if err := s.Write(buf); err != nil {
					t.Fatal(err)
				}
			}
			bs, err := format.Source(buf.Bytes())
			if err != nil {
				fmt.Println(buf.String())
				t.Fatal(err)
			}
			code := string(bs)
			if code != c.Code {
				t.Errorf("%s: got\n%s\ngot vs. expected:\n%s", c.Name, code, codegen.Diff(t, code, c.Code))
			}
		})
	}
}
