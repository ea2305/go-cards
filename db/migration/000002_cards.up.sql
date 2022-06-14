create table IF NOT EXISTS cards (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  value VARCHAR(50) NOT NULL,
  suit VARCHAR(50) NOT NULL,
  code VARCHAR(50) NOT null,
  created_at timestamp
);