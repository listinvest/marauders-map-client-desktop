package deploy

import (
	"marauders-map-client-desktop/tools/ostools"
	"os"
	"path"
	"path/filepath"
)

type Watchtower struct {
	homename string
	homepath string

	built     bool
	homeBuilt bool
	roomBuilt bool
}

// Builds watchtower
func (w *Watchtower) BuildWatchtower() {
	w.buildWatchtowerHome()
	w.buildRoom()
	w.setupBin()

	w.built = true
}

// Build watchtower directory path
func (w *Watchtower) buildWatchtowerHome() string {
	usrhome, err := os.UserHomeDir()
	_ = err

	// path: $USERHOME/system
	wtpath := path.Join(usrhome, w.GetHomeName())
	os.MkdirAll(wtpath, os.ModePerm)

	w.homepath = wtpath

	w.homeBuilt = true

	return wtpath
}

// Copy binary (self copy) to the watchtower
func (w *Watchtower) setupBin() {

	if !w.roomBuilt {
		w.buildRoom()
	}

	p := path.Join(w.GetWatchtowerPath(), w.GetBinaryName())

	// Copy binary to HOME
	ostools.CopyFile(w.GetBinaryPath(), p)
}

// Setup watchtower directory structure
func (w *Watchtower) buildRoom() {

	if !w.homeBuilt {
		w.buildWatchtowerHome()
	}

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

// Setup watchtower directory structure
func (w *Watchtower) Daemonize() {
	// TODO: create autoexecution mechanism depending de OS
}

func (w *Watchtower) GetWatchtowerPath() string {
	return w.homepath
}

func (w *Watchtower) GetHomeName() string {
	return w.homename
}

func (w *Watchtower) GetBinaryName() string {
	return filepath.Base(os.Args[0])
}

func (w *Watchtower) GetBinaryPath() string {
	return os.Args[0]
}

func NewWatchtower() *Watchtower {
	return &Watchtower{
		homename: "system",
	}
}
