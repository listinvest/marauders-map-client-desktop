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

func (s *Subject) Notify(string_json string) {
	for _, l := range s.Observers {
		if l != nil {
			l.execute(string_json)
		}
	}
}

// ===========================================
// Observer interface
// ===========================================
type ObserverInterface interface {
	execute(cmd string)
}
