package services

import (
	"PlantApp/common/dto"
	"context"
	"github.com/google/uuid"
)

type Plants interface {
	Create(ctx context.Context, body dto.CreatePlantRequest) (*dto.Plant, error)
	Get(ctx context.Context, id uuid.UUID) (*dto.Plant, error)
	Update(ctx context.Context, body dto.PlantUpdateRequest) (*dto.Plant, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
