package main

import (
	"fmt"
	// database "monitoring/databases"
	"monitoring/handlers"
	// "monitoring/models"
	"net/http"
)

func main() {
	// database.Connect()

	// services := []models.Service{
	// 	{"Google", "https://www.google.com", "GET"},
	// 	{"GO", "http://35.222.30.92:8080/users/1", "GET"},
	// }

	// go handlers.MonitorServices(services)

	// Set up HTTP server
	http.HandleFunc("/ws", handlers.HandleWebSocket)
	fmt.Println("WebSocket server running on ws://localhost:4444/ws")
	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}

	// select {} // Keeps the main function alive

}
