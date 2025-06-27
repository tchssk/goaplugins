package testdata

var SimpleCode = `
// Service is the Service service interface.
type Service interface {
	// Method implements Method.
	Method(context.Context, *Payload) (res *Result, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "test api"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Service"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"Method"}

type Child struct {
	AttributeBoolean            *bool
	AttributeGrandChild         *GrandChild
	RequiredAttributeGrandChild *GrandChild
}

type GrandChild struct {
	AttributeBoolean *bool
}

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
	AttributeAny             any
	AttributeChild           *Child
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
	RequiredAttributeAny     any
	RequiredAttributeChild   *Child
	Ignored                  *string
}

// Result is the result type of the Service service Method method.
type Result struct {
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
	AttributeAny             any
	AttributeChild           *Child
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
	RequiredAttributeAny     any
	RequiredAttributeChild   *Child
	Ignored                  *string
}

// NewResult initializes result type Result from viewed result type Result.
func NewResult(vres *serviceviews.Result) *Result {
	return newResult(vres.Projected)
}

// NewViewedResult initializes viewed result type Result from result type
// Result using the given view.
func NewViewedResult(res *Result, view string) *serviceviews.Result {
	p := newResultView(res)
	return &serviceviews.Result{Projected: p, View: "default"}
}

// newResult converts projected type Result to service type Result.
func newResult(vres *serviceviews.ResultView) *Result {
	res := &Result{
		AttributeBoolean:       vres.AttributeBoolean,
		AttributeInt:           vres.AttributeInt,
		AttributeInt32:         vres.AttributeInt32,
		AttributeInt64:         vres.AttributeInt64,
		AttributeUInt:          vres.AttributeUInt,
		AttributeUInt32:        vres.AttributeUInt32,
		AttributeUInt64:        vres.AttributeUInt64,
		AttributeFloat32:       vres.AttributeFloat32,
		AttributeFloat64:       vres.AttributeFloat64,
		AttributeString:        vres.AttributeString,
		AttributeBytes:         vres.AttributeBytes,
		AttributeAny:           vres.AttributeAny,
		RequiredAttributeBytes: vres.RequiredAttributeBytes,
		RequiredAttributeAny:   vres.RequiredAttributeAny,
		Ignored:                vres.Ignored,
	}
	if vres.RequiredAttributeBoolean != nil {
		res.RequiredAttributeBoolean = *vres.RequiredAttributeBoolean
	}
	if vres.RequiredAttributeInt != nil {
		res.RequiredAttributeInt = *vres.RequiredAttributeInt
	}
	if vres.RequiredAttributeInt32 != nil {
		res.RequiredAttributeInt32 = *vres.RequiredAttributeInt32
	}
	if vres.RequiredAttributeInt64 != nil {
		res.RequiredAttributeInt64 = *vres.RequiredAttributeInt64
	}
	if vres.RequiredAttributeUInt != nil {
		res.RequiredAttributeUInt = *vres.RequiredAttributeUInt
	}
	if vres.RequiredAttributeUInt32 != nil {
		res.RequiredAttributeUInt32 = *vres.RequiredAttributeUInt32
	}
	if vres.RequiredAttributeUInt64 != nil {
		res.RequiredAttributeUInt64 = *vres.RequiredAttributeUInt64
	}
	if vres.RequiredAttributeFloat32 != nil {
		res.RequiredAttributeFloat32 = *vres.RequiredAttributeFloat32
	}
	if vres.RequiredAttributeFloat64 != nil {
		res.RequiredAttributeFloat64 = *vres.RequiredAttributeFloat64
	}
	if vres.RequiredAttributeString != nil {
		res.RequiredAttributeString = *vres.RequiredAttributeString
	}
	if vres.AttributeChild != nil {
		res.AttributeChild = transformServiceviewsChildViewToChild(vres.AttributeChild)
	}
	if vres.RequiredAttributeChild != nil {
		res.RequiredAttributeChild = transformServiceviewsChildViewToChild(vres.RequiredAttributeChild)
	}
	return res
}

// newResultView projects result type Result to projected type ResultView using
// the "default" view.
func newResultView(res *Result) *serviceviews.ResultView {
	vres := &serviceviews.ResultView{
		AttributeBoolean:         res.AttributeBoolean,
		AttributeInt:             res.AttributeInt,
		AttributeInt32:           res.AttributeInt32,
		AttributeInt64:           res.AttributeInt64,
		AttributeUInt:            res.AttributeUInt,
		AttributeUInt32:          res.AttributeUInt32,
		AttributeUInt64:          res.AttributeUInt64,
		AttributeFloat32:         res.AttributeFloat32,
		AttributeFloat64:         res.AttributeFloat64,
		AttributeString:          res.AttributeString,
		AttributeBytes:           res.AttributeBytes,
		AttributeAny:             res.AttributeAny,
		RequiredAttributeBoolean: &res.RequiredAttributeBoolean,
		RequiredAttributeInt:     &res.RequiredAttributeInt,
		RequiredAttributeInt32:   &res.RequiredAttributeInt32,
		RequiredAttributeInt64:   &res.RequiredAttributeInt64,
		RequiredAttributeUInt:    &res.RequiredAttributeUInt,
		RequiredAttributeUInt32:  &res.RequiredAttributeUInt32,
		RequiredAttributeUInt64:  &res.RequiredAttributeUInt64,
		RequiredAttributeFloat32: &res.RequiredAttributeFloat32,
		RequiredAttributeFloat64: &res.RequiredAttributeFloat64,
		RequiredAttributeString:  &res.RequiredAttributeString,
		RequiredAttributeBytes:   res.RequiredAttributeBytes,
		RequiredAttributeAny:     res.RequiredAttributeAny,
		Ignored:                  res.Ignored,
	}
	if res.AttributeChild != nil {
		vres.AttributeChild = transformChildToServiceviewsChildView(res.AttributeChild)
	}
	if res.RequiredAttributeChild != nil {
		vres.RequiredAttributeChild = transformChildToServiceviewsChildView(res.RequiredAttributeChild)
	}
	return vres
}

// transformServiceviewsChildViewToChild builds a value of type *Child from a
// value of type *serviceviews.ChildView.
func transformServiceviewsChildViewToChild(v *serviceviews.ChildView) *Child {
	if v == nil {
		return nil
	}
	res := &Child{
		AttributeBoolean: v.AttributeBoolean,
	}
	if v.AttributeGrandChild != nil {
		res.AttributeGrandChild = transformServiceviewsGrandChildViewToGrandChild(v.AttributeGrandChild)
	}
	if v.RequiredAttributeGrandChild != nil {
		res.RequiredAttributeGrandChild = transformServiceviewsGrandChildViewToGrandChild(v.RequiredAttributeGrandChild)
	}

	return res
}

// transformServiceviewsGrandChildViewToGrandChild builds a value of type
// *GrandChild from a value of type *serviceviews.GrandChildView.
func transformServiceviewsGrandChildViewToGrandChild(v *serviceviews.GrandChildView) *GrandChild {
	if v == nil {
		return nil
	}
	res := &GrandChild{
		AttributeBoolean: v.AttributeBoolean,
	}

	return res
}

// transformChildToServiceviewsChildView builds a value of type
// *serviceviews.ChildView from a value of type *Child.
func transformChildToServiceviewsChildView(v *Child) *serviceviews.ChildView {
	if v == nil {
		return nil
	}
	res := &serviceviews.ChildView{
		AttributeBoolean: v.AttributeBoolean,
	}
	if v.AttributeGrandChild != nil {
		res.AttributeGrandChild = transformGrandChildToServiceviewsGrandChildView(v.AttributeGrandChild)
	}
	if v.RequiredAttributeGrandChild != nil {
		res.RequiredAttributeGrandChild = transformGrandChildToServiceviewsGrandChildView(v.RequiredAttributeGrandChild)
	}

	return res
}

// transformGrandChildToServiceviewsGrandChildView builds a value of type
// *serviceviews.GrandChildView from a value of type *GrandChild.
func transformGrandChildToServiceviewsGrandChildView(v *GrandChild) *serviceviews.GrandChildView {
	if v == nil {
		return nil
	}
	res := &serviceviews.GrandChildView{
		AttributeBoolean: v.AttributeBoolean,
	}

	return res
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

func (p *Payload) GetAttributeAny() any {
	return p.AttributeAny
}

func (p *Payload) GetAttributeChild() *Child {
	return p.AttributeChild
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

func (p *Payload) GetRequiredAttributeAny() any {
	return p.RequiredAttributeAny
}

func (p *Payload) GetRequiredAttributeChild() *Child {
	return p.RequiredAttributeChild
}

func (p *Result) GetAttributeBoolean() *bool {
	return p.AttributeBoolean
}

func (p *Result) GetAttributeInt() *int {
	return p.AttributeInt
}

func (p *Result) GetAttributeInt32() *int32 {
	return p.AttributeInt32
}

func (p *Result) GetAttributeInt64() *int64 {
	return p.AttributeInt64
}

func (p *Result) GetAttributeUInt() *uint {
	return p.AttributeUInt
}

func (p *Result) GetAttributeUInt32() *uint32 {
	return p.AttributeUInt32
}

func (p *Result) GetAttributeUInt64() *uint64 {
	return p.AttributeUInt64
}

func (p *Result) GetAttributeFloat32() *float32 {
	return p.AttributeFloat32
}

func (p *Result) GetAttributeFloat64() *float64 {
	return p.AttributeFloat64
}

func (p *Result) GetAttributeString() *string {
	return p.AttributeString
}

func (p *Result) GetAttributeBytes() []byte {
	return p.AttributeBytes
}

func (p *Result) GetAttributeAny() any {
	return p.AttributeAny
}

func (p *Result) GetAttributeChild() *Child {
	return p.AttributeChild
}

func (p *Result) GetRequiredAttributeBoolean() bool {
	return p.RequiredAttributeBoolean
}

func (p *Result) GetRequiredAttributeInt() int {
	return p.RequiredAttributeInt
}

func (p *Result) GetRequiredAttributeInt32() int32 {
	return p.RequiredAttributeInt32
}

func (p *Result) GetRequiredAttributeInt64() int64 {
	return p.RequiredAttributeInt64
}

func (p *Result) GetRequiredAttributeUInt() uint {
	return p.RequiredAttributeUInt
}

func (p *Result) GetRequiredAttributeUInt32() uint32 {
	return p.RequiredAttributeUInt32
}

func (p *Result) GetRequiredAttributeUInt64() uint64 {
	return p.RequiredAttributeUInt64
}

func (p *Result) GetRequiredAttributeFloat32() float32 {
	return p.RequiredAttributeFloat32
}

func (p *Result) GetRequiredAttributeFloat64() float64 {
	return p.RequiredAttributeFloat64
}

func (p *Result) GetRequiredAttributeString() string {
	return p.RequiredAttributeString
}

func (p *Result) GetRequiredAttributeBytes() []byte {
	return p.RequiredAttributeBytes
}

func (p *Result) GetRequiredAttributeAny() any {
	return p.RequiredAttributeAny
}

func (p *Result) GetRequiredAttributeChild() *Child {
	return p.RequiredAttributeChild
}

func (p *Child) GetAttributeBoolean() *bool {
	return p.AttributeBoolean
}

func (p *Child) GetAttributeGrandChild() GrandChild {
	return p.AttributeGrandChild
}

func (p *Child) GetRequiredAttributeGrandChild() GrandChild {
	return p.RequiredAttributeGrandChild
}

func (p *GrandChild) GetAttributeBoolean() *bool {
	return p.AttributeBoolean
}
`
