package api

import (
	"PlantApp/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

func (s *Server) getPlants(w http.ResponseWriter, r *http.Request) {
	plants, err := s.plants.GetOverview(r.Context())
	if err != nil {
		logger.Get().Error("Failed to fetch plants", "error", err)
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, plants)
}

func (s *Server) getPlant(w http.ResponseWriter, r *http.Request) {
	plantID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	plant, err := s.plants.Get(r.Context(), plantID)
	if err != nil {
		logger.Get().Errorw("Failed to fetch plant", "plantID", plantID, "error", err)
	}

	render.Status(r, http.StatusOK)
	render.Respond(w, r, plant)
}
