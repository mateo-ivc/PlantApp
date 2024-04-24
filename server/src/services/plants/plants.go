package plants

import (
	"PlantApp/common/dto"
	"PlantApp/database"
	"PlantApp/logger"
	"PlantApp/services"
	"context"
	"github.com/google/uuid"
)

type PlantService struct {
	database *database.Database
}

func (s *PlantService) Create(ctx context.Context, body dto.CreatePlantRequest) (*dto.Plant, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PlantService) Get(ctx context.Context, id uuid.UUID) (*dto.Plant, error) {
	plant, err := s.database.GetPlant(id)
	if err != nil {
		logger.Get().Errorw("Failed to fetch plant", "error", err)
		return nil, err
	}
	return new(dto.Plant).From(plant), nil
}

func (s *PlantService) Update(ctx context.Context, body dto.PlantUpdateRequest) (*dto.Plant, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PlantService) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *PlantService) New(db *database.Database) services.Plants {
	p := new(PlantService)
	p.database = db
	return p
}
