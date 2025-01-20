package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"monitoring/models"
	"net/http"
)

type DiscordWebhook struct {
	Content string `json:"content"`
}

const discordhook = "https://discord.com/api/webhooks/1329684257051775042/wEcFbIj5N_jmmTiXBuoX6GX9qcIdKCXW4OtEN2xXjH3jnEVAGlZsqaK4aSRBw52Ji0hR"

func alertError(service models.Service, err error, time string) {

	message := fmt.Sprintf("❌ Service %s is **DOWN!** Error: %v at %s\n", service.Name, err, time)

	//!Hook discord

	payload := DiscordWebhook{Content: message}

	// Marshal the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		// return fmt.Errorf("error marshalling payload: %v", err)
	}

	// Send the POST request to Discord
	req, err := http.NewRequest("POST", discordhook, bytes.NewBuffer(payloadBytes))
	if err != nil {
		// return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		// return fmt.Errorf("error: received non-200 response code: %v", resp.StatusCode)
	}

}
