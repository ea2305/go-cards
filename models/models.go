package models

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
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

var decks []Deck

func CreateDeck(shuffled bool, selection []string) (Deck, error) {
	rawCards := GetCards()
	var cards []Card

	if len(selection) > 0 {
		for _, card := range rawCards {
			for index, find := range selection {
				if card.Code == find {
					cards = append(cards, card)
					// removes element if the selection matches to validate the empty list later.
					selection = append(selection[:index], selection[index+1:]...)
				}
			}
		}
	} else {
		cards = rawCards
	}

	if shuffled {
		for i := range cards {
			j := rand.Intn(i + 1)
			cards[i], cards[j] = cards[j], cards[i]
		}
	}

	var deck = Deck{
		Id:        uuid.NewString(), // TODO implement uuid
		Shuffled:  shuffled,
		Remaining: len(cards),
		Cards:     cards,
	}

	// TODO provisional store strategy
	decks = append(decks, deck)

	if len(selection) > 0 {
		// logs
		return deck, errors.New("some cards we not found: " + fmt.Sprintf("%v != %v", deck.Cards, selection))
	} else {
		return deck, nil
	}
}

func GetCards() []Card {
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

	return cards
}

func GetDeck(id string) (Deck, error) {
	// uses in memory deck, TODO implement database strategy
	for _, deck := range decks {
		if deck.Id == id {
			return deck, nil
		}
	}
	return Deck{}, errors.New("deck not found")
}

func DrawCard(id string, count int) ([]Card, error) {
	// inmemory implementation
	var deck, err = GetDeck(id)
	if err != nil {
		return nil, err
	}

	if len(deck.Cards) < count {
		return nil, errors.New("not enough cards")
	}
	// get elements from the beginning of the list
	var index = 0
	var cards = make([]Card, count)
	var copyCards = make([]Card, len(deck.Cards))
	copy(copyCards, deck.Cards)
	cards = copySlice(copyCards, count)

	cardsUpdate := append(deck.Cards[:index], deck.Cards[index+count:]...)
	copy(deck.Cards, cardsUpdate)

	return cards, nil
}

func copySlice(slice []Card, count int) []Card {
	var cards []Card
	for index, card := range slice {
		cards = append(cards, card)
		if index == count-1 {
			return cards
		}
	}
	return cards
}
