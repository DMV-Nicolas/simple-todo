package handlers

import (
	"github.com/DMV-Nicolas/simple-todo/views/todo"
	"github.com/labstack/echo/v4"
)

func (h *Handler) ListTodos(c echo.Context) error {
	todos := h.queries.ListTodos(0, -1)
	return RenderTemplate(c, todo.List(todos))
}
