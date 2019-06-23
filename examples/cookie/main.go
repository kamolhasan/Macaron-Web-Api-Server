package main

import (
	"log"
	"net/http"

	"github.com/go-macaron/auth"
	"gopkg.in/macaron.v1"
)

func main() {

	m := macaron.Classic()

	m.Get("/set", setCookies)
	m.Get("/get", getCookies)
	m.NotFound(func() string {
		return "Unknown URL...!"
	})
	m.Use(auth.Basic("kamol", "hasan"))

	log.Println("Server running... ...")
	log.Println(http.ListenAndServe(":8080", m))
}

func setCookies(ctx *macaron.Context) {
	ctx.SetSecureCookie("kamol", "hasan")
}

func getCookies(ctx *macaron.Context) string {
	name, _ := ctx.GetSecureCookie("kamol")
	return name
}
