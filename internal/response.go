package internal

type SendFileCommand struct {
	AbsolutePath string
	wsc          *WSClient
}

func NewSendFileCommand(wsc *WSClient) *SendFileCommand {
	return &SendFileCommand{
		wsc: wsc,
	}
}
