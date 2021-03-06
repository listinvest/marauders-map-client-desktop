package internal

import (
	"marauders-map-client-desktop/tools"
	"os"
	"path"
	"path/filepath"
)

const (
	binFolder       = "bin"
	resourcesFolder = "resources"
	recFolder       = "rec"
	downloadsFolder = "downloads"
	kldataFolder    = "kl"
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
	tools.CopyFile(w.GetBinaryPath(), p)
}

// Setup watchtower directory structure
func (w *Watchtower) buildRoom() {

	if !w.homeBuilt {
		w.buildWatchtowerHome()
	}

	folders := []string{
		binFolder,
		resourcesFolder,
		recFolder,
		downloadsFolder,
		kldataFolder,
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
	// TODO: create autoexecution mechanism depending the OS
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

// ==============================
// Watchtower HOME folders
// ==============================
func (w *Watchtower) GetBinFolderName() string {
	return binFolder
}

func (w *Watchtower) GetResourcesFolderName() string {
	return resourcesFolder
}

func (w *Watchtower) GetRecordingFolderName() string {
	return recFolder
}

func (w *Watchtower) GetDownloadsFolderName() string {
	return downloadsFolder
}

func (w *Watchtower) GetAbsoluteDownloadsFolderPath() string {
	// Watchtower home path
	homePath := watchtower.homepath

	// Download path
	downloadsPath := path.Join(homePath, watchtower.GetDownloadsFolderName())

	return downloadsPath
}

// Watchtower constructor
func NewWatchtower() *Watchtower {
	return &Watchtower{
		homename: "system",
	}
}
