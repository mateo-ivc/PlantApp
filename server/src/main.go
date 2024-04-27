package main

import (
	"PlantApp/api"
	"PlantApp/database"
	"PlantApp/logger"
	"PlantApp/services/plants"
	"log"

	"net/http"
)

func main() {
	logger.EnableDevelopmentLogger()
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
func run() error {
	dbConnection := database.New()
	boardService := plants.NewPlantService(dbConnection)
	s := api.New(boardService)
	port := ":8080"
	logger.Get().Infow("starting server", "port", port)
	return http.ListenAndServe(port, s)
}
