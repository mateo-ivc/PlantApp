package database

import (
	"fmt"
	"time"
)

type Plant struct {
	ID           int
	Name         string
	Temperature  int
	Moisture     int
	Humidity     int
	Lighting     int
	CreatedAt    time.Time
	ControllerID interface{}
}

type PlantUpdate struct {
	ID          int
	Name        string
	Temperature int
	Moisture    int
	Humidity    int
	Lighting    int
}

func (d *Database) GetPlant(plantID int) (Plant, error) {
	var plant Plant
	if err := d.db.QueryRow(fmt.Sprintf("SELECT * FROM plants p where p.id = %d", plantID)).
		Scan(&plant.ID, &plant.Name, &plant.Temperature, &plant.Moisture, &plant.Humidity, &plant.Lighting, &plant.CreatedAt, &plant.ControllerID); err != nil {
		return Plant{}, err
	}
	return plant, nil

}

func (d *Database) UpdatePlant(plant PlantUpdate) (Plant, error) {

	_, err := d.db.Exec("UPDATE plants SET name = $1, temperature = $2, moisture = $3, humidity = $4, lighting = $5 WHERE id = $6", plant.Name, plant.Temperature, plant.Moisture, plant.Humidity, plant.Lighting, plant.ID)
	if err != nil {
		return Plant{}, err
	}
	updatedPlant, _ := d.GetPlant(plant.ID)
	return updatedPlant, err
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
		err := rows.Scan(&plant.ID, &plant.Name, &plant.Temperature, &plant.Moisture, &plant.Humidity, &plant.Lighting, &plant.CreatedAt, &plant.ControllerID)
		if err != nil {
			panic(err)
		}
		// Append the plant to the slice
		plants = append(plants, plant)
	}
	return plants, nil

}
