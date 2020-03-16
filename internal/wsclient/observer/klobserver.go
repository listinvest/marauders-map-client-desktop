package observer

import "log"

// KeyLogger commands Observer
type KeyloggerCmdObserver struct {
}

func (o *KeyloggerCmdObserver) execute(cmd string, data []string) {
	if cmd != "kl" {
		return
	}

	log.Println("KeyloggerCmdObserver: new action triggered")
}

func NewKeyloggerCmdObserver() *KeyloggerCmdObserver {
	return &KeyloggerCmdObserver{}
}
