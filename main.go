package main

import (
	"log"

	"github.com/DMV-Nicolas/simple-todo/db"
	"github.com/DMV-Nicolas/simple-todo/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	// connect to SQLite database
	sqlDB, err := db.ConnectDatabase("todo.db")
	if err != nil {
		log.Fatal("cannot connect database")
	}

	// create an object queries for the database functions
	queries := db.NewQueries(sqlDB)

	app := echo.New()

	// create an handler object for the handlers functions
	handler := handlers.NewHandler(queries)

	app.GET("/", handler.ListTodos)
	app.POST("/", handler.CreateTodo)
	app.DELETE("/:id", handler.DeleteTodo)
	app.PUT("/:id", handler.UpdateTodo)

	// start server
	app.Start(":5000")
}
