package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ea2305/go-cards/controllers"
)

func (a *App) initRoutes(r *mux.Router) {
	r.HandleFunc("/health_check", controllers.HealthCheck).Methods(http.MethodGet)
}
