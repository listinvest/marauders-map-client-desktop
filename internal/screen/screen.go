package screen

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"marauders-map-client-desktop/internal/deploy"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/kbinani/screenshot"
)

type Screenshot struct {
	FileGroup string
	FileName  string
	FilePath  string
}

type ScreenRecorder struct {
	secondsPerShot int
	recording      bool
	mux            *sync.Mutex
}

// Take screenshots and save them inside its own group folder
// inside the RECORDING folder
func (s *ScreenRecorder) ScreenShot(group string) *Screenshot {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Active display not found")
	}

	secTimeStamp := time.Now().Unix()

	if n <= 0 {
		return nil
	}

	bounds := screenshot.GetDisplayBounds(0)

	// Image name will be equal to the actual timestamp
	fileName := fmt.Sprintf("%d.png", uint64(secTimeStamp))

	// The watchtower HOME directory
	homePath := deploy.GetWatchtower().GetWatchtowerPath()

	// The watchtower RECORDING directory
	recPath := path.Join(homePath, deploy.GetWatchtower().GetRecordingFolderName())

	// Creates the folder of grouped shots
	// this path is absolute
	shotsGroup := path.Join(recPath, group)

	// Always creates it in case it doesnt exists
	os.MkdirAll(shotsGroup, os.ModePerm)

	// Finishet image absolute filePath
	filePath := path.Join(shotsGroup, fileName)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil
	}

	s.saveImage(img, filePath)

	return &Screenshot{
		FileGroup: group,
		FileName:  fileName,
		FilePath:  filePath,
	}

	// TODO: handle many monitors
	// // Iterate monitors
	// var all image.Rectangle = image.Rect(0, 0, 0, 0)
	// for i := 0; i < n; i++ {
	// 	bounds := screenshot.GetDisplayBounds(i)
	// 	all = bounds.Union(all)

	// 	img, err := screenshot.CaptureRect(bounds)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
	// 	fileName := fmt.Sprintf("%d.png", uint64(secTimeStamp))
	// 	home := deploy.GetWatchtower().GetWatchtowerPath()

	// 	filePath := path.Join(home, deploy.GetWatchtower().GetRecordingFolderName())
	// 	filePath = path.Join(filePath, fileName)

	// 	s.saveImage(img, filePath)

	// 	fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	// }
}

// Save RGBA image data to a filepath
func (s *ScreenRecorder) saveImage(img *image.RGBA, filepath string) {
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
func (s *ScreenRecorder) StartCapturing(ch chan *Screenshot) {
	s.mux.Lock()
	if s.recording {
		s.mux.Unlock()
		return
	}
	s.mux.Unlock()

	log.Println("Inside StartCapturing()")

	go func() {
		s.mux.Lock()
		s.recording = true
		s.mux.Unlock()

		// Group the images inside a folder named
		// with a timestamp
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)

		for {
			shot := s.ScreenShot(timestamp)

			s.mux.Lock()
			if !s.recording {
				s.mux.Unlock()
				break
			}
			s.mux.Unlock()

			ch <- shot

			time.Sleep(time.Duration(s.secondsPerShot) * time.Second)
		}
	}()
}

// Stop recording
func (s *ScreenRecorder) StopCapturing() {
	s.mux.Lock()
	s.recording = false
	s.mux.Unlock()
}

// Screen constructor
func NewScreenRecorder(secondsPerShot int) *ScreenRecorder {

	if secondsPerShot <= 0 {
		secondsPerShot = 5
	}

	return &ScreenRecorder{
		secondsPerShot: secondsPerShot,
		recording:      false,
		mux:            &sync.Mutex{},
	}
}
