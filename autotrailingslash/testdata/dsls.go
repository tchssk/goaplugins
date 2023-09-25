package testdata

import (
	. "goa.design/goa/v3/dsl"
)

var SingleServiceDSL = func() {
	Service("Service", func() {
		Method("Method", func() {
			HTTP(func() {
				GET("foo")
			})
		})
	})
}
