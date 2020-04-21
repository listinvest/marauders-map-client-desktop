package internal

import (
	"fmt"
	"log"
	"net/url"

	gostompclient "github.com/apal7/go-stomp-client"
)

type WSClient struct {
	// Websocket URL
	WSConfiguration

	connected bool

	stompcredentials StompCredentials
	stompclient      *gostompclient.Client
}

type WSConfiguration struct {
	wsscheme string
	wshost   string
	wsport   string
	wspath   string
}

type StompCredentials struct {
	username string
	password string
}

type Subscription struct {
	Queue  string
	Worker Worker
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

/*
 * Disconnects from WebSocket server
 */
func (wsc *WSClient) Disconnect() {
	log.Println("Closing websocket connection")
	wsc.stompclient.Disconnect()
}

/*
 * Add device information to a specific header
 */
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
 * Send a Message to a queue
 */
func (wsc *WSClient) Send(queue string, message []byte) {
	wsc.stompclient.Send(queue, message)
}

/*
 * Subscribe for a message
 */
//  func (wsc *WSClient) Subscribe(queue string, ch chan *gostompclient.Frame) {
func (wsc *WSClient) Subscribe(sub *Subscription) {
	wsc.stompclient.Subscribe(sub.Queue, nil, sub.Worker.GetChannel())
}

/*
 * Subscribe for a message
 */
func (wsc *WSClient) SubscribeWithChan(queue string, ch chan *gostompclient.Frame) {
	wsc.stompclient.Subscribe(queue, nil, ch)
}

/*
 * Configure queue subscriptions
 */
func (wsc *WSClient) SetupSubscriptions(subs []*Subscription) {
	for _, sub := range subs {
		wsc.Subscribe(sub)
	}
}

// ==================================================
// Constructors
// ==================================================

func NewSubscription(queue string, worker Worker) *Subscription {
	return &Subscription{
		Queue:  queue,
		Worker: worker,
	}
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
