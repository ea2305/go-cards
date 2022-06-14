package factory

import (
	m "github.com/ea2305/go-cards/models"
	"github.com/google/uuid"
)

func FactoryCards() []m.Card {
	var cards []m.Card
	var suits = [4]string{"CLUBS", "DIAMONDS", "HEARTS", "SPADES"}
	var suit_codes = [4]string{"C", "D", "H", "S"}
	var values = [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			var suit = suits[i]
			var suitCode = suit_codes[i]
			var value = values[j]
			var code = suitCode + value

			var card = m.Card{
				Id:    uuid.NewString(),
				Value: value,
				Suit:  suit,
				Code:  code,
			}

			cards = append(cards, card)
		}
	}

	return cards
}
