package testdata

import (
	. "goa.design/goa/v3/dsl"
)

var SingleServiceDSL = func() {
	var grandChild = Type("GrandChild", func() {
		Attribute("AttributeBoolean", Boolean)
	})
	var child = Type("Child", func() {
		Attribute("AttributeBoolean", Boolean)
		Attribute("AttributeGrandChild", grandChild)
		Attribute("RequiredAttributeGrandChild", grandChild)
		Required(
			"RequiredAttributeGrandChild",
		)
	})
	var payload = Type("Payload", func() {
		Attribute("AttributeBoolean", Boolean)
		Attribute("AttributeInt", Int)
		Attribute("AttributeInt32", Int32)
		Attribute("AttributeInt64", Int64)
		Attribute("AttributeUInt", UInt)
		Attribute("AttributeUInt32", UInt32)
		Attribute("AttributeUInt64", UInt64)
		Attribute("AttributeFloat32", Float32)
		Attribute("AttributeFloat64", Float64)
		Attribute("AttributeString", String)
		Attribute("AttributeBytes", Bytes)
		Attribute("AttributeAny", Any)
		Attribute("AttributeChild", child)
		Attribute("RequiredAttributeBoolean", Boolean)
		Attribute("RequiredAttributeInt", Int)
		Attribute("RequiredAttributeInt32", Int32)
		Attribute("RequiredAttributeInt64", Int64)
		Attribute("RequiredAttributeUInt", UInt)
		Attribute("RequiredAttributeUInt32", UInt32)
		Attribute("RequiredAttributeUInt64", UInt64)
		Attribute("RequiredAttributeFloat32", Float32)
		Attribute("RequiredAttributeFloat64", Float64)
		Attribute("RequiredAttributeString", String)
		Attribute("RequiredAttributeBytes", Bytes)
		Attribute("RequiredAttributeAny", Any)
		Attribute("RequiredAttributeChild", child)
		Attribute("Ignored", String, func() {
			Meta("attributegetter:generate", "false")
		})
		Required(
			"RequiredAttributeBoolean",
			"RequiredAttributeInt",
			"RequiredAttributeInt32",
			"RequiredAttributeInt64",
			"RequiredAttributeUInt",
			"RequiredAttributeUInt32",
			"RequiredAttributeUInt64",
			"RequiredAttributeFloat32",
			"RequiredAttributeFloat64",
			"RequiredAttributeString",
			"RequiredAttributeBytes",
			"RequiredAttributeAny",
			"RequiredAttributeChild",
		)
	})
	var result = ResultType("application/vnd.result", func() {
		Attribute("AttributeBoolean", Boolean)
		Attribute("AttributeInt", Int)
		Attribute("AttributeInt32", Int32)
		Attribute("AttributeInt64", Int64)
		Attribute("AttributeUInt", UInt)
		Attribute("AttributeUInt32", UInt32)
		Attribute("AttributeUInt64", UInt64)
		Attribute("AttributeFloat32", Float32)
		Attribute("AttributeFloat64", Float64)
		Attribute("AttributeString", String)
		Attribute("AttributeBytes", Bytes)
		Attribute("AttributeAny", Any)
		Attribute("AttributeChild", child)
		Attribute("RequiredAttributeBoolean", Boolean)
		Attribute("RequiredAttributeInt", Int)
		Attribute("RequiredAttributeInt32", Int32)
		Attribute("RequiredAttributeInt64", Int64)
		Attribute("RequiredAttributeUInt", UInt)
		Attribute("RequiredAttributeUInt32", UInt32)
		Attribute("RequiredAttributeUInt64", UInt64)
		Attribute("RequiredAttributeFloat32", Float32)
		Attribute("RequiredAttributeFloat64", Float64)
		Attribute("RequiredAttributeString", String)
		Attribute("RequiredAttributeBytes", Bytes)
		Attribute("RequiredAttributeAny", Any)
		Attribute("RequiredAttributeChild", child)
		Attribute("Ignored", String, func() {
			Meta("attributegetter:generate", "false")
		})
		Required(
			"RequiredAttributeBoolean",
			"RequiredAttributeInt",
			"RequiredAttributeInt32",
			"RequiredAttributeInt64",
			"RequiredAttributeUInt",
			"RequiredAttributeUInt32",
			"RequiredAttributeUInt64",
			"RequiredAttributeFloat32",
			"RequiredAttributeFloat64",
			"RequiredAttributeString",
			"RequiredAttributeBytes",
			"RequiredAttributeAny",
			"RequiredAttributeChild",
		)
	})
	Service("Service", func() {
		Method("Method", func() {
			Payload(payload)
			Result(result)
			HTTP(func() {
				GET("/")
			})
		})
	})
}
