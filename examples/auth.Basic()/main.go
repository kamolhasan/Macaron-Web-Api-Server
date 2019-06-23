package main

import (
	"log"
	"net/http"

	"github.com/go-macaron/auth"

	"gopkg.in/macaron.v1"
)

func main() {

	m := macaron.Classic()
	m.Get("/", myHandler)
	m.Use(auth.Basic("name", "kamol"))

	log.Println("Server running... ...")
	log.Println(http.ListenAndServe(":8080", m))
}

func myHandler(ctx *macaron.Context) string {
	return "Got response from: " + ctx.Req.RequestURI
}
