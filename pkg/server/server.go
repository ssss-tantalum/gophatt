package server

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ssss-tantalum/gophatt/pkg/app"
)

type Server struct {
	app *app.App

	echo *echo.Echo
}

func New(app *app.App) *Server {
	srv := &Server{
		app:  app,
		echo: echo.New(),
	}

	srv.initRoutes(app)

	return srv
}

func (s *Server) Start(port string) error {
	return s.echo.Start(port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
