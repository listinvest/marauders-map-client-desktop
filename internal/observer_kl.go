package internal

// KeyLogger commands Observer
type KeyloggerCmdObserver struct {
}

func (o *KeyloggerCmdObserver) execute(string_json string) {
}

func NewKeyloggerCmdObserver() *KeyloggerCmdObserver {
	return &KeyloggerCmdObserver{}
}
