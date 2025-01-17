package main

import (
	"fmt"
	"net/http"
	"time"
)

type Service struct {
	Name string
	URL  string
}

const INTERVAL = 10 * time.Second

func monitorService(service Service, interval time.Duration) {
	for {
		resp, err := http.Get(service.URL)

		fmt.Printf("Checking service %s\n", resp)

		if err != nil || resp.StatusCode != http.StatusOK {
			fmt.Printf("Service %s is DOWN! Error: %v\n", service.Name, err)
		} else {
			fmt.Printf("Service %s is UP. Status: %s\n", service.Name, resp.Status)
		}

		time.Sleep(interval)
	}
}

func main() {
	services := []Service{
		{"Google", "https://www.google.com"},
		{"Example", "https://example.com"},
	}

	for _, service := range services {
		go monitorService(service, INTERVAL)
	}

	select {} // Keeps the main function alive
}
