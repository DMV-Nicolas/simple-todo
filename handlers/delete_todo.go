package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteTodo(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	h.queries.DeleteTodo(uint(id))

	return c.JSON(http.StatusOK, id)
}
