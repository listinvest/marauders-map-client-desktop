package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/accesspoint"}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Connecto to server
	log.Printf("Connecting to '%s'", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial: ", err)
	}
	defer c.Close()

	// Read messages
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Read ERR:", err)
				return
			}
			log.Printf("Recv: %s", message)

			// Dispatch observer event
			// ..
		}
	}()

	c.WriteMessage(websocket.TextMessage, []byte("Hello from Client"))

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:

			data := []byte(t.String())

			log.Println("Sending random data:", data)

			err := c.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println("Write ERR:", err)
				return
			}
		case <-interrupt:
			log.Println("Interrupted..")

			// Cleanly close the connection by sending a close message and then
			// Waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Write Close ERR:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}
