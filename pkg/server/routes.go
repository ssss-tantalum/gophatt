package server

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ssss-tantalum/gophatt/pkg/app"
	"github.com/ssss-tantalum/gophatt/pkg/handler/api"
	"github.com/ssss-tantalum/gophatt/pkg/handler/view"
)

func (s *Server) initRoutes(app *app.App) {
	router := s.echo

	// Serve static files
	router.Static("/static", "static")

	// Error handling
	router.HTTPErrorHandler = customHTTPErrorHandler

	// Middlewares
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(session.Middleware(sessions.NewCookieStore([]byte(app.Cfg().SecretKey))))

	// View routes
	{
		homeHandler := view.NewHomeHandler(app)
		router.GET("/", homeHandler.Index)

	}

	// API routes
	apiRoutes := router.Group("/api")
	{
		authHandler := api.NewAuthHandler(app)
		apiRoutes.POST("/auth/signup", authHandler.SignUp)
		apiRoutes.POST("/auth/login", authHandler.Login)
		apiRoutes.GET("/auth/logout", authHandler.Logout, AuthMiddleware)
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)

	switch code {
	case 401:
	case 404:
	default:
	}
}
