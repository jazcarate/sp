// Package sourcing contains state and ways to change it
package sourcing

type Participant struct {
	Name  string
	Split int
}

// A State represents the current state.
// TODO split inti "ministates" whith ops for each.
type State struct {
	Name          string
	Participants  ([]Participant)
	Configuration SigningConfiguration
}

// NewState constructor with sensible default value
// TODO Can we make the zero value?
func NewState() *State {
	return &State{
		Name:          "Split Chain",
		Participants:  nil,
		Configuration: Trust,
	}
}

// Apply an operation to a state.
func (s *State) Apply(op StateChanger) (*State, error) {
	if s == nil {
		s = NewState()
	}

	return op.apply(s)
}

func (s *State) findParticipant(name string) (*Participant, error) {
	for i, k := range s.Participants {
		if k.Name == name {
			return &s.Participants[i], nil
		}
	}

	return nil, ErrNoParticipant
}
