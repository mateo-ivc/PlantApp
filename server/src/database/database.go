package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	db *sql.DB
}

const (
	host     = "185.202.239.207"
	port     = 5342
	user     = "mateo"
	password = "Mateo2002"
	dbname   = "postgres"
)

func New() *Database {

	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))
	if err != nil {
		return new(Database)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Database{
		conn,
	}
}
