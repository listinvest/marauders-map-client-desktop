package internal

import (
	"fmt"
	"log"
	"net/url"
	"time"

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
	wsscheme string
	wshost   string
	wsport   string
	wspath   string

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

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		wsc.connected = false

		return
	}

	log.Println("Connected to server")

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
			for !wsc.connected {
				time.Sleep(1000)
			}

			_, message, err := wsc.conn.ReadMessage()
			if err != nil {
				wsc.connected = false
				continue
			}

			ch <- string(message)
		}

	}(ch)
}

// Entrypoint for starting communications with Server
// via websockets
func (wsc *WSClient) StartCommunications(subject *Subject) {
	ch := make(chan string)

	// Infinite loop for reconnecting mechanism
	go func() {
		for {
			if wsc.connected {
				continue
			}

			wsc.Connect(wsc.wsscheme, fmt.Sprintf("%s:%s", wsc.wshost, wsc.wsport), wsc.wspath)
		}
	}()

	wsc.StartReadsMessages(ch)

	for {
		rawcmd, ok := <-ch
		if !ok {
			log.Println("Connection closed!")
			break
		}

		_ = rawcmd

		subject.Notify(rawcmd)
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
		wsscheme: scheme,
		wshost:   host,
		wsport:   port,
		wspath:   path,
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
