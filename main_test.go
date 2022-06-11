package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var config = AppConfig{
		Addr: ":8081",
	}
	app := App{}
	app.initApp(app, config)

	code := m.Run()
	os.Exit(code)
}
