package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	// "github.com/gorilla/websocket"
	//"golang.org/x/net/websocket"
	//"nhooyr.io/websocket"

	// "nhooyr.io/websocket"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections
	},
}

func main() {
	flag.Parse()

	http.HandleFunc("/", home)
	http.HandleFunc("/events", events)
	http.HandleFunc("/eventStream", eventStream)
	// Handle the /events route which streams server time
	http.HandleFunc("/eventStreamHandler", eventStreamHandler)
	//http.HandleFunc("/wsHandler", wsHandler)

	http.HandleFunc("/ws", handleConnection)

	http.ListenAndServe(":3000", nil)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		// Log the message received from client
		fmt.Printf("Received: %s\n", p)

		// Send a message back to the client
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}

// func wsHandler(w http.ResponseWriter, r *http.Request) {
// 	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions,{

// 	 	CompressionMode: websocket.CompressionDisabled,
// 	})

// 	if err != nil {
// 		http.Error(w, "Error Accepting webSocket connection", 400)
// 		return
// 	}
// 	defer conn.Close(websocket.StatusNormalClosure, "closed")

// }

// Event stream handler that sends the current server time every second
func eventStreamHandler(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE (Server-Sent Events)
	w.Header().Set("Content-Type", "text/event-stream")
	// w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	currentTime := time.Now().Format(time.RFC1123)

	// Infinite loop to send time every second
	for {
		// Get the current time
		//currentTime = time.Now().Format(time.RFC1123)

		currentTime = time.Now().Format("Monday, January 2, 2006 03:04:05 PM")

		//formattedTime := currentTime.Format("03:04:05 PM")

		content := fmt.Sprintf("data: %s\n\n", currentTime)
		//content := fmt.Sprintf("data: %s\n\n", formattedTime)
		// Send the current time as an SSE event
		// w.Write([]byte(content))

		// Flush the response to ensure the message is sent immediately

		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		// flusher, ok := w.(http.Flusher)
		// if ok {
		// 	flusher.Flush()
		// } else {

		// }

		// Sleep for 1 second before sending the next time update
		time.Sleep(1 * time.Second)

	}
	// Close the connection
	w.WriteHeader(http.StatusOK)

}

// Event handler that will send SSE events to the browser
func eventStream(w http.ResponseWriter, r *http.Request) {
	// Set appropriate headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	// w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Connection", "keep-alive")
	//w.Header().Set("Access-Control-Allow-Origin", "172.29.80.1")
	// Allow all domains
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Set the maximum number of events to send
	for i := 1; i <= 10; i++ {
		// Write the event data
		//fmt.Fprintf(w, "data: Event %d at %s\n\n", i, time.Now().Format("2006-01-02 15:04:05"))
		//w.Write([]byte(content))
		//w.(http.Flusher).Flush()

		content := fmt.Sprintf("data: Event %d at %s\n\n", i, time.Now().Format("2006-01-02 15:04:05"))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		// Flush the data to the client
		// flusher, ok := w.(http.Flusher)
		// if ok {
		// 	flusher.Flush()
		// }
		// Wait for 1 second before sending the next event
		time.Sleep(1 * time.Second)
	}

	// Close the connection
	w.WriteHeader(http.StatusOK)
}

func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	tokens := []string{"this", "is", "a", "this", "is", "a"}

	for _, token := range tokens {
		content := fmt.Sprintf("data: %s\n\n", string(token))
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		// intentional wait
		time.Sleep(time.Millisecond * 420)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
