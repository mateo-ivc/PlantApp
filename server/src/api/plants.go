package api

import (
	"PlantApp/logger"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func (s *Server) getPlants(w http.ResponseWriter, r *http.Request) {

	plantID := r.Context().Value("Plant").(uuid.UUID)

	plant, err := s.plants.Get(r.Context(), plantID)
	if err != nil {
		logger.Get().Error("Failed to fetch plants", "error", err)
	}
	fmt.Println(plant)
}
