package handlers

import (
	"net/http"

	"github.com/DMV-Nicolas/todo/db"
	"github.com/DMV-Nicolas/todo/views/todo"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	Queries *db.Queries
}

type createTodoRequest struct {
	Title string `json:"title"`
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	return nil
}

type listTodosRequest struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

func (h *TodoHandler) ListTodos(c echo.Context) error {
	req := new(listTodosRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	todos := h.Queries.ListTodos(req.Offset, -1)
	return RenderTemplate(c, todo.List(todos))
}
