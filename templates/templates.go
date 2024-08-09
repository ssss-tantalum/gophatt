package templates

import (
	"github.com/a-h/templ"
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, component)
}

func Redirect(c echo.Context, url string) error {
	return htmx.NewResponse().Redirect(url).Write(c.Response().Writer)
}
