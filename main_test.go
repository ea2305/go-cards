package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	m "github.com/ea2305/go-cards/models"
	"github.com/stretchr/testify/assert"
)

var app = App{}

func setup() {
	var config = AppConfig{
		Addr: ":8081",
		Database: DatabaseConfig{
			driver: "postgres",
			user:   os.Getenv("DB_USER"),
			pass:   os.Getenv("DB_PASS"),
			name:   os.Getenv("DB_NAME"),
			host:   os.Getenv("DB_HOST"),
			port:   os.Getenv("DB_PORT"),
			ssl:    os.Getenv("DB_SSL"),
		},
	}

	app.initApp(config)
}

func TestMain(m *testing.M) {
	setup()
	fmt.Println("===================| Before |===================")
	app.migrateTables()

	code := m.Run()

	fmt.Println("===================| After |===================")
	app.rollbackTables()

	os.Exit(code)
}

func createRequest(verb string, route string, body io.Reader) (*http.Request, *httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(verb, route, body)
	if err != nil {
		log.Printf("Error creating request: [%v](%v)\n", verb, route)
		return nil, nil, err
	}
	res := httptest.NewRecorder()
	return req, res, nil
}

func TestHealthCheckResponse(t *testing.T) {
	var request, response, err = createRequest("GET", "/api/v1/health_check", nil)
	if err != nil {
		t.Fail()
	}
	app.Router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)

	var data map[string]string
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.Equal(t, "ok", data["status"])
}

func TestCreateDeck(t *testing.T) {
	var request, response, err = createRequest("POST", "/api/v1/decks", nil)
	if err != nil {
		t.Fail()
	}
	app.Router.ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code)

	var data m.Deck
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.NotNil(t, data.Id)
	assert.NotNil(t, data.Shuffled)
	assert.Equal(t, false, data.Shuffled)
	assert.NotNil(t, data.Remaining)
	assert.Equal(t, 52, data.Remaining)
}

func TestCreateShuffledDeck(t *testing.T) {
	var request, response, err = createRequest("POST", "/api/v1/decks?shuffled=true", nil)
	if err != nil {
		t.Fail()
	}
	app.Router.ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code)

	var data m.Deck
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.NotNil(t, data.Shuffled)
	assert.Equal(t, true, data.Shuffled)
	assert.NotNil(t, data.Remaining)
	assert.Equal(t, 52, data.Remaining)
}

func TestCreateCustomDeck(t *testing.T) {
	var reqCards = []string{"CQ", "DJ", "H7", "H8"}
	var reqCardStr = strings.Join(reqCards[:], ",")
	var request, response, err = createRequest("POST", "/api/v1/decks?cards="+reqCardStr, nil)
	if err != nil {
		t.Fail()
	}
	app.Router.ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code)

	var data m.Deck
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.NotNil(t, data.Remaining)
	assert.Equal(t, len(reqCards), data.Remaining)
}

func TestCreateCustomDeckWithWrongFormat(t *testing.T) {
	var reqCards = []string{"CQ", "DJ", "ZY", "ZX"}
	var reqCardStr = strings.Join(reqCards[:], ",")
	var request, response, err = createRequest("POST", "/api/v1/decks?cards="+reqCardStr, nil)
	if err != nil {
		t.Fail()
	}
	app.Router.ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code)

	var data map[string]string
	json.Unmarshal(response.Body.Bytes(), &data)
	assert.NotNil(t, data["error"])
	assert.Contains(t, data["error"], "some cards we not found:")
}

func TestCreateCustomShuffledDeck(t *testing.T) {
	var reqCards = []string{"CQ", "DJ", "H7", "H8"}
	var reqCardStr = strings.Join(reqCards[:], ",")
	var request, response, err = createRequest("POST", "/api/v1/decks?shuffled=true&cards="+reqCardStr, nil)
	if err != nil {
		t.Fail()
	}
	app.Router.ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code)

	var data m.Deck
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.NotNil(t, data.Shuffled)
	assert.NotNil(t, data.Remaining)
	assert.Equal(t, true, data.Shuffled)
	assert.Equal(t, len(reqCards), data.Remaining)
}

func TestGetDeck(t *testing.T) {
	var deck, _ = m.CreateDeck(false, nil)
	var id = deck.Id
	var request, response, err = createRequest("GET", "/api/v1/decks/"+id, nil)
	if err != nil {
		t.Fail()
	}
	app.Router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)

	var data m.Deck
	json.Unmarshal(response.Body.Bytes(), &data)
	assert.NotNil(t, data.Id)
	assert.Equal(t, data.Id, id)
	assert.NotNil(t, data.Cards)
	assert.Equal(t, len(deck.Cards), len(data.Cards))
}

func TestGetDeckNotFound(t *testing.T) {
	var request, response, err = createRequest("GET", "/api/v1/decks/3871bbef-2736-4416-b04f-d7bfb51b75a2", nil)
	if err != nil {
		t.Fail()
	}
	app.Router.ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code)

	var data map[string]string
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.NotNil(t, data["error"])
	assert.Contains(t, data["error"], "deck not found")
}

func TestGetDeckWrongUuidFormat(t *testing.T) {
	var request, response, err = createRequest("GET", "/api/v1/decks/3871b75a2", nil)
	if err != nil {
		t.Fail()
	}
	app.Router.ServeHTTP(response, request)
	assert.Equal(t, 404, response.Code)
}

func TestDrawCards(t *testing.T) {
	var deck, _ = m.CreateDeck(false, nil)
	var fistCardCode = deck.Cards[0].Code
	var id = deck.Id
	var request, response, err = createRequest("PATCH", fmt.Sprintf("/api/v1/decks/%v?count=1", id), nil)
	if err != nil {
		t.Fail()
	}
	app.Router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)

	var data map[string][]m.Card
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.Equal(t, 1, len(data))
	if len(data) > 0 {
		// pick elements from the beginning
		assert.Equal(t, fistCardCode, data["cards"][0].Code)
	}
}

func TestDrawCardsWithGreaterCount(t *testing.T) {
	var deck, _ = m.CreateDeck(false, nil)
	var id = deck.Id
	var request, response, err = createRequest("PATCH", fmt.Sprintf("/api/v1/decks/%v?count=54", id), nil)
	if err != nil {
		t.Fail()
	}
	app.Router.ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code)

	var data map[string]string
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.Contains(t, data["error"], "not enough cards")
}
