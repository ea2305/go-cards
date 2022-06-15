package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ea2305/go-cards/database"
	"github.com/ea2305/go-cards/middlewares"
	"github.com/ea2305/go-cards/routes"
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
	routes.InitRoutes(router)

	router.Use(middlewares.LoggingMiddleware)

	server := &http.Server{
		Addr:         config.Addr,
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
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

	database.SetConnection(db)
	app.Server = server
	app.Router = router
}

func (app *App) migrateTables() {
	driver, err := postgres.WithInstance(database.Connection.Unsafe().DB, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://sql/migration",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	m.Up()
}

func (app *App) rollbackTables() {
	driver, err := postgres.WithInstance(database.Connection.Unsafe().DB, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://sql/migration",
		"postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	m.Down()
}
