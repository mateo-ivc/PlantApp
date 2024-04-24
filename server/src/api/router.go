package api

import (
	"PlantApp/services"
	"github.com/go-chi/chi/v5"
)

// define structs here
type Server struct {
	plants services.Plants
}

func New() chi.Router {
	r := chi.NewRouter()
	s := Server{}
	s.publicRoutes(r)
	return r
}

func (s *Server) publicRoutes(r chi.Router) chi.Router {
	return r.Group(func(r chi.Router) {
		r.Get("/", s.getPlants)
	})
}
