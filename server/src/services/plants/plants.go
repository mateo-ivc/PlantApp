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

func NewPlantService(db *database.Database) services.Plants {
	p := new(PlantService)
	p.database = db
	return p
}

func (s *PlantService) Create(ctx context.Context, body dto.CreatePlantRequest) (*dto.Plant, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PlantService) Get(ctx context.Context, id int) (*dto.Plant, error) {
	plant, err := s.database.GetPlant(id)
	if err != nil {
		logger.Get().Errorw("Failed to fetch plant", "error", err)
		return nil, err
	}
	return new(dto.Plant).From(plant), nil
}
func (s *PlantService) GetOverview(ctx context.Context) ([]*dto.Plant, error) {

	plants, err := s.database.GetOverview()
	plantOverview := make([]*dto.Plant, len(plants))
	for index, plant := range plants {
		plantOverview[index] = new(dto.Plant).From(plant)
	}
	if err != nil {
		logger.Get().Errorw("Failed to fetch plant", "error", err)
		return nil, err
	}
	return plantOverview, nil
}

func (s *PlantService) UpdatePlant(ctx context.Context, body dto.PlantUpdateRequest) (*dto.Plant, error) {
	//TODO implement me
	update := database.PlantUpdate{
		ID:          body.ID,
		Name:        body.Name,
		Temperature: body.Temperature,
		Moisture:    body.Moisture,
		Humidity:    body.Humidity,
		Lighting:    body.Lighting,
	}
	plant, err := s.database.UpdatePlant(update)
	if err != nil {
		return nil, err
	}
	return new(dto.Plant).From(plant), err
}

func (s *PlantService) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
