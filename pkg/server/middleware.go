package server

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ssss-tantalum/gophatt/pkg/auth"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get(auth.SESSION_KEY, c)
		if auth, ok := sess.Values[auth.AUTH_KEY].(bool); !ok || !auth {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		if userID, ok := sess.Values[auth.USER_ID_KEY].(string); ok && len(userID) != 0 {
			c.Set(auth.USER_ID_KEY, userID)
		}

		return next(c)
	}
}
