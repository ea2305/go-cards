package controllers

import (
	"encoding/json"
	"net/http"
)

var cards = []string{"card1", "card2"}

func CreateDeck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(cards)
}
