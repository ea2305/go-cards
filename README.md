# Go Cards API

This Repository implements methods to create, open and draw cards from one deck of BlackJack game. All the methods are accessible through the API. 

### Requirements
- Docker
- Docker Compose

### API reference

|method|verb|endpoint|body|query-params|constrains|
|---|---|---|---|---|---|
|create deck|POST|/api/v1/decks|{}|shuffled boolean,cards []card_code| Selection list length should be more than 0, and s I has to be separated with semicolons. Shuffled is optional, if the query param is not provided the API will return the default order (card's number: A-K, suits: C,D,H,S)
|open deck|GET|/api/v1/decks/{id:uuid}|{}|{}|_ID_ is mandatory param, if the param is not provide the API will return an error. The ID should have the _UUID_ format|
|draw card|PATCH|/api/v1/decks/{id:uuid}|{}|count int|count should be greater than zero. If the count exceeds the remaining amount of cards will return an error|

### Get started
 
> Install and run project locally with docker

- `./scripts/start.sh`

> Run integration test with docker

- `./scripts/test.sh`

### Migrations

_Setup_

In order to create new migrations install the following dependency:

`brew install golang-migrate`

_Create new migration_

`migrate create -ext sql -dir db/migration -seq table_names`

> Migrations will be execute automatically at the init phase of the server