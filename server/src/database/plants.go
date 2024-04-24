package database

import (
	"fmt"
	"github.com/google/uuid"
)

type Plant struct {
	ID   uuid.UUID
	Name string
	Type string
}

func (d *Database) GetPlant(plantID uuid.UUID) (Plant, error) {
	var plant Plant
	if err := d.db.QueryRow(fmt.Sprintf("SELECT * FROM Plants p where p == %s", plantID)).Scan(&plant); err != nil {
		return Plant{}, err
	}
	return plant, nil

}
