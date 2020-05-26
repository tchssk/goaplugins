# CodegenDisabler Plugin

The `codegendisabler` plugin is a [Goa v3](https://github.com/goadesign/goa/tree/v3) plugin
that makes it possible to disable a part of the code generator.

## Enabling the Plugin

To enable the plugin simply import one of the codegendisabler packages as follows:

```go
import (
  _ "github.com/tchssk/goaplugins/codegendisabler/gen/http/client/types/client_body_init"
  . "goa.design/goa/v3/dsl"
)
```

This code disables a part of the code generator which generates by a section template named
`client-body-init` to `gen/http/<service>/client/types.go`.
