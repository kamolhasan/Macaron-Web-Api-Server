package main

import (
	"log"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var engine *xorm.Engine

type Books struct {
	ID     int64
	Name   string
	Author string
}

func main() {
	var err error
	connStr := "user=postgres password=postgres host=127.0.0.1 port=5432 dbname=demo sslmode=disable"

	if engine, err = xorm.NewEngine("postgres", connStr); err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create("postgres.log")
	if err != nil {
		println(err.Error())
	}
	logger := xorm.NewSimpleLogger(f)
	logger.ShowSQL(true)
	engine.SetLogger(logger)
	engine.Limit()
}
