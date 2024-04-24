package dto

import (
	"PlantApp/database"
	"github.com/google/uuid"
)

type Plant struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Type string    `json:"type,omitempty"`
}
type CreatePlantRequest struct {
	Name string `json:"name"`
}

type PlantUpdateRequest struct {
	Name string `json:"name"`
}

func (b *Plant) From(plant database.Plant) *Plant {
	b.ID = plant.ID
	b.Name = plant.Name
	b.Type = plant.Type
	return b
}
