package screen_test

import (
	"marauders-map-client-desktop/internal/screen"
	"testing"
)

func TestScreenCapture(t *testing.T) {
	recorder := screen.NewScreenRecorder(5)

	ch := make(chan string)
	recorder.StartCapturing(ch)

	chdat := <-ch
	println("CHDAT: ", chdat)

	t.Fail()
}
