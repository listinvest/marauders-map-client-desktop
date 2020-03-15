package main

import (
	"fmt"
	"log"
	"marauders-map-client-desktop/internal/wsclient"
)

func main() {

	ch := make(chan string)
	wsclient.Connect("ws", "localhost:8080", "/accesspoint")

	wsclient.StartReadsMessages(ch)

	for {
		data, ok := <-ch
		if !ok {
			log.Println("Connection closed!")
			break
		}

		println("Received: ", data)

		err := wsclient.SendMessage(fmt.Sprintf("I received: %s", data))
		if err != nil {
			log.Printf("ERROR sending message: %s", data)
			log.Printf("ERROR reason: %s", err)
		}
	}

}
