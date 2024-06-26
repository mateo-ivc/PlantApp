package dto

import (
	"PlantApp/database"
	"time"
)

type Plant struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Temperature  int         `json:"temperature"`
	Moisture     int         `json:"moisture"`
	Humidity     int         `json:"humidity"`
	Lighting     int         `json:"lighting"`
	CreatedAt    time.Time   `json:"created_at"`
	ControllerID interface{} `json:"controllerID"`
}
type CreatePlantRequest struct {
	Name string `json:"name"`
}

type PlantUpdateRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Temperature int    `json:"temperature"`
	Moisture    int    `json:"moisture"`
	Humidity    int    `json:"humidity"`
	Lighting    int    `json:"lighting"`
}

func (p *Plant) From(plant database.Plant) *Plant {
	p.ID = plant.ID
	p.Name = plant.Name
	p.Temperature = plant.Temperature
	p.Moisture = plant.Moisture
	p.Humidity = plant.Humidity
	p.Lighting = plant.Lighting
	p.CreatedAt = plant.CreatedAt
	p.ControllerID = plant.ControllerID
	return p
}
