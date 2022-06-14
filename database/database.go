package database

import "github.com/jmoiron/sqlx"

var Connection *sqlx.DB

func SetConnection(connection *sqlx.DB) {
	Connection = connection
}
