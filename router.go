package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) initRoutes(r *mux.Router) {
	var regexPattern = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"
	var getDeckWithId = fmt.Sprintf("/api/v1/decks/{id:%v}", regexPattern)
	r.HandleFunc("/api/v1/health_check", a.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/decks", a.CreateDeck).Methods(http.MethodPost)
	r.HandleFunc(getDeckWithId, a.GetDeck).Methods(http.MethodGet)
}
