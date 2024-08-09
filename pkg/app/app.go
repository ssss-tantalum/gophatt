package app

import (
	"context"

	"github.com/ssss-tantalum/gophatt/ent"
)

type appCtxKey struct{}

func ContextWithApp(ctx context.Context, app *App) context.Context {
	ctx = context.WithValue(ctx, appCtxKey{}, app)
	return ctx
}

type App struct {
	ctx context.Context
	cfg *Config

	client *ent.Client
}

func New(ctx context.Context, cfg *Config, client *ent.Client) *App {
	app := &App{
		cfg: cfg,
	}
	app.ctx = ContextWithApp(ctx, app)
	app.client = client

	return app
}

func (app *App) Ctx() context.Context {
	return app.ctx
}

func (app *App) Cfg() *Config {
	return app.cfg
}

func (app *App) Client() *ent.Client {
	return app.client
}
