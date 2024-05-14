package api

import (
	"PlantApp/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// define structs here
type Server struct {
	plants services.Plants
}

func New(plants services.Plants) chi.Router {
	r := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	r.Use(c.Handler)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	s := Server{
		plants: plants,
	}
	s.publicRoutes(r)
	return r
}

func (s *Server) publicRoutes(r chi.Router) chi.Router {
	return r.Group(func(r chi.Router) {
		r.Get("/plants", s.getPlants)
		r.Get("/plants/{id}", s.getPlant)
		r.Get("/plants/{id}/image", s.servePlantImage)
		r.Post("/plants/{id}/image", s.uploadPlantImage)
	})
}
