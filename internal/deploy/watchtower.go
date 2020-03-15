package deploy

import (
	"log"
	"os"
	"path"
)

type Watchtower struct {
	homename string
	path     string
	execname string
}

// Builds watchtower
func (w *Watchtower) BuildWatchtower() {
	w.buildWatchTowerPath()
	w.setupBin()
	w.buildRoom()
}

// Build watchtower directory path
func (w *Watchtower) buildWatchTowerPath() string {
	usrhome, err := os.UserHomeDir()
	_ = err

	// path: $USERHOME/system
	wtpath := path.Join(usrhome, w.GetHomeName())
	os.MkdirAll(wtpath, os.ModePerm)

	w.path = wtpath

	log.Println("Directory created: ", wtpath, w.GetHomeName())

	return wtpath
}

// Copy binary (self copy) to the watchtower
func (w *Watchtower) setupBin() {
}

// Setup watchtower directory structure
func (w *Watchtower) buildRoom() {
	folders := []string{
		"bin",
		"resources",
		"rec",
	}

	// Creates folders
	for _, f := range folders {
		// Concatenates watchtower path with folder name
		absolutefolder := path.Join(w.GetWatchtowerPath(), f)
		// Creates folder
		os.MkdirAll(absolutefolder, os.ModePerm)
	}
}

func (w *Watchtower) GetWatchtowerPath() string {
	return w.path
}

func (w *Watchtower) GetHomeName() string {
	return w.homename
}

func NewWatchtower() *Watchtower {
	return &Watchtower{
		homename: "system",
	}
}
