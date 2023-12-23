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

	handler := handlers.NewHandler(queries)

	app.GET("/", handler.ListTodos)
	app.POST("/", handler.CreateTodo)
	app.DELETE("/:id", handler.DeleteTodo)
	app.PUT("/:id", handler.UpdateTodo)

	app.Start(":5000")

}
