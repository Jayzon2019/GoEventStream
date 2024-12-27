// client.go

package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	// Connect to the WebSocket server
	serverURL := "ws://localhost:3000/ws"
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatal("Dial failed:", err)
	}
	defer conn.Close()

	// Send a message to the server
	message := []byte("Hello, WebSocket server!")
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatal("WriteMessage failed:", err)
	}

	// Receive a response from the server
	_, response, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("ReadMessage failed:", err)
	}
	fmt.Printf("Received from server: %s\n", response)
}
