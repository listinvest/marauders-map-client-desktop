package internal

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

func (s *Subject) Notify(cmd string, data []string) {
	for _, l := range s.Observers {
		if l != nil {
			l.execute(cmd, data)
		}
	}
}

// ===========================================
// Observer interface
// ===========================================
type ObserverInterface interface {
	execute(cmd string, data []string)
}
