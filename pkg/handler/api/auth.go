package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ssss-tantalum/gophatt/ent/user"
	"github.com/ssss-tantalum/gophatt/pkg/app"
	"github.com/ssss-tantalum/gophatt/pkg/auth"
	v "github.com/ssss-tantalum/gophatt/pkg/validate"
	"github.com/ssss-tantalum/gophatt/templates"
	pages "github.com/ssss-tantalum/gophatt/templates/pages/users"
	"golang.org/x/crypto/bcrypt"
)

var (
	signUpSchema = v.Schema{
		"username": v.Rules(v.Min(2), v.Max(50)),
		"email":    v.Rules(v.Email),
		"password": v.Rules(
			v.ContainsSpecial,
			v.ContainsUpper,
			v.Min(7),
			v.Max(50),
		),
	}

	loginSchema = v.Schema{
		"email":    v.Rules(v.Email),
		"password": v.Rules(v.Required),
	}
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
	var err error
	ctx := c.Request().Context()

	var values pages.SignUpFormValues
	errors, ok := v.Request(c.Request(), &values, signUpSchema)
	if !ok {
		return templates.Render(c, pages.SignUpForm(values, errors))
	}

	if values.Password != values.ConfirmPassword {
		errors.Add("confirm_password", "password do no match")
		return templates.Render(c, pages.SignUpForm(values, errors))
	}

	hashedPassword, err := hashPassword(values.Password)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	tx, err := h.app.Client().Tx(ctx)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	user, err := tx.User.Create().
		SetID(uuid.New()).
		SetUsername(values.Username).
		SetEmail(values.Email).
		SetPassword(hashedPassword).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		c.Logger().Error(err)
		errors.Add("form_error", "dupulicate username or email")
		return templates.Render(c, pages.SignUpForm(values, errors))
	}

	if err := tx.Commit(); err != nil {
		c.Logger().Error(err)
		return err
	}

	if err := auth.SetAuthSession(c, user.ID.String()); err != nil {
		c.Logger().Error(err)
		return err
	}

	return templates.Redirect(c, "/")
}

func (h AuthHandler) Login(c echo.Context) error {
	var err error
	ctx := c.Request().Context()

	var values pages.LoginFormValues
	errors, ok := v.Request(c.Request(), &values, loginSchema)
	if !ok {
		return templates.Render(c, pages.LoginForm(values, errors))
	}

	user, err := h.app.Client().User.Query().
		Where(user.EmailEQ(values.Email)).
		Only(ctx)
	if err != nil {
		errors.Add("form_error", "email or password is invalid")
		return templates.Render(c, pages.LoginForm(values, errors))
	}

	if err := comparePasswords(user.Password, values.Password); err != nil {
		errors.Add("form_error", "email or password is invalid")
		return templates.Render(c, pages.LoginForm(values, errors))
	}

	if err := auth.SetAuthSession(c, user.ID.String()); err != nil {
		c.Logger().Error(err)
		return err
	}

	return templates.Redirect(c, "/")
}

func (h AuthHandler) Logout(c echo.Context) error {
	if isAuthenticated := auth.IsAuthenticated(c); !isAuthenticated {
		return c.Redirect(http.StatusSeeOther, "/")
	}

	if err := auth.DestoryAuthSession(c); err != nil {
		c.Logger().Error(err)
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func hashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func comparePasswords(hash, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return err
	}
	return nil
}
