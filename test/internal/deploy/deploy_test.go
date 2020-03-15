package deploy

import (
	"log"
	"marauders-map-client-desktop/internal/deploy"
	"os"
	"testing"
)

func TestDeploy(t *testing.T) {
	deploy.Deploy()
	watchtower := deploy.GetWatchtower()

	wthomepath := watchtower.GetWatchtowerPath()
	wtbinpath := watchtower.GetBinaryPath()

	if _, err := os.Stat(wthomepath); os.IsNotExist(err) {
		log.Println("Watchtower directory wasn't created")
		log.Printf("expected: '%s'  - given: '%s'", wthomepath, wthomepath)
		t.Fail()
	}

	if _, err := os.Stat(wtbinpath); os.IsNotExist(err) {
		log.Println("Watchtower binary wasn't created")
		log.Printf("expected: '%s'  - given: '%s'", wthomepath, wthomepath)
		t.Fail()
	}
}
