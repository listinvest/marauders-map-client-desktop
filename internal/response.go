package internal

import "log"

type SendFileCommand struct {
	AbsolutePath string
	wsc          *WSClient
}

func (cmd *SendFileCommand) Send(filepath string) {
	log.Println("Sending shot:", filepath)
}

func NewSendFileCommand(wsc *WSClient) *SendFileCommand {
	return &SendFileCommand{
		wsc: wsc,
	}
}
