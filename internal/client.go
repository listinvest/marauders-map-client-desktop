package internal

import (
	"fmt"
	"log"
	"net/url"

	gostompclient "github.com/apal7/go-stomp-client"
)

type WSClient struct {
	// ====================
	// Websocket URL
	// ====================
	WSConfiguration

	connected bool

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

	wsc.stompclient = gostompclient.NewClient(wsurl.String(), nil)
	wsc.stompclient.Connect(clientheaders)

	wsc.stompclient.InitInboundChannel()

	wsc.connected = wsc.stompclient.IsSTOMPConnected()
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
	wsc.stompclient.Disconnect()
}

/*
 * Send a Message to a queue
 */
func (wsc *WSClient) Send(queue string, message []byte) {
	wsc.stompclient.Send(queue, message)
}

/*
 * Subscribe for a message
 */
func (wsc *WSClient) Subscribe(queue string, ch chan *gostompclient.Frame) {
	wsc.stompclient.Subscribe(queue, nil, ch)
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
