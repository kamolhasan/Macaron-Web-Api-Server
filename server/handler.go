package server

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/macaron.v1"
)

func GetBooks(ctx *macaron.Context) {
	fmt.Println("GetBooks called!")

	var books []Book
	if err := Engine.Find(&books); err != nil {
		fmt.Println("got books!")
		panic(err)
	}
	bookList := BookList{
		Items: books,
	}
	ctx.JSON(http.StatusOK, bookList)
}

func GetBook(ctx *macaron.Context) {
	book := Book{
		ID: ctx.Params("id"),
	}
	if exist, err := Engine.Get(&book); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else if exist {
		ctx.JSON(http.StatusFound, book)
	} else {
		ctx.JSON(http.StatusNotFound, nil)
	}
}

func PostBook(ctx *macaron.Context, list BookList) string {
	st, _ := ctx.Req.Body().String()

	for _, val := range list.Items {
		log.Println(val)
	}
	return st + "\n postbook.. called!"
}

func UpdateBook(ctx *macaron.Context, book Book) string {
	st, _ := ctx.Req.Body().String()
	log.Println(book)
	return st + "UpdateBook... called!"
}
