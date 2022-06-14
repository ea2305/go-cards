package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type App struct {
	Server *http.Server
	Router *mux.Router
	DB     *sqlx.DB
}

type AppConfig struct {
	Addr     string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	driver string
	user   string
	pass   string
	name   string
	host   string
	port   string
	ssl    string
}

func (app *App) initApp(config AppConfig) {
	router := mux.NewRouter()
	app.initRoutes(router)

	server := &http.Server{
		Addr:    config.Addr,
		Handler: router,
	}

	var connectionString = fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.Database.user,
		config.Database.pass,
		config.Database.host,
		config.Database.port,
		config.Database.name,
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}

	migrateTables(db)

	app.Server = server
	app.Router = router
	app.DB = db
}

func migrateTables(db *sqlx.DB) {
	const migrations = `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

		create table IF NOT EXISTS decks (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			shuffled BOOLEAN NOT NULL,
			remaining INT NOT null,
			created_at timestamp
		);
		
		create table IF NOT EXISTS cards (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			value VARCHAR(50) NOT NULL,
			suit VARCHAR(50) NOT NULL,
			code VARCHAR(50) NOT null,
			created_at timestamp
		);
		
		create table IF NOT EXISTS card_deck (
			id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
			card_id uuid NOT null REFERENCES cards (id),
				deck_id uuid NOT null REFERENCES decks (id),
				created_at timestamp
		);
	`

	if _, err := db.Exec(migrations); err != nil {
		log.Fatal(err)
	}
}
