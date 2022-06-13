package main

import (
	"encoding/json"
	"net/http"
)

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	var status = map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}

func (a *App) CreateDeck(w http.ResponseWriter, r *http.Request) {
	model := GetDeck()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	// w.Write(response)
	if err := json.NewEncoder(w).Encode(model); err != nil {
		http.Error(w, "bad request", 400)
	}
}
