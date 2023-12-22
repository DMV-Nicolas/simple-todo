package main

import (
	"log"

	"github.com/DMV-Nicolas/todo/db"
	"github.com/DMV-Nicolas/todo/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	sqlDB, err := db.ConnectDatabase("todo.db")
	if err != nil {
		log.Fatal("cannot connect database")
	}

	queries := db.NewQueries(sqlDB)

	app := echo.New()
	todoHandler := &handlers.TodoHandler{
		Queries: queries,
	}

	app.GET("/todo", todoHandler.ListTodos)

	app.Start(":5000")

}
