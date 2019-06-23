// Macaron Practice 2: Basic server with custom handler fuction

package main

import (
	"log"
	"net/http"

	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	m.Get("/", myHandlerFunc)

	log.Println("Server is running...")
	log.Println(http.ListenAndServe(":8080", m))
}

func myHandlerFunc(ctx *macaron.Context) string {
	return "the request path is: " + ctx.Req.RequestURI
}
