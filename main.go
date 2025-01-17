package main

import (
	database "monitoring/databases"
	"monitoring/handlers"
	"monitoring/models"
)

func main() {

	database.Connect()

	services := []models.Service{
		{"Google", "https://www.google.com"},
		{"Example", "https://example.com"},
	}

	go handlers.MonitorServices(services)

	select {} // Keeps the main function alive
}
