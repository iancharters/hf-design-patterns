package observer

type Observer interface {
	Update(interface{})
}

type Subject struct {
	observers []Observer
}

func NewSubject() *Subject {
	return &Subject{}
}

func (s *Subject) RegisterObserver(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) RemoveObserver(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

func (s *Subject) NotifyObservers(v interface{}) {
	for _, observer := range s.observers {
		observer.Update(v)
	}
}
