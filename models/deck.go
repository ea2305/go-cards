package models

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/ea2305/go-cards/database"
)

type Deck struct {
	Id        string `json:"deck_id" db:"id"`
	Shuffled  bool   `json:"shuffled" db:"shuffled"`
	Remaining int    `json:"remaining" db:"remaining"`
	Cards     []Card `json:"cards" db:"card"`
	CreatedAt string `json:"-" db:"created_at"`
}

type Card struct {
	Id        string `json:"-" db:"id"`
	Value     string `json:"value" db:"value"`
	Suit      string `json:"suit" db:"suit"`
	Code      string `json:"code" db:"code"`
	CreatedAt string `json:"-" db:"created_at"`
}

func CreateDeck(shuffled bool, selection []string) (Deck, error) {
	rawCards, err := QueryAllCards()
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

	deck, storeErr := StoreDeck(shuffled, len(cards), cards)

	if storeErr != nil {
		// logs
		return deck, storeErr
	}

	if len(selection) > 0 {
		// logs
		return deck, errors.New("some cards we not found: " + fmt.Sprint(selection))
	} else {
		return deck, nil
	}
}

func GetDeck(id string) (Deck, error) {
	var deck Deck
	var queryDeck = "select * from decks where id = $1;"
	err := database.Connection.Get(&deck, queryDeck, id)
	if err != nil {
		// logs
		return Deck{}, errors.New("deck not found")
	}

	var cards []Card
	cardErr := database.Connection.Select(&cards, database.GetCardsByDeckId(), id)
	if cardErr != nil {
		// logs
		return Deck{}, errors.New("cards in deck missing")
	}

	deck.Cards = cards

	return deck, nil
}

func DrawCard(id string, count int) ([]Card, error) {
	var deck, err = GetDeck(id)
	if err != nil {
		return nil, err
	}

	if len(deck.Cards) < count {
		return nil, errors.New("not enough cards")
	}

	// get elements from the beginning of the list
	cards := deck.Cards[:count]

	tx := database.Connection.MustBegin()
	for _, card := range cards {
		tx.MustExec(database.DeleteCardInDeck(), deck.Id, card.Id)
	}

	tx.MustExec(database.UpdateRemainingCardsFromDeck(), deck.Remaining-count, deck.Id)
	commitErr := tx.Commit()
	if commitErr != nil {
		return nil, commitErr
	}

	return cards, nil
}

func QueryAllCards() ([]Card, error) {
	var cards []Card
	err := database.Connection.Select(&cards, "select * from cards;")
	if err != nil {
		// logs
		return nil, err
	}

	return cards, nil
}

func StoreDeck(shuffled bool, remaining int, cards []Card) (Deck, error) {
	tx := database.Connection.MustBegin()
	var deckId string
	tx.QueryRow(database.InsertDeck(), shuffled, remaining).Scan(&deckId)

	for _, card := range cards {
		tx.MustExec(database.InsertCardInDeck(), card.Id, deckId)
	}
	commitErr := tx.Commit()
	if commitErr != nil {
		return Deck{}, commitErr
	}

	var deck = Deck{
		Id:        deckId,
		Shuffled:  shuffled,
		Remaining: remaining,
		Cards:     cards,
	}

	return deck, nil
}
