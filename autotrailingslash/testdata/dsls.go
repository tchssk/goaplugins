package testdata

import (
	. "goa.design/goa/v3/dsl"
)

var SingleServiceDSL = func() {
	Service("Service", func() {
		Method("Method1", func() {
			HTTP(func() {
				GET("/foo")
			})
		})
		Method("Method2", func() {
			Payload(func() {
				Attribute("param1")
			})
			HTTP(func() {
				GET("/foo/{param1}")
			})
		})
	})
}
