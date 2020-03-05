package testdata

var ServiceWithOptionalBodyCode = `
// Service is the Service1 service interface.
type Service interface {
	// Method1 implements Method1.
	Method1(context.Context, *Payload) (err error)
	// Method2 implements Method2.
	Method2(context.Context, *Payload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Service1"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"Method1", "Method2"}

// Payload is the payload type of the Service1 service Method1 method.
type Payload struct {
	Attribute *string
	EmptyBody bool
}
`

var ServiceWithoutOptionalBodyCode = `
// Service is the Service2 service interface.
type Service interface {
	// Method1 implements Method1.
	Method1(context.Context, *Payload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Service2"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"Method1"}

// Payload is the payload type of the Service2 service Method1 method.
type Payload struct {
	Attribute *string
}
`

var EncodeDecodeWithOptionalBodyCode = `// EncodeMethod1Response returns an encoder for responses returned by the
// Service1 Method1 endpoint.
func EncodeMethod1Response(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeMethod1Request returns a decoder for requests sent to the Service1
// Method1 endpoint.
func DecodeMethod1Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body Method1RequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err != io.EOF {
				return nil, goa.DecodePayloadError(err.Error())
			}
		}
		payload := NewMethod1Payload(&body)

		return payload, nil
	}
}

// EncodeMethod2Response returns an encoder for responses returned by the
// Service1 Method2 endpoint.
func EncodeMethod2Response(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeMethod2Request returns a decoder for requests sent to the Service1
// Method2 endpoint.
func DecodeMethod2Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body Method2RequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewMethod2Payload(&body)

		return payload, nil
	}
}
`

var EncodeDecodeWithoutOptionalBodyCode = `// EncodeMethod1Response returns an encoder for responses returned by the
// Service2 Method1 endpoint.
func EncodeMethod1Response(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeMethod1Request returns a decoder for requests sent to the Service2
// Method1 endpoint.
func DecodeMethod1Request(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body Method1RequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewMethod1Payload(&body)

		return payload, nil
	}
}
`

var TypesWithOptionalBodyCode = `// Method1RequestBody is the type of the "Service1" service "Method1" endpoint
// HTTP request body.
type Method1RequestBody struct {
	Attribute *string ` + "`" + `form:"Attribute,omitempty" json:"Attribute,omitempty" xml:"Attribute,omitempty"` + "`" + `
}

// Method2RequestBody is the type of the "Service1" service "Method2" endpoint
// HTTP request body.
type Method2RequestBody struct {
	Attribute *string ` + "`" + `form:"Attribute,omitempty" json:"Attribute,omitempty" xml:"Attribute,omitempty"` + "`" + `
}

// NewMethod1Payload builds a Service1 service Method1 endpoint payload. It
// allows an empty body.
func NewMethod1Payload(body *Method1RequestBody) *service1.Payload {
	var v *service1.Payload
	if body == nil {
		v = &service1.Payload{
			EmptyBody: true,
		}
	} else {
		v = &service1.Payload{
			Attribute: body.Attribute,
		}
	}
	return v
}

// NewMethod2Payload builds a Service1 service Method2 endpoint payload.
func NewMethod2Payload(body *Method2RequestBody) *service1.Payload {
	v := &service1.Payload{
		Attribute: body.Attribute,
	}
	return v
}
`

var TypesWithoutOptionalBodyCode = `// Method1RequestBody is the type of the "Service2" service "Method1" endpoint
// HTTP request body.
type Method1RequestBody struct {
	Attribute *string ` + "`" + `form:"Attribute,omitempty" json:"Attribute,omitempty" xml:"Attribute,omitempty"` + "`" + `
}

// NewMethod1Payload builds a Service2 service Method1 endpoint payload.
func NewMethod1Payload(body *Method1RequestBody) *service2.Payload {
	v := &service2.Payload{
		Attribute: body.Attribute,
	}
	return v
}
`
