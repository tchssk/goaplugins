# ErrorsWithStack Plugin

The `errorswithstack` plugin is a [Goa v3](https://github.com/goadesign/goa/tree/v3) plugin
that adds a stack trace to every original service error.
This plugin depends on [`github.com/cockroachdb/errors/withstack`](github.com/cockroachdb/errors/withstack).

## Enabling the Plugin

To enable the plugin simply import as follows:

```go
import (
  _ "github.com/tchssk/goaplugins/v3/errorswithstack"
  . "goa.design/goa/v3/dsl"
)
```

## Effects on Code Generation

Enabling the plugin changes the behavior of the `gen` command of the `goa` tool.

The `gen` command output is modified as follows:

1. All error initialization helper functions are modified to add a stack trace to the original service error using [`WithStackDepth()`](https://pkg.go.dev/github.com/cockroachdb/errors/withstack#WithStackDepth).

    ```diff
    func MakeInternalError(err error) *goa.ServiceError {
    -	return goa.NewServiceError(err, "internal_error", false, false, true)
    +	return goa.NewServiceError(withstack.WithStackDepth(err, 1), "internal_error", false, false, true)
    }
    ```

## Middleware

You can capture errors using the Goa endpoint middleware. The topmost caller can be extracted using [`GetOneLineSource()`](https://pkg.go.dev/github.com/cockroachdb/errors/withstack#GetOneLineSource):

```go
func ErrorLogger(logger *log.Logger) func(goa.Endpoint) goa.Endpoint {
	return func(e goa.Endpoint) goa.Endpoint {
		return goa.Endpoint(func(ctx context.Context, req any) (any, error) {
			res, err := e(ctx, req)
			if err != nil {
				file, line, _, ok := withstack.GetOneLineSource(err)
				if ok {
					logger.Printf("%s:%d: %v", file, line, err) // file.go:15 something went wrong
				}
			}
			return res, err
		})
	}
}
```

or [`GetReportableStackTrace()`](https://pkg.go.dev/github.com/cockroachdb/errors/withstack#GetReportableStackTrace):

```go
func ErrorLogger(logger *log.Logger) func(goa.Endpoint) goa.Endpoint {
	return func(e goa.Endpoint) goa.Endpoint {
		return goa.Endpoint(func(ctx context.Context, req any) (any, error) {
			res, err := e(ctx, req)
			if err != nil {
				if st := withstack.GetReportableStackTrace(errors.Unwrap(err)); st != nil {
					if len(st.Frames) >= 1 {
						frame := st.Frames[len(st.Frames)-1]
						logger.Printf("%s:%d: %v", frame.AbsPath, frame.Lineno, err) // /path/to/file.go:15 something went wrong
					}
				}
			}
			return res, err
		})
	}
}
```

The error's underlying concrete value is [`ServiceError`](https://pkg.go.dev/goa.design/goa/v3/pkg#ServiceError). You can also create conditions by making type assertion.

```go
func ErrorLogger(logger *log.Logger) func(goa.Endpoint) goa.Endpoint {
	return func(e goa.Endpoint) goa.Endpoint {
		return goa.Endpoint(func(ctx context.Context, req any) (any, error) {
			res, err := e(ctx, req)
			if err != nil {
				if serviceError, ok := err.(*goa.ServiceError); ok {
               if serviceError.Fault {
						file, line, _, ok := withstack.GetOneLineSource(err)
						if ok {
							logger.Printf("%s:%d: %v", file, line, err) // file.go:15 something went wrong
						}
               }
				}
			}
			return res, err
		})
	}
}
```

You can also report errors to [Sentry](https://sentry.io) using [`report.ReportError`](https://pkg.go.dev/github.com/cockroachdb/errors/report#ReportError).

```go
func ErrorReporter() func(goa.Endpoint) goa.Endpoint {
	return func(e goa.Endpoint) goa.Endpoint {
		return goa.Endpoint(func(ctx context.Context, req any) (any, error) {
			res, err := e(ctx, req)
			if err != nil {
				report.ReportError(err)
			}
			return res, err
		})
	}
}
```
