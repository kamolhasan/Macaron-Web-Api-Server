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
	m.Use(auth.BasicFunc(func(username string, password string) bool {
		return auth.SecureCompare(username, "kamol") && auth.SecureCompare(password, "hasan")
	}))

	log.Println("Server running... ...")
	log.Println(http.ListenAndServe(":8080", m))
}

func myHandler(ctx *macaron.Context) string {
	return "Got response from: " + ctx.Req.RequestURI
}
