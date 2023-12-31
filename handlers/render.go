package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// RenderTemplate takes a echo context and a templ component to render
func RenderTemplate(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
