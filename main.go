package main

import (
	"log"
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
	app.migrateTables()

	log.Printf("Starting server at http://localhost%v \n", config.Addr)
	log.Fatal(app.Server.ListenAndServe())
}
