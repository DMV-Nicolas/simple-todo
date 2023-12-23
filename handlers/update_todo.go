package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UpdateTodo(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var done bool
	doneStr := c.FormValue("done")
	if doneStr == "" {
		done = false
	} else {
		done = true
	}

	h.queries.UpdateTodo(uint(id), done)

	return c.JSON(http.StatusOK, nil)
}
