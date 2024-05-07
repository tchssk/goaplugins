package expr

import (
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

// Root is the design root expression.
var Root = &RootExpr{
	HTTPEndpoints: map[string]map[string]*HTTPEndpointExpr{},
}

type (
	// RootExpr keeps track of the optional bodies defined in the design.
	RootExpr struct {
		// HTTPEndpoints lists all the service endpoints indexed by service and method strings.
		HTTPEndpoints map[string]map[string]*HTTPEndpointExpr
	}
)

// Register design root with eval engine.
func init() {
	if err := eval.Register(Root); err != nil {
		panic(err) // bug
	}
}

// EvalName returns the name used in error messages.
func (r *RootExpr) EvalName() string {
	return "Optional Body plugin"
}

// WalkSets iterates over the API-level and service-level optional body definitions.
func (r *RootExpr) WalkSets(walk eval.SetWalker) {
	oexps := make(eval.ExpressionSet, 0, len(r.HTTPEndpoints))
	for _, s := range r.HTTPEndpoints {
		for _, m := range s {
			oexps = append(oexps, m)
		}
	}
	walk(oexps)
}

// DependsOn tells the eval engine to run the goa DSL first.
func (r *RootExpr) DependsOn() []eval.Root {
	return []eval.Root{expr.Root}
}

// Packages returns the import path to the Go packages that make
// up the DSL. This is used to skip frames that point to files
// in these packages when computing the location of errors.
func (r *RootExpr) Packages() []string {
	return []string{"github.com/tchssk/goaplugins/v3/optionalbody/dsl"}
}
