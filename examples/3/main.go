// Macaron practice 3: m.Use() examples, order of middleware execution

package main

import (
	"log"
	"net/http"

	"gopkg.in/macaron.v1"
)

func main() {
	myMiddleware := func(ctx *macaron.Context) {
		if ctx.Req.Header.Get("API-KEY") != "secret" {
			ctx.Resp.WriteHeader(http.StatusUnauthorized)
		}
	}
	myMiddleware2 := func(ctx *macaron.Context) {
		log.Println("middleware 2 successfully executed!")
	}

	m := macaron.Classic()
	m.Get("/", myHandler)

	// m.Use() always follows the order, i.e. myMiddleware2 will executed after myMiddleware
	m.Use(myMiddleware)
	m.Use(myMiddleware2)

	log.Println("Server running... ...")
	log.Println(http.ListenAndServe(":8080", m))
}

func myHandler(ctx *macaron.Context) string {
	return "Got response from: " + ctx.Req.RequestURI
}
