package internal_tests

import (
	"marauders-map-client-desktop/internal"
	"marauders-map-client-desktop/tools"
	"strings"
	"testing"
)

func TestObserverNotified(t *testing.T) {

	o := internal.NewKeyloggerCmdObserver()

	subject := internal.Subject{Observers: make([]internal.ObserverInterface, 0)}
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

		c = tools.CleanWhiteSpaces(c)
		xs := strings.Split(c, " ")

		// Notify to the observer
		subject.Notify(xs[0], xs[1:])

		// Check if Msg notified is equal
		// TODO
	}
}
