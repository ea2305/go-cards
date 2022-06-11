package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) initRoutes(r *mux.Router) {
	r.HandleFunc("/health_check", a.HealthCheck).Methods(http.MethodGet)
}
