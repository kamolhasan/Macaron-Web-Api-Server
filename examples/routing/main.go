package main

import (
	"log"
	"net/http"

	"gopkg.in/macaron.v1"
)

func myHandler(ctx *macaron.Context) string {
	return "Hello " + ctx.Params(":name")
}

func myHandler2(ctx *macaron.Context) string {
	return "Full Name: " + ctx.Params(":firstname") + " " + ctx.Params(":lastname")
}

func myHandler3(ctx *macaron.Context) string {
	return "Print: " + ctx.Params("*")
}

func routing(m *macaron.Macaron) {

	m.Group("/home", func() {
		m.Get("/hello/:name", myHandler)
		m.Get("/hello/:firstname/:lastname", myHandler2)
		m.Get("/all/*", myHandler3)
	})

}

func main() {
	m := macaron.Classic()
	routing(m)

	log.Println("Server running... ...")
	log.Println(http.ListenAndServe(":8080", m))
}
