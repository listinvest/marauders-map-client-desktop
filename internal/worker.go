package internal

import gostompclient "github.com/apal7/go-stomp-client"

type Worker interface {
	GetChannel() chan *gostompclient.Frame
	ProcessFrame(frame *gostompclient.Frame)
}
