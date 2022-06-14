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
