package models

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/ea2305/go-cards/database"
	"github.com/google/uuid"
)

type Deck struct {
	Id        string `json:"deck_id" db:"id"`
	Shuffled  bool   `json:"shuffled" db:"shuffled"`
	Remaining int    `json:"remaining" db:"remaining"`
	Cards     []Card `json:"cards"`
}

type Card struct {
	Id        string `json:"-" db:"id"`
	Value     string `json:"value" db:"value"`
	Suit      string `json:"suit" db:"suit"`
	Code      string `json:"code" db:"code"`
	CreatedAt string `json:"-" db:"created_at"`
}

// TODO remove when database is in place

var decks []Deck

func CreateDeck(shuffled bool, selection []string) (Deck, error) {
	rawCards, err := GetCards()
	if err != nil {
		return Deck{}, err
	}

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
		Id:        uuid.NewString(),
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

func GetCards() ([]Card, error) {
	var cards []Card
	err := database.Connection.Select(&cards, "select * from cards;")
	if err != nil {
		// logs
		return nil, err
	}

	return cards, nil
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
