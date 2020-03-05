package testdata

import (
	"github.com/tchssk/goaplugins/optionalbody/dsl"
	. "goa.design/goa/v3/dsl"
)

var SimpleDSL = func() {
	var payload = Type("Payload", func() {
		Attribute("Attribute", String)
	})
	Service("Service1", func() {
		Method("Method1", func() {
			Payload(payload)
			HTTP(func() {
				POST("/")
				dsl.OptionalBody()
			})
		})
		Method("Method2", func() {
			Payload(payload)
			HTTP(func() {
				POST("/")
			})
		})
	})
	Service("Service2", func() {
		Method("Method1", func() {
			Payload(payload)
			HTTP(func() {
				POST("/")
			})
		})
	})
}
