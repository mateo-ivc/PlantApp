package migration

import (
	"PlantApp/logger"
	"database/sql"
	"embed"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	host     = "localhost"
	port     = 5342
	user     = "mateo"
	password = "Mateo2002"
	dbname   = "postgres"
)

var Migrations embed.FS

func Open() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://mateo:Mateo2002@localhost:5432/PlantApp?sslmode=disable")
	if err != nil {
		logger.Get().Errorw("Could not connect to database", "error", err)
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Get().Errorw("driver nil, idk", "error", err)
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migration/sql",
		"postgres", driver)
	if err != nil {
		logger.Get().Errorw(" idk", "error", err)
		return nil, err
	}
	err = m.Up()
	if err == nil || errors.Is(err, migrate.ErrNoChange) {
		return db, nil

	}

	logger.Get().Errorw(" Migration dirty", "error", err)
	return nil, err
}
