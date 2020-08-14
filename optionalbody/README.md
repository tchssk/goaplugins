# OptionalBody Plugin

The `optionalbody` plugin is a [Goa v3](https://github.com/goadesign/goa/tree/v3) plugin
that makes it possible to define an optional body for HTTP server endpoints.

## Note

Goa supports the optional HTTP request body since v3.2.1.

## Enabling the Plugin

To enable the plugin and make use of the OptionalBody DSL simply import both the `optionalbody` and
the `dsl` packages as follows:

```go
import (
  "github.com/tchssk/goaplugins/optionalbody/dsl"
  . "goa.design/goa/v3/dsl"
)
```
Note the use of blank identifier to import the `optionalbody` package which is necessary
as the package is imported solely for its side-effects (initialization).

## Effects on Code Generation

Enabling the plugin changes the behavior of the `gen` command of the `goa` tool.

The `gen` command output is modified as follows:

1. A new private field is appended to the HTTP server payload type. The field  indicates
   whether the HTTP body is empty.
2. A new payload initializer is appended to the HTTP server code.
3. A new method is appended to the HTTP server payload type. The method reports
   whether the HTTP body is empty.
4. The HTTP request decoder is modified to bypass the missing payload error and the
   request payload validator when the request body is empty.

## Design

This plugin adds the following functions to the goa DSL:

* `OptionalPayload` is used in method `HTTP` DSL to define a HTTP request body.

Here is an example defining an optional body.

```go
var _ = Service("user", func() {
  Method("show", func() {
    Payload(Userpayload)
    HTTP(func() {
      POST("/")
      dsl.OptionalBody()  // Sets the HTTP request body optional.
    })
  })
})
```
