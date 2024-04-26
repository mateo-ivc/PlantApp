package database

import (
	"PlantApp/database/migration"
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

const (
	host     = "localhost"
	port     = 5342
	user     = "mateo"
	password = "Mateo2002"
	dbname   = "postgres"
)

func New() *Database {
	conn, err := migration.Open()
	if err != nil {
		return new(Database)
	}

	return &Database{
		conn,
	}
}
