package main

import (
	"fmt"
	"os"
)

func main() {
	app := App{}
	var config = AppConfig{
		Addr: ":8080",
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

	fmt.Printf("[Starting server in localhost%v...]", config.Addr)
	app.Server.ListenAndServe()
}
