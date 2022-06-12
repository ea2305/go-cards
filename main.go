package main

import (
	"fmt"
)

func main() {
	app := App{}
	var config = AppConfig{
		Addr: ":8080",
	}

	server, _ := app.initApp(app, config)

	fmt.Printf("[Starting server in localhost%v...]", config.Addr)
	server.ListenAndServe()
}
