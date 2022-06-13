package main

import (
	"math/rand"
	"strconv"
)

type Deck struct {
	Id        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []Card `json:"cards"`
}

type Card struct {
	Id    string `json:"card_id"`
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

func GetDeck(shuffled bool) Deck {
	cards := GetCards(shuffled)

	var deck = Deck{
		Id:        "uudid", // TODO implement uuid
		Shuffled:  shuffled,
		Remaining: len(cards),
		Cards:     cards,
	}
	return deck
}

func GetCards(shuffled bool) []Card {
	var cards []Card
	var suits = [4]string{"CLUBS", "DIAMONDS", "HEARTS", "SPADES"}
	var suit_codes = [4]string{"C", "D", "H", "S"}
	var values = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			var suit = suits[i]
			var suitCode = suit_codes[i]
			var value = values[j]
			var code = suitCode + value

			var card = Card{
				Id:    suit + strconv.Itoa(j), // TODO: implements uuid generator
				Value: value,
				Suit:  suit,
				Code:  code,
			}

			cards = append(cards, card)
		}
	}

	if shuffled {
		for i := range cards {
			j := rand.Intn(i + 1)
			cards[i], cards[j] = cards[j], cards[i]
		}
	}

	return cards
}
