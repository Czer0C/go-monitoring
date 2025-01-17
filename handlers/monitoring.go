package handlers

import (
	"fmt"
	database "monitoring/databases"
	"monitoring/models"
	"net/http"
	"time"
)

const INTERVAL = 10 * time.Second

func monitorService(service models.Service, interval time.Duration) {
	for {
		start := time.Now()

		resp, err := http.Get(service.URL)

		latency := time.Since(start)

		if err != nil || resp.StatusCode != http.StatusOK {
			// fmt.Printf("Service %s is DOWN! Error: %v\n", service.Name, err)
			alertError(service, err, start.Local().Format("2006-01-02 15:04:05"))

		} else {
			// fmt.Printf("Service %s is UP. Status: %s - in %s\n", service.Name, resp.Status, latency)
		}

		strCode := fmt.Sprintf("%d", resp.StatusCode)

		database.InsertServiceStatus(service.Name, service.URL, strCode, latency.String(), start)

		time.Sleep(interval)
	}
}

func MonitorServices(services []models.Service) {
	for _, service := range services {
		go monitorService(service, INTERVAL)
	}
}
