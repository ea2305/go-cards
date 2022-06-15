package database

func GetCardsByDeckId() string {
	query := `
		select
			cards.id as id,
			cards.value as value,
			cards.suit as suit,
			cards.code as code
		from 
			card_deck card_deck 
		left join 
			cards as cards
		on 
			card_deck.card_id = cards.id 
		where 
			card_deck.deck_id = $1
		order by 
			card_deck.created_at asc;
	`
	return query
}

func DeleteCardInDeck() string {
	query := `
	delete 
		from card_deck 
	where 
		deck_id = $1 and card_id = $2
	`
	return query
}

func UpdateRemainingCardsFromDeck() string {
	query := `
		update decks set remaining = $1 
		where 
			id = $2
	`
	return query
}

func InsertDeck() string {
	query := `
		insert into decks (shuffled, remaining) values ($1, $2) RETURNING id;
	`
	return query
}

func InsertCardInDeck() string {
	query := `
		insert into card_deck (card_id, deck_id) values ($1, $2);
	`
	return query
}
