package app

import (
	"context"
	sh "sneakers/internal/api/sneaker"
	"sneakers/internal/config"
	sr "sneakers/internal/repository/sneaker"
	"sneakers/internal/router"
	ss "sneakers/internal/service/sneaker"

	"github.com/labstack/echo/v4"
)

type App struct {
	c      *config.Config
	router *echo.Echo
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{
		router: router.New(),
	}
	return app, app.initDeps(ctx)
}

func (a *App) initDeps(ctx context.Context) error {
	intits := []func(context.Context) error{
		a.initConfig,
		a.initService,
	}
	for _, f := range intits {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	a.c = config.DefaultConfig()
	return nil
}

func (a *App) initService(_ context.Context) error {
	v1 := a.router.Group("/api/v1")

	repo := sr.NewRepository()
	sneakerService := ss.NewService(repo)
	handler := sh.NewHandler(sneakerService)
	handler.Register(v1)

	return nil
}

func (a *App) Run() error {
	err := a.router.Start(a.c.Addr())
	if err != nil {
		return err
	}
	return nil
}
