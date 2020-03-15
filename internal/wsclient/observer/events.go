package observer

import (
	"marauders-map-client-desktop/tools/string_tools"
	"strings"
)

// ===========================================
// This is the subject where all
// listeners will be observing
// ===========================================
type Subject struct {
	Observers []ObserverInterface
}

func (s *Subject) AddListener(l ObserverInterface) {
	s.Observers = append(s.Observers, l)
}

func (s *Subject) Notify(m string) {
	for _, l := range s.Observers {
		if l != nil {
			l.execute(m)
		}
	}
}

// ===========================================
// Observer interface
// ===========================================
type ObserverInterface interface {
	execute(m string)
}

// KeyLogger commands Observer
type KeyloggerObserver struct {
	Msg string
}

func (o *KeyloggerObserver) execute(m string) {
	m = string_tools.CleanWhiteSpaces(m)
	o.Msg = m

	ms := strings.Split(m, " ")
	_ = ms
}

func NewKeyloggerObserver() *KeyloggerObserver {
	return new(KeyloggerObserver)
}
