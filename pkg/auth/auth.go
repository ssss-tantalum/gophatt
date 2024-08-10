package auth

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	SESSION_KEY string = "authenticate-sessions"
	AUTH_KEY    string = "autenticated"
	USER_ID_KEY string = "user_id"
)

func IsAuthenticated(c echo.Context) bool {
	sess, _ := session.Get(SESSION_KEY, c)

	return sess.Values[AUTH_KEY] == true
}

func SetAuthSession(c echo.Context, userID string) error {
	sess, err := session.Get(SESSION_KEY, c)
	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	sess.Values = map[interface{}]interface{}{
		AUTH_KEY:    true,
		USER_ID_KEY: userID,
	}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return nil
}

func DestoryAuthSession(c echo.Context) error {
	sess, _ := session.Get(SESSION_KEY, c)
	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return nil
}
