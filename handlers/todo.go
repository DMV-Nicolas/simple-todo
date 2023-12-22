package handlers

import (
	"github.com/DMV-Nicolas/todo/db"
	"github.com/DMV-Nicolas/todo/views/todo"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	Queries *db.Queries
}

func (h *TodoHandler) ListTodos(c echo.Context) error {
	todos := h.Queries.ListTodos(0, -1)
	return RenderTemplate(c, todo.List(todos))
}
