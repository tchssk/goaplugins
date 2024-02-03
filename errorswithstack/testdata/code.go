package testdata

var SimpleCode = `
// Service is the Service service interface.
type Service interface {
	// Method implements Method.
	Method(context.Context) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Service"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"Method"}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return goa.NewServiceError(withstack.WithStackDepth(err, 1), "internal_error", false, false, false)
}
`
