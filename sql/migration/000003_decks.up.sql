create table IF NOT EXISTS decks (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  shuffled BOOLEAN NOT NULL,
  remaining INT NOT null,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);