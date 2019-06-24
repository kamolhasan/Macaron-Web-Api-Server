package main

import (
	"fmt"
	"log"

	"github.com/go-macaron/binding"

	"gopkg.in/macaron.v1"
)

type myStruct struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func main() {

	m := macaron.Classic()

	m.Post("/", binding.Json(myStruct{}), func(ctx *macaron.Context, str myStruct) string {
		return fmt.Sprintf("Name: %s \nEmail: %s \nMessage: %s\n", str.Name, str.Email, str.Message)
	})

	log.Println("Server running... ...")
	m.Run()
}
