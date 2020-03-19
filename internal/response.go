package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SendFileCommand struct {
	wsc *WSClient
}

func (cmd *SendFileCommand) Send(filepath string) {
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
		return
	}
	defer file.Close()

	// POST it
	res, err := http.Post(posturl, "binary/octet-stream", file)
	if err != nil {
		log.Printf("Couldnt send file %s to URL %s\n", filepath, posturl)
		log.Println("ERROR: ", err)
		return
	}
	defer res.Body.Close()

	message, _ := ioutil.ReadAll(res.Body)
	log.Println(string(message))
}

func NewSendFileCommand(wsc *WSClient) *SendFileCommand {
	return &SendFileCommand{
		wsc: wsc,
	}
}
