package main

import (
	"fmt"
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
	fmt.Println("[Starting test]")

	code := m.Run()

	fmt.Println("[Ending test]")
	os.Exit(code)
}

func TestHealthCheckResponse(t *testing.T) {
	request, _ := http.NewRequest("GET", "/health_check", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
