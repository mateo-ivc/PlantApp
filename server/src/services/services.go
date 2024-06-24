package services

import (
	"PlantApp/common/dto"
	"context"
	"github.com/google/uuid"
)

type Plants interface {
	Create(ctx context.Context, body dto.CreatePlantRequest) (*dto.Plant, error)
	Get(ctx context.Context, id int) (*dto.Plant, error)
	GetOverview(ctx context.Context) ([]*dto.Plant, error)
	UpdatePlant(ctx context.Context, body dto.PlantUpdateRequest) (*dto.Plant, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
