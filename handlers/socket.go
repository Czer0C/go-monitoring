package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

// HandleWebSocket handles WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer ws.Close()

	fmt.Println("Client connected")

	// Listen for messages from the client
	for {
		// Read raw message from the client
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		// Print the received message
		fmt.Printf("Received: %s\n", msg)

		// Echo the message back to the client
		err = ws.WriteMessage(messageType, []byte(fmt.Sprintf("You said: %s", msg)))
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
