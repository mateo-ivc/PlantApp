package database

import (
	"fmt"
	"time"
)

type Plant struct {
	ID           int
	Name         string
	Moisture     int
	Humidity     int
	Lighting     int
	CreatedAt    time.Time
	ControllerID interface{}
}

func (d *Database) GetPlant(plantID int) (Plant, error) {
	var plant Plant
	if err := d.db.QueryRow(fmt.Sprintf("SELECT * FROM Plants p where p.id = %d", plantID)).
		Scan(&plant.ID, &plant.Name, &plant.Moisture, &plant.Humidity, &plant.Lighting, &plant.CreatedAt, &plant.ControllerID); err != nil {
		return Plant{}, err
	}
	return plant, nil

}

func (d *Database) GetOverview() ([]Plant, error) {
	rows, err := d.db.Query("SELECT * FROM plants")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var plants []Plant
	for rows.Next() {
		var plant Plant
		// Scan the values from the row into the struct fields
		err := rows.Scan(&plant.ID, &plant.Name, &plant.Moisture, &plant.Humidity, &plant.Lighting, &plant.CreatedAt, &plant.ControllerID)
		if err != nil {
			panic(err)
		}
		// Append the plant to the slice
		plants = append(plants, plant)
	}
	return plants, nil

}
