package app

import (
	"fmt"
	"net/http"
)

type App struct {
	name   string
	router http.Handler
}

func NewApp() *App {
	app := &App{
		router: LoadRouter(),
		name:   "start",
	}

	return app
}

func (a *App) Start() error {
	err := http.ListenAndServe(":3000", a.router)

	if err != nil {
		return fmt.Errorf("failed to start server %w", err)
	}
	return nil

}
