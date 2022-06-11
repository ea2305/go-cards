package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App struct{}

type AppConfig struct {
	Addr string
}

func (a *App) initApp(app App, config AppConfig) *http.Server {
	r := mux.NewRouter()
	app.initRoutes(r)

	server := &http.Server{
		Addr:    config.Addr,
		Handler: r,
	}
	return server
}
