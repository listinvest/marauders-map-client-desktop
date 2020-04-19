package internal

import (
	"log"

	gostompclient "github.com/apal7/go-stomp-client"
)

type BashWorker struct {
	ch      chan *gostompclient.Frame
	running bool
}

func (w *BashWorker) GetChannel() chan *gostompclient.Frame {
	return w.ch
}

func (w *BashWorker) SetChannel(ch chan *gostompclient.Frame) {
	w.ch = ch
}

func (w *BashWorker) Start() {
	if w.running {
		return
	}

	if w.ch == nil {
		w.SetupChannel()
	}

	w.running = true
	go func() {
		for w.running {
			log.Println("BashWorker - waiting for frame")
			frame := <-w.ch
			log.Println("Frame received:", string(frame.Body))
		}
	}()
}

func (w *BashWorker) Stop() {
	w.running = false
	close(w.ch)
	log.Println("BashWorker - channel closed")
}

func (w *BashWorker) SetupChannel() {
	w.ch = make(chan *gostompclient.Frame)
}

func NewBashWorker() *BashWorker {
	bworker := &BashWorker{}
	bworker.SetupChannel()

	return bworker
}
