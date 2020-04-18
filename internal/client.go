package internal

import (
	"fmt"
	"log"
	"net/url"

	gostompclient "github.com/apal7/go-stomp-client"

	"github.com/gorilla/websocket"
)

type WSClient struct {
	// ====================
	// Websocket URL
	// ====================
	WSConfiguration

	registered bool

	stompcredentials StompCredentials
	stompclient      *gostompclient.Client
}

type StompCredentials struct {
	username string
	password string
}

type WSConfiguration struct {
	wsscheme string
	wshost   string
	wsport   string
	wspath   string

	conn      *websocket.Conn
	connected bool

	stompclient *gostompclient.Client
}

/**
 * Connects to WebSocket Server
 */
func (wsc *WSClient) Connect() {
	wsurl := url.URL{Scheme: wsc.wsscheme, Host: fmt.Sprintf("%s:%s", wsc.wshost, wsc.wsport), Path: wsc.wspath}

	// Load device info
	dev := NewDevice()
	dev.LoadInfo()

	clientheaders := gostompclient.NewHeader()
	wsc.AddDeviceToHeaders(dev, clientheaders)

	clientheaders.AddHeader("username", wsc.stompcredentials.username)
	clientheaders.AddHeader("password", wsc.stompcredentials.password)

	stompclient := gostompclient.NewClient(wsurl.String(), nil)
	stompclient.Connect(clientheaders)
}

func (wsc *WSClient) AddDeviceToHeaders(device Device, headers *gostompclient.Header) *gostompclient.Header {
	if headers == nil {
		headers = gostompclient.NewHeader()
	}

	// Load device info
	devmap := device.ToMap()

	for k, v := range devmap {
		headers.AddHeader(k, v)
	}

	headers.AddHeader("foo", "bar")

	return headers
}

/**
 * Connects to WebSocket Server
 */
func (wsc *WSClient) SetCredentials(username string, password string) {
	wsc.stompcredentials = StompCredentials{
		username: username,
		password: password,
	}
}

/*
 * Disconnects from WebSocket server
 */
func (wsc *WSClient) Disconnect() {
	log.Println("Closing websocket connection")
	wsc.conn.Close()
	wsc.stompclient.Disconnect()
}

func NewWSClient(wsconf WSConfiguration) *WSClient {
	return &WSClient{
		WSConfiguration: wsconf,
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
