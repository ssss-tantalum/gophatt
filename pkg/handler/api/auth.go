package api

import (
	"github.com/labstack/echo/v4"
	"github.com/ssss-tantalum/gophatt/pkg/app"
)

type AuthHandler struct {
	app *app.App
}

func NewAuthHandler(app *app.App) *AuthHandler {
	return &AuthHandler{
		app: app,
	}
}

func (h AuthHandler) SignUp(c echo.Context) error {
	return nil
}
