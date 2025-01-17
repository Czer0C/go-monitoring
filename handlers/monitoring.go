package handlers

import (
	"fmt"
	database "monitoring/databases"
	"monitoring/models"
	"net/http"
	"time"
)

const INTERVAL = 10 * time.Second

func monitorService(service models.Service) {

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	start := time.Now()
	resp, err := client.Get(service.URL)
	latency := time.Since(start)

	var statusCode string
	if err != nil {
		alertError(service, err, start.Local().Format("2006-01-02 15:04:05"))
		statusCode = "0"
	} else {
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			alertError(service, fmt.Errorf("unexpected status code: %d", resp.StatusCode), start.Local().Format("2006-01-02 15:04:05"))
		}
		statusCode = fmt.Sprintf("%d", resp.StatusCode)
	}

	database.InsertServiceStatus(service.Name, service.URL, statusCode, latency.String(), start)

}

func MonitorServices(services []models.Service) {

	for {
		for _, service := range services {
			go monitorService(service)
		}

		time.Sleep(INTERVAL)
	}

}
