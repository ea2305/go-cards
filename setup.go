package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	app.Server = server
	app.Router = router
	app.DB = db
}

func (app *App) migrateTables() {
	driver, err := postgres.WithInstance(app.DB.Unsafe().DB, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migration",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	m.Up()
}

func (app *App) rollbackTables() {
	driver, err := postgres.WithInstance(app.DB.Unsafe().DB, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migration",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	m.Down()
}
