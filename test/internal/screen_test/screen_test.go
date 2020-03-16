package screen_test

import (
	"marauders-map-client-desktop/internal/deploy"
	"marauders-map-client-desktop/internal/screen"
	"testing"
)

func TestSimulateScreenCapture(t *testing.T) {
	recorder := screen.NewScreenRecorder(5)

	// Deploy for directories creation
	// And instances needed
	deploy.Deploy()

	// recorder.StartCapturing(ch)
	recorder.ScreenShot("test-group")
}
