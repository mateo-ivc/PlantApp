package api

import (
	"PlantApp/logger"
	"encoding/base64"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"os"
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

func (s *Server) uploadPlantImage(w http.ResponseWriter, r *http.Request) {
	plantID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	contentType := r.Header.Get("Content-type")

	if contentType == "image/jpg" || contentType == "image/png" {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Get().Errorw("Failed read Body")
		}

		err = os.WriteFile(fmt.Sprintf("static/images/plant_%d.png", plantID), data, 0666)
		if err != nil {
			logger.Get().Errorw("Failed to save file")
			return
		}

		logger.Get().Debugln("Image received and saved!")

		render.Status(r, http.StatusOK)
		render.Respond(w, r, nil)
	} else {
		render.Status(r, http.StatusForbidden)
	}
}

func (s *Server) servePlantImage(w http.ResponseWriter, r *http.Request) {
	plantID, _ := strconv.Atoi(chi.URLParam(r, "id"))

	file, err := os.ReadFile(fmt.Sprintf("static/images/plant_%d.png", plantID))
	if err != nil {
		logger.Get().Errorw("Failed to read file", "error", err, "plantID", plantID)
	}
	base64Img := base64.StdEncoding.EncodeToString(file)

	// Write the base64 encoded image data to the response
	w.Header().Set("Content-Type", "text/plain") // Set the appropriate content type
	w.Write([]byte(base64Img))
}
