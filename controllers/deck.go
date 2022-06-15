package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	m "github.com/ea2305/go-cards/models"
	"github.com/gorilla/mux"
)

type DeckResponse struct {
	Id        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
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
			log.Printf("[controller:create-deck] shuffled query serialization fails \n")
			responseError(w, http.StatusBadRequest, "bad request")
			return
		}
		shuffled = parsed
	}

	deck, err := m.CreateDeck(shuffled, selection)
	if err != nil {
		log.Printf("[controller:create-deck] create deck failed \n")
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	customResponse := DeckResponse{
		Id:        deck.Id,
		Shuffled:  deck.Shuffled,
		Remaining: deck.Remaining,
	}
	responseJson(w, http.StatusCreated, customResponse)
}

func GetDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Printf("[controller:get-deck] param id was not defined \n")
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
		log.Printf("[controller:draw-card] param id was not defined \n")
		responseError(w, http.StatusBadRequest, "missing id parameter")
		return
	}
	var queryCount = r.URL.Query().Get("count")
	var count = 0
	if queryCount != "" {
		var parsed, err = strconv.Atoi(queryCount)
		if err != nil {
			log.Printf("[controller:draw-card] count parse error \n")
			responseError(w, http.StatusBadRequest, "wrong count format")
			return
		}
		count = parsed
	}
	if count <= 0 {
		log.Printf("[controller:draw-card] count is less equal than zero \n")
		responseError(w, http.StatusBadRequest, "count should be a positive integer greater than zero")
		return
	}

	var cards, err = m.DrawCard(id, count)
	if err != nil {
		log.Printf("[controller:draw-card] drawn cards from model fails \n")
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseJson(w, http.StatusOK, map[string][]m.Card{"cards": cards})
}
