create table IF NOT EXISTS card_deck (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  card_id uuid NOT null REFERENCES cards (id),
  deck_id uuid NOT null REFERENCES decks (id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);