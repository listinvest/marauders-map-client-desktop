package internal

import (
	"io"
	"log"
	"marauders-map-client-desktop/tools"
	"net/http"
	"os"
	"path"
)

// ==========================================================
// Observer for sending to server files
// ==========================================================
type FileCmdObserver struct {
	sendFileCmd *SendFileCommand
	watchtower  *Watchtower
}

func (o *FileCmdObserver) execute(cmd string, data []string) {
	if cmd != "file" {
		return
	}

	if len(data) <= 0 {
		return
	}

	action := data[0]

	switch action {
	case "send":
		files := data[1:]
		o.sendFiles(files)
		break
	case "download":
		urls := data[1:]
		for _, url := range urls {
			o.downloadFile(url)
		}
		break
	}

}

func (o *FileCmdObserver) sendFiles(files []string) {
	for _, f := range files {
		if !tools.FileExists(f) {
			log.Printf("File requested '%s' doesn't exist\n", f)
			continue
		}

		o.sendFileCmd.Send(f)
	}
}

func (o *FileCmdObserver) downloadFile(url string) error {
	log.Println("Downloading: ", url)

	downloadsfolder := watchtower.GetAbsoluteDownloadsFolderPath()
	filename := tools.ExtractFileNameFromURL(url)

	// Absolute filePath
	finalFilePath := path.Join(downloadsfolder, filename)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Println("ERROR downloading:", err)
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(finalFilePath)
	if err != nil {
		log.Println("ERROR saving downloaded file:", err)
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	log.Printf("File %s downloaded\n", url)
	return err
}

func NewFileCmdObserver(sendFileCmd *SendFileCommand, watchtower *Watchtower) *FileCmdObserver {
	return &FileCmdObserver{
		sendFileCmd: sendFileCmd,
		watchtower:  watchtower,
	}
}
