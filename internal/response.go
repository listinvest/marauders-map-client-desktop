package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SendFileCommand struct {
	AbsolutePath string
	wsc          *WSClient
}

func (cmd *SendFileCommand) Send(filepath string) {
	log.Println("Sending shot:", filepath)

	protocol := cmd.wsc.httpprotocol
	port := cmd.wsc.httpport
	domain := cmd.wsc.httpdomain
	uploaduri := cmd.wsc.uploaduri

	posturl := fmt.Sprintf("%s://%s:%s%s", protocol, domain, port, uploaduri)

	log.Println("Uploading to: ", posturl)

	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("File %s to send not found in directory\n", filepath)
		log.Println("ERROR:", err)
		return
	}
	defer file.Close()

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
