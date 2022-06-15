package routes

import (
	"fmt"
	"net/http"

	"github.com/ea2305/go-cards/controllers"
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	var regexPattern = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"
	var ApiIdFormat = fmt.Sprintf("/api/v1/decks/{id:%v}", regexPattern)
	r.HandleFunc("/api/v1/health_check", controllers.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/decks", controllers.CreateDeck).Methods(http.MethodPost)
	r.HandleFunc(ApiIdFormat, controllers.GetDeck).Methods(http.MethodGet)
	r.HandleFunc(ApiIdFormat, controllers.DrawCard).Methods(http.MethodPatch)
}
