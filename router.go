package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) initRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/health_check", a.HealthCheck).Methods(http.MethodGet)
}
