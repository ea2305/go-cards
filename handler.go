package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func responseError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func responseJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json, err := json.Marshal(payload)
	if err != nil {
		// logs
		responseError(w, http.StatusInternalServerError, "response error")
		return
	}
	w.Write(json)
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	var status = map[string]string{"status": "ok"}
	responseJson(w, http.StatusOK, status)
}

func (a *App) CreateDeck(w http.ResponseWriter, r *http.Request) {
	var queryShuffle = r.URL.Query().Get("shuffled")
	var queryCards = r.URL.Query().Get("cards")
	var shuffled = false
	var selection []string
	if queryCards != "" {
		var split = strings.Split(queryCards, ",")
		selection = split
	}
	if queryShuffle != "" {
		var parsed, err = strconv.ParseBool(queryShuffle)
		if err != nil {
			// logs
			responseError(w, http.StatusBadRequest, "bad request")
			return
		}
		shuffled = parsed
	}

	deck, err := GetDeck(shuffled, selection)
	if err != nil {
		// logs
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseJson(w, http.StatusCreated, deck)
}
