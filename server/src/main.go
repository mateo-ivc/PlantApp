package main

import (
	"PlantApp/api"
	"PlantApp/database"
	"PlantApp/logger"
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
	_ = database.New()
	s := api.New()
	port := ":8080"
	//logger.Get().Infow("starting server", "base-path", basePath, "port", port)
	return http.ListenAndServe(port, s)
}
