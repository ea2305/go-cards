package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/ea2305/go-gards/main"
)

func main() {
	r := mux.NewRouter()
	var addr = ":8080"
	app := App{}
	app.initRoutes(r)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	fmt.Printf("[Starting server in localhost%v...]", addr)
	server.ListenAndServe()
}
