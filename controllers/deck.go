package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	m "github.com/ea2305/go-cards/models"
	"github.com/gorilla/mux"
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

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	var status = map[string]string{"status": "ok"}
	responseJson(w, http.StatusOK, status)
}

func CreateDeck(w http.ResponseWriter, r *http.Request) {
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

	deck, err := m.CreateDeck(shuffled, selection)
	if err != nil {
		// logs
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseJson(w, http.StatusCreated, deck)
}

func GetDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		// logs
		responseError(w, http.StatusBadRequest, "missing id parameter")
		return
	}
	var deck, err = m.GetDeck(id)
	if err != nil {
		responseError(w, http.StatusNotFound, err.Error())
		return
	}
	responseJson(w, http.StatusOK, deck)
}

func DrawCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		// logs
		responseError(w, http.StatusBadRequest, "missing id parameter")
		return
	}
	var queryCount = r.URL.Query().Get("count")
	var count = 0
	if queryCount != "" {
		var parsed, err = strconv.Atoi(queryCount)
		if err != nil {
			// logs
			responseError(w, http.StatusBadRequest, "wrong count format")
			return
		}
		count = parsed
	}
	if count <= 0 {
		// logs
		responseError(w, http.StatusBadRequest, "count should be a positive integer greater than zero")
		return
	}

	var cards, err = m.DrawCard(id, count)
	if err != nil {
		// logs
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseJson(w, http.StatusOK, map[string][]m.Card{"cards": cards})
}
