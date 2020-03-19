package internal_tests

import (
	"marauders-map-client-desktop/internal"
	"testing"
)

func TestSimulateScreenCapture(t *testing.T) {
	recorder := internal.NewScreenRecorder(5)

	// Deploy for directories creation
	// And instances needed
	internal.Deploy()

	// recorder.StartCapturing(ch)
	recorder.ScreenShot("test-group")
}
