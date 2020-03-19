package internal

import (
	"log"
	"marauders-map-client-desktop/tools"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	// ====================
	// Websocket URL
	// ====================
	WSConfiguration

	// ====================
	// HTTP Address
	// ====================
	HTTPConfiguration
}

type WSConfiguration struct {
	Wsscheme string
	Wshost   string
	Wsport   string
	Wspath   string

	conn      *websocket.Conn
	connected bool
}

type HTTPConfiguration struct {
	httpprotocol string
	httpdomain   string
	httpport     string

	uploaduri string // Server URI for uploading files
}

/**
 * Connects to WebSocket Server
 */
func (wsc *WSClient) Connect(scheme string, host string, path string) {
	u := url.URL{Scheme: scheme, Host: host, Path: path}

	log.Printf("Connecting to '%s'", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial: ", err)
	}

	wsc.conn = c
	wsc.connected = true
}

/*
 * Disconnects from WebSocket server
 */
func (wsc *WSClient) Disconnect() {
	log.Println("Closing websocket connection")
	wsc.conn.Close()
}

/*
 * Writes message to the socket
 */
func (wsc *WSClient) SendMessage(msg string) error {
	if wsc.conn == nil || !wsc.connected {
		panic("WebSocket Connection needed!")
	}

	return wsc.conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

/*
 * Reads incoming messages from the socket
 */
func (wsc *WSClient) StartReadsMessages(ch chan string) {

	// Start decoupled Goroutine for reading messages
	go func(ch chan string) {
		defer close(ch)

		for {
			_, message, err := wsc.conn.ReadMessage()
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
func (wsc *WSClient) StartCommunications(subject *Subject) {
	ch := make(chan string)

	// TODO: goroutine here for reconnecting mechanism
	wsc.Connect("ws", "localhost:8080", "/accesspoint")

	wsc.StartReadsMessages(ch)

	for {
		rawcmd, ok := <-ch
		if !ok {
			log.Println("Connection closed!")
			break
		}

		log.Println("Command Received:", rawcmd)
		rawcmd = tools.CleanWhiteSpaces(rawcmd)
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

func NewWSClient(wsconf WSConfiguration, httpconf HTTPConfiguration) *WSClient {
	return &WSClient{
		WSConfiguration:   wsconf,
		HTTPConfiguration: httpconf,
	}
}

func NewWSConfiguration(scheme, host, port, path string) WSConfiguration {
	return WSConfiguration{
		Wsscheme: scheme,
		Wshost:   host,
		Wsport:   port,
		Wspath:   path,
	}
}

func NewHTTPConfiguration(httpprotocol, httpdomain, httpport, uploaduri string) HTTPConfiguration {
	return HTTPConfiguration{
		httpprotocol: httpprotocol,
		httpdomain:   httpdomain,
		httpport:     httpport,

		uploaduri: uploaduri,
	}
}
