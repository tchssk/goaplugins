# AutoTrailingSlash Plugin

The `autotrailingslash` plugin is a [Goa v3](https://github.com/goadesign/goa/tree/v3) plugin
that makes it possible to handle HTTP requests with a trailing slash.

Goa v3 switched the default HTTP router to chi in [v3.13.0](https://github.com/goadesign/goa/releases/tag/v3.13.0). Due to that change, requests with a trailing slash are no longer handled as same as v3.12.4 or earier. This plugin reproduces the previous httptreemux-like behavior.

## Enabling the Plugin

To enable the plugin simply import as follows:

```go
import (
  _ "github.com/tchssk/goaplugins/v3/autotrailingslash"
  . "goa.design/goa/v3/dsl"
)
```
