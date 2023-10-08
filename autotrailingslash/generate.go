package autotrailingslash

import (
	"strings"

	"goa.design/goa/v3/codegen"
	"goa.design/goa/v3/eval"
	"goa.design/goa/v3/expr"
)

func init() {
	codegen.RegisterPlugin("autotrailingslash", "gen", Prepare, Generate)
}

func Generate(genpkg string, roots []eval.Root, files []*codegen.File) ([]*codegen.File, error) {
	return files, nil
}

func Prepare(genpkg string, roots []eval.Root) error {
	for _, root := range roots {
		if r, ok := root.(*expr.RootExpr); ok {
			if r.API == nil {
				continue
			}
			if r.API.HTTP == nil {
				continue
			}
			for _, service := range r.API.HTTP.Services {
				for _, endpoint := range service.HTTPEndpoints {
					var routes []*expr.RouteExpr
					for _, route := range endpoint.Routes {
						if strings.Contains(route.Path, "/{*") {
							continue
						}
						routes = append(routes, &expr.RouteExpr{
							Method:   route.Method,
							Path:     route.Path + "/./",
							Endpoint: route.Endpoint,
							Meta:     route.Meta,
						})

					}
					endpoint.Routes = append(endpoint.Routes, routes...)
				}
			}
		}
	}
	return nil
}
