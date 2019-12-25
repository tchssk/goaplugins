package testdata

var SimpleCode = `
// Service is the Service service interface.
type Service interface {
	// Method implements Method.
	Method(context.Context, *Payload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Service"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"Method"}

// Payload is the payload type of the Service service Method method.
type Payload struct {
	AttributeBoolean         *bool
	AttributeInt             *int
	AttributeInt32           *int32
	AttributeInt64           *int64
	AttributeUInt            *uint
	AttributeUInt32          *uint32
	AttributeUInt64          *uint64
	AttributeFloat32         *float32
	AttributeFloat64         *float64
	AttributeString          *string
	AttributeBytes           []byte
	AttributeAny             interface{}
	RequiredAttributeBoolean bool
	RequiredAttributeInt     int
	RequiredAttributeInt32   int32
	RequiredAttributeInt64   int64
	RequiredAttributeUInt    uint
	RequiredAttributeUInt32  uint32
	RequiredAttributeUInt64  uint64
	RequiredAttributeFloat32 float32
	RequiredAttributeFloat64 float64
	RequiredAttributeString  string
	RequiredAttributeBytes   []byte
	RequiredAttributeAny     interface{}
	Ignored                  *string
}

func (p *Payload) GetAttributeBoolean() *bool {
	return p.AttributeBoolean
}

func (p *Payload) GetAttributeInt() *int {
	return p.AttributeInt
}

func (p *Payload) GetAttributeInt32() *int32 {
	return p.AttributeInt32
}

func (p *Payload) GetAttributeInt64() *int64 {
	return p.AttributeInt64
}

func (p *Payload) GetAttributeUInt() *uint {
	return p.AttributeUInt
}

func (p *Payload) GetAttributeUInt32() *uint32 {
	return p.AttributeUInt32
}

func (p *Payload) GetAttributeUInt64() *uint64 {
	return p.AttributeUInt64
}

func (p *Payload) GetAttributeFloat32() *float32 {
	return p.AttributeFloat32
}

func (p *Payload) GetAttributeFloat64() *float64 {
	return p.AttributeFloat64
}

func (p *Payload) GetAttributeString() *string {
	return p.AttributeString
}

func (p *Payload) GetAttributeBytes() []byte {
	return p.AttributeBytes
}

func (p *Payload) GetAttributeAny() interface{} {
	return p.AttributeAny
}

func (p *Payload) GetRequiredAttributeBoolean() bool {
	return p.RequiredAttributeBoolean
}

func (p *Payload) GetRequiredAttributeInt() int {
	return p.RequiredAttributeInt
}

func (p *Payload) GetRequiredAttributeInt32() int32 {
	return p.RequiredAttributeInt32
}

func (p *Payload) GetRequiredAttributeInt64() int64 {
	return p.RequiredAttributeInt64
}

func (p *Payload) GetRequiredAttributeUInt() uint {
	return p.RequiredAttributeUInt
}

func (p *Payload) GetRequiredAttributeUInt32() uint32 {
	return p.RequiredAttributeUInt32
}

func (p *Payload) GetRequiredAttributeUInt64() uint64 {
	return p.RequiredAttributeUInt64
}

func (p *Payload) GetRequiredAttributeFloat32() float32 {
	return p.RequiredAttributeFloat32
}

func (p *Payload) GetRequiredAttributeFloat64() float64 {
	return p.RequiredAttributeFloat64
}

func (p *Payload) GetRequiredAttributeString() string {
	return p.RequiredAttributeString
}

func (p *Payload) GetRequiredAttributeBytes() []byte {
	return p.RequiredAttributeBytes
}

func (p *Payload) GetRequiredAttributeAny() interface{} {
	return p.RequiredAttributeAny
}
`
