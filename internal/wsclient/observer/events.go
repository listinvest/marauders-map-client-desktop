package observer

import (
	"fmt"
	"marauders-map-client-desktop/tools/string_tools"
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
	fmt.Printf("%q KLObserver - message received:\n", m)
	o.Msg = m
}

func NewKeyloggerObserver() *KeyloggerObserver {
	return new(KeyloggerObserver)
}
