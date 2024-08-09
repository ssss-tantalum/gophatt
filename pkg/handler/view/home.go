package view

import (
	"github.com/labstack/echo/v4"
	"github.com/ssss-tantalum/gophatt/pkg/app"
	"github.com/ssss-tantalum/gophatt/templates"
	"github.com/ssss-tantalum/gophatt/templates/layouts"
	pages "github.com/ssss-tantalum/gophatt/templates/pages/home"
)

type HomeHandler struct {
	app *app.App
}

func NewHomeHandler(app *app.App) *HomeHandler {
	return &HomeHandler{
		app: app,
	}
}

func (h HomeHandler) Index(c echo.Context) error {
	props := pages.HomePageProps{
		BaseProps: layouts.BaseProps{
			Title: "Gophatt | Home",
		},
	}

	return templates.Render(c, pages.Home(props))
}
