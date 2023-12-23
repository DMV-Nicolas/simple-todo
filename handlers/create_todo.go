package handlers

import (
	"errors"
	"net/http"

	"github.com/DMV-Nicolas/todo/views/todo"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateTodo(c echo.Context) error {
	title := c.FormValue("title")
	if title == "" {
		err := errors.New("title not provided")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	t := h.queries.CreateTodo(title)

	return RenderTemplate(c, todo.Get(t))
}
