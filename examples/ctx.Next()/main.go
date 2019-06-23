// |-------------------------------------------------------|
// | Execution order:                                      |
// |-------------------------------------------------------|
// |2019/06/23 16:50:08 Server running... ...              |
// |[Macaron] 2019-06-23 16:50:11: Started GET / for [::1] |
// |[Macaron] middleware1: before a request!               |
// |[Macaron] middleware2: before a request!               |
// |2019/06/23 16:50:11 sending response....!              |
// |[Macaron] middleware2: after a request!                |
// |[Macaron] middleware1: after a request!                |
// |-------------------------------------------------------|

package main

import (
	"log"
	"net/http"

	"gopkg.in/macaron.v1"
)

func main() {
	middleware1 := func(ctx *macaron.Context, lg *log.Logger) {
		lg.Println("middleware1: before a request!")

		ctx.Next()

		lg.Println("middleware1: after a request!")
	}

	middleware2 := func(ctx *macaron.Context, lg *log.Logger) {
		lg.Println("middleware2: before a request!")

		ctx.Next()

		lg.Println("middleware2: after a request!")
	}

	m := macaron.Classic()
	m.Get("/", myHandler)

	m.Use(middleware1)
	m.Use(middleware2)

	log.Println("Server running... ...")
	log.Println(http.ListenAndServe(":8080", m))
}

func myHandler(ctx *macaron.Context) string {
	log.Println("sending response....!")
	return "Got response from: " + ctx.Req.RequestURI
}
