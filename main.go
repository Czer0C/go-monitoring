package main

import (
	database "monitoring/databases"
	"monitoring/handlers"
	"monitoring/models"
)

func main() {

	database.Connect()

	services := []models.Service{
		{"Google", "https://www.google.com", "GET"},
		{"GO", "http://35.222.30.92:8080/users/1", "GET"},
	}

	go handlers.MonitorServices(services)

	select {} // Keeps the main function alive
}
