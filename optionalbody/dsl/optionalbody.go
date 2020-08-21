package dsl

import (
	"goa.design/goa/v3/dsl"
	"goa.design/goa/v3/eval"
	goaexpr "goa.design/goa/v3/expr"

	"github.com/tchssk/goaplugins/v3/optionalbody/expr"

	// Register code generators for the Optional Body plugin
	_ "github.com/tchssk/goaplugins/v3/optionalbody"
)

// OptionalBody describes an optional HTTP request body.
//
// OptionalBody must appear in a Method HTTP expression.
//
func OptionalBody(args ...interface{}) {
	e, ok := eval.Current().(*goaexpr.HTTPEndpointExpr)
	if !ok {
		eval.IncompatibleDSL()
		return
	}
	if _, ok := expr.Root.HTTPEndpoints[e.MethodExpr.Service.Name]; !ok {
		expr.Root.HTTPEndpoints[e.MethodExpr.Service.Name] = make(map[string]*expr.HTTPEndpointExpr)
	}
	expr.Root.HTTPEndpoints[e.MethodExpr.Service.Name][e.MethodExpr.Name] = &expr.HTTPEndpointExpr{
		Service:      e.Service,
		MethodExpr:   e.MethodExpr,
		OptionalBody: true,
	}
	if len(args) != 0 {
		dsl.Body(args...)
		return
	}
}
