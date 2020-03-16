package observer

import "log"

// KeyLogger commands Observer
type ScreenshotCmdObserver struct {
}

func (o *ScreenshotCmdObserver) execute(cmd string, data []string) {
	if cmd != "screen" {
		return
	}

	log.Println("ScreenshotCmdObserver: new action triggered")
}

func NewScreenshotCmdObserver() *ScreenshotCmdObserver {
	return &ScreenshotCmdObserver{}
}
