package testdata

var SimpleCode = `// Server lists the Service service endpoint HTTP handlers.
type Server struct {
	Mounts  []*MountPoint
	Method1 http.Handler
	Method2 http.Handler
	Method3 http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the Service service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *service.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Method1", "GET", "/foo"},
			{"Method1", "GET", "/bar/"},
			{"Method1", "GET", "/foo/"},
			{"Method2", "GET", "/foo/{param1}"},
			{"Method2", "GET", "/foo/{param1}/"},
			{"Method3", "GET", "/foo/{*param2}"},
		},
		Method1: NewMethod1Handler(e.Method1, mux, decoder, encoder, errhandler, formatter),
		Method2: NewMethod2Handler(e.Method2, mux, decoder, encoder, errhandler, formatter),
		Method3: NewMethod3Handler(e.Method3, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "Service" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Method1 = m(s.Method1)
	s.Method2 = m(s.Method2)
	s.Method3 = m(s.Method3)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return service.MethodNames[:] }

// Mount configures the mux to serve the Service endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountMethod1Handler(mux, h.Method1)
	MountMethod2Handler(mux, h.Method2)
	MountMethod3Handler(mux, h.Method3)
}

// Mount configures the mux to serve the Service endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountMethod1Handler configures the mux to serve the "Service" service
// "Method1" endpoint.
func MountMethod1Handler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/foo", f)
	mux.Handle("GET", "/bar/", f)
	mux.Handle("GET", "/foo/", f)
}

// NewMethod1Handler creates a HTTP handler which loads the HTTP request and
// calls the "Service" service "Method1" endpoint.
func NewMethod1Handler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeMethod1Response(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Method1")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Service")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil && errhandler != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			if errhandler != nil {
				errhandler(ctx, w, err)
			}
		}
	})
}

// MountMethod2Handler configures the mux to serve the "Service" service
// "Method2" endpoint.
func MountMethod2Handler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/foo/{param1}", f)
	mux.Handle("GET", "/foo/{param1}/", f)
}

// NewMethod2Handler creates a HTTP handler which loads the HTTP request and
// calls the "Service" service "Method2" endpoint.
func NewMethod2Handler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMethod2Request(mux, decoder)
		encodeResponse = EncodeMethod2Response(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Method2")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Service")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil && errhandler != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil && errhandler != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			if errhandler != nil {
				errhandler(ctx, w, err)
			}
		}
	})
}

// MountMethod3Handler configures the mux to serve the "Service" service
// "Method3" endpoint.
func MountMethod3Handler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/foo/{*param2}", f)
}

// NewMethod3Handler creates a HTTP handler which loads the HTTP request and
// calls the "Service" service "Method3" endpoint.
func NewMethod3Handler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMethod3Request(mux, decoder)
		encodeResponse = EncodeMethod3Response(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Method3")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Service")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil && errhandler != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil && errhandler != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			if errhandler != nil {
				errhandler(ctx, w, err)
			}
		}
	})
}

// EncodeMethod1Response returns an encoder for responses returned by the
// Service Method1 endpoint.
func EncodeMethod1Response(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// EncodeMethod2Response returns an encoder for responses returned by the
// Service Method2 endpoint.
func EncodeMethod2Response(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeMethod2Request returns a decoder for requests sent to the Service
// Method2 endpoint.
func DecodeMethod2Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (*service.Method2Payload, error) {
	return func(r *http.Request) (*service.Method2Payload, error) {
		var (
			param1 string

			params = mux.Vars(r)
		)
		param1 = params["param1"]
		payload := NewMethod2Payload(param1)

		return payload, nil
	}
}

// EncodeMethod3Response returns an encoder for responses returned by the
// Service Method3 endpoint.
func EncodeMethod3Response(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeMethod3Request returns a decoder for requests sent to the Service
// Method3 endpoint.
func DecodeMethod3Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (*service.Method3Payload, error) {
	return func(r *http.Request) (*service.Method3Payload, error) {
		var (
			param2 string

			params = mux.Vars(r)
		)
		param2 = params["param2"]
		payload := NewMethod3Payload(param2)

		return payload, nil
	}
}
`
