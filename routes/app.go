package routes

import (
	"net/http"

	"github.com/ea2305/go-cards/controllers"
	"github.com/ea2305/go-cards/util"
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/health_check", controllers.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/decks", controllers.CreateDeck).Methods(http.MethodPost)
	r.HandleFunc(util.ApiIdFormat, controllers.GetDeck).Methods(http.MethodGet)
	r.HandleFunc(util.ApiIdFormat, controllers.DrawCard).Methods(http.MethodPatch)
	// handle 404 cases with json format

	r.NotFoundHandler = http.HandlerFunc(controllers.HandleNotFound)
}
