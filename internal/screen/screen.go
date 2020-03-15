package screen

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

type ScreenRecorder struct {
	secondsPerShot int
}

// Take a screenshot with all monitors inside the image
func (s *ScreenRecorder) ScreenShot() {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Active display not found")
	}

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	secTimeStamp := time.Now().Unix()

	// Iterate monitors
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		all = bounds.Union(all)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}

		// fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		fileName := fmt.Sprintf("%d.png", uint64(secTimeStamp))
		s.saveScreenShot(img, fileName)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
}

// Save RGBA image data to a filepath
func (s *ScreenRecorder) saveScreenShot(img *image.RGBA, filepath string) {
	file, err := os.Create(filepath)
	if err != nil {
		log.Printf("ERROR saving screenshot: %s", err)
		return
	}
	defer file.Close()
	png.Encode(file, img)
}

// Starts a decoupled Goroutine for recording all
// the monitors detected
func (s *ScreenRecorder) StartCapturing(ch chan string) {
	go func() {
		for {
			println("shot..")
			s.ScreenShot()
			time.Sleep(time.Duration(s.secondsPerShot) * time.Second)
		}

		// Close channel if one was given
		// if ch != nil {
		// 	close(ch)
		// }
	}()
}

// Screen constructor
func NewScreenRecorder(secondsPerShot int) *ScreenRecorder {

	if secondsPerShot <= 0 {
		secondsPerShot = 5
	}

	return &ScreenRecorder{
		secondsPerShot: secondsPerShot,
	}
}
