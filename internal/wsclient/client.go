package wsclient

import (
	"log"
	"marauders-map-client-desktop/internal/wsclient/observer"
	"marauders-map-client-desktop/tools/string_tools"
	"net/url"
	"strings"

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

			ch <- string(message)
		}

	}(ch)
}

// Entrypoint for starting communications with Server
// via websockets
func StartCommunications(subject *observer.Subject) {
	ch := make(chan string)

	Connect("ws", "localhost:8080", "/accesspoint")

	StartReadsMessages(ch)

	for {
		rawcmd, ok := <-ch
		if !ok {
			log.Println("Connection closed!")
			break
		}

		log.Println("Command Received:", rawcmd)
		rawcmd = string_tools.CleanWhiteSpaces(rawcmd)
		scmd := strings.Split(rawcmd, " ")

		if len(scmd) >= 1 {
			cmd := scmd[0]
			cdata := scmd[1:]

			subject.Notify(cmd, cdata)
		}

		// TODO: delete this
		// thanksmsg := fmt.Sprintf("Thank you! ...for your message: \"%s\"", data)
		// err := SendMessage(thanksmsg)
		// if err != nil {
		// 	log.Printf("ERROR sending message: %s", data)
		// 	log.Printf("ERROR reason: %s", err)
		// }
		// log.Printf("Message Sent: %s", thanksmsg)
	}
}
