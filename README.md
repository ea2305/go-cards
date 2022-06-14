# Go Cards API

This Repository implements methods to create, open and draw cards from one deck of BlackJack game. All the methods are accessible through the API. 

### Requirements
- Docker
- Docker Compose

### API reference

|method|verb|endpoint|body|query-params|
|---|---|---|---|---|
|create deck|POST|/api/v1/decks|none|shuffled(boolean),cards([]card_code)|
|open deck|GET|/api/v1/decks/{id:uuid}|none|none|
|draw card|PATCH|/api/v1/decks/{id:uuid}|none|count(int>0)|

### Get started
 
> Install and run project locally with docker

- `./scripts/start.sh`

> Run integration test with docker

- `./scripts/test.sh`