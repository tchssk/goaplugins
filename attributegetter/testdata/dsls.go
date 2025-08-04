package testdata

import (
	. "goa.design/goa/v3/dsl"
)

var SingleServiceDSL = func() {
	var grandChild = Type("grand-child", func() {
		Attribute("attribute-boolean", Boolean)
	})
	var child = Type("child", func() {
		Attribute("attribute-boolean", Boolean)
		Attribute("attribute-grand-child", grandChild)
		Attribute("required-attribute-grand-child", grandChild)
		Required(
			"required-attribute-grand-child",
		)
	})
	var payload = Type("payload", func() {
		Attribute("attribute-boolean", Boolean)
		Attribute("attribute-int", Int)
		Attribute("attribute-int32", Int32)
		Attribute("attribute-int64", Int64)
		Attribute("attribute-uint", UInt)
		Attribute("attribute-uint32", UInt32)
		Attribute("attribute-uint64", UInt64)
		Attribute("attribute-float32", Float32)
		Attribute("attribute-float64", Float64)
		Attribute("attribute-string", String)
		Attribute("attribute-bytes", Bytes)
		Attribute("attribute-any", Any)
		Attribute("attribute-array", ArrayOf(String))
		Attribute("attribute-map", MapOf(String, String))
		Attribute("attribute-child", child)
		Attribute("required-attribute-boolean", Boolean)
		Attribute("required-attribute-int", Int)
		Attribute("required-attribute-int32", Int32)
		Attribute("required-attribute-int64", Int64)
		Attribute("required-attribute-uint", UInt)
		Attribute("required-attribute-uint32", UInt32)
		Attribute("required-attribute-uint64", UInt64)
		Attribute("required-attribute-float32", Float32)
		Attribute("required-attribute-float64", Float64)
		Attribute("required-attribute-string", String)
		Attribute("required-attribute-bytes", Bytes)
		Attribute("required-attribute-any", Any)
		Attribute("required-attribute-array", ArrayOf(String))
		Attribute("required-attribute-map", MapOf(String, String))
		Attribute("required-attribute-child", child)
		Attribute("ignored", String, func() {
			Meta("attributegetter:generate", "false")
		})
		Required(
			"required-attribute-boolean",
			"required-attribute-int",
			"required-attribute-int32",
			"required-attribute-int64",
			"required-attribute-uint",
			"required-attribute-uint32",
			"required-attribute-uint64",
			"required-attribute-float32",
			"required-attribute-float64",
			"required-attribute-string",
			"required-attribute-bytes",
			"required-attribute-any",
			"required-attribute-array",
			"required-attribute-map",
			"required-attribute-child",
		)
	})
	var result = ResultType("application/vnd.result", func() {
		Attribute("attribute-boolean", Boolean)
		Attribute("attribute-int", Int)
		Attribute("attribute-int32", Int32)
		Attribute("attribute-int64", Int64)
		Attribute("attribute-uint", UInt)
		Attribute("attribute-uint32", UInt32)
		Attribute("attribute-uint64", UInt64)
		Attribute("attribute-float32", Float32)
		Attribute("attribute-float64", Float64)
		Attribute("attribute-string", String)
		Attribute("attribute-bytes", Bytes)
		Attribute("attribute-any", Any)
		Attribute("attribute-array", ArrayOf(String))
		Attribute("attribute-map", MapOf(String, String))
		Attribute("attribute-child", child)
		Attribute("required-attribute-boolean", Boolean)
		Attribute("required-attribute-int", Int)
		Attribute("required-attribute-int32", Int32)
		Attribute("required-attribute-int64", Int64)
		Attribute("required-attribute-uint", UInt)
		Attribute("required-attribute-uint32", UInt32)
		Attribute("required-attribute-uint64", UInt64)
		Attribute("required-attribute-float32", Float32)
		Attribute("required-attribute-float64", Float64)
		Attribute("required-attribute-string", String)
		Attribute("required-attribute-bytes", Bytes)
		Attribute("required-attribute-any", Any)
		Attribute("required-attribute-array", ArrayOf(String))
		Attribute("required-attribute-map", MapOf(String, String))
		Attribute("required-attribute-child", child)
		Attribute("ignored", String, func() {
			Meta("attributegetter:generate", "false")
		})
		Required(
			"required-attribute-boolean",
			"required-attribute-int",
			"required-attribute-int32",
			"required-attribute-int64",
			"required-attribute-uint",
			"required-attribute-uint32",
			"required-attribute-uint64",
			"required-attribute-float32",
			"required-attribute-float64",
			"required-attribute-string",
			"required-attribute-bytes",
			"required-attribute-any",
			"required-attribute-array",
			"required-attribute-map",
			"required-attribute-child",
		)
	})
	Service("service", func() {
		Method("method", func() {
			Payload(payload)
			Result(result)
			HTTP(func() {
				GET("/")
			})
		})
	})
}

var ServiceWithCollectionDSL = func() {
	var result = ResultType("application/vnd.result", func() {
		Attribute("attribute-boolean", Boolean)
	})
	Service("service", func() {
		Method("method", func() {
			Payload(Empty)
			Result(CollectionOf(result))
			HTTP(func() {
				GET("/")
			})
		})
	})
}
