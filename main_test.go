package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var router *mux.Router

func setup() {
	var config = AppConfig{
		Addr: ":8081",
	}
	app := App{}

	// TODO data store strategy
	_, r := app.initApp(app, config)
	router = r
}

func TestMain(m *testing.M) {
	setup()
	fmt.Println("===================| Before |===================")

	code := m.Run()

	fmt.Println("===================| After |===================")

	os.Exit(code)
}

func createRequest(verb string, route string, body io.Reader) (*http.Request, *httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(verb, route, body)
	if err != nil {
		fmt.Printf("Error creating request: [%v](%v)", verb, route)
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
	router.ServeHTTP(response, request)
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
	router.ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code)

	var data Deck
	json.Unmarshal(response.Body.Bytes(), &data)

	assert.NotNil(t, data.Id)
	assert.NotNil(t, data.Shuffled)
	assert.Equal(t, false, data.Shuffled)
	assert.NotNil(t, data.Remaining)
	assert.Equal(t, 52, data.Remaining)
	assert.Equal(t, 52, len(data.Cards))
}
