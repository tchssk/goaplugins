package expr

import (
	"fmt"

	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

type (
	// HTTPEndpointExpr describes a HTTP endpoint.
	HTTPEndpointExpr struct {
		// MethodExpr is the underlying method expression.
		MethodExpr *expr.MethodExpr
		// Service is the parent service.
		Service *expr.HTTPServiceExpr
		// OptionalBody indicates that the HTTP request body is optional.
		OptionalBody bool
	}
)

// Name of HTTP endpoint
func (e *HTTPEndpointExpr) Name() string {
	return e.MethodExpr.Name
}

// EvalName returns the generic expression name used in error messages.
func (e *HTTPEndpointExpr) EvalName() string {
	var prefix, suffix string
	if e.Name() != "" {
		suffix = fmt.Sprintf("HTTP endpoint %#v", e.Name())
	} else {
		suffix = "unnamed HTTP endpoint"
	}
	if e.Service != nil {
		prefix = e.Service.EvalName() + " "
	}
	return prefix + suffix
}

// Validate validates the endpoint expression.
func (e *HTTPEndpointExpr) Validate() *eval.ValidationErrors {
	verr := new(eval.ValidationErrors)
	// TODO Implement.
	return verr
}
