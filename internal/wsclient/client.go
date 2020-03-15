package wsclient

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

var conn *websocket.Conn
var connected bool

/**
 * Connects to WebSocket Server
 */
func Connect(scheme string, host string, path string) {
	u := url.URL{Scheme: scheme, Host: host, Path: path}

	log.Printf("Connecting to '%s'", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial: ", err)
	}

	conn = c
	connected = true
}

/*
 * Disconnects from WebSocket server
 */
func Disconnect() {
	log.Println("Closing websocket connection")
	conn.Close()
}

/*
 * Writes message to the socket
 */
func SendMessage(msg string) error {
	if conn == nil || !connected {
		panic("WebSocket Connection needed!")
	}

	return conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

/*
 * Reads incoming messages from the socket
 */
func StartReadsMessages(ch chan string) {

	// Start decoupled Goroutine for reading messages
	go func(ch chan string) {
		defer close(ch)

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read ERROR:", err)
				break
			}

			log.Printf("Received message: %s", message)
			ch <- string(message)
		}

	}(ch)
}
