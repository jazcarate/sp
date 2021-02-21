// Package sourcing contains state and ways to change it
package sourcing

type participant struct {
	name    string
	enabled bool
	split   *int
}

// A State represents the current state.
type State struct {
	participants ([]participant)
}

// Participants returns current participant slice from a state.
func (s *State) Participants() []string {
	if s == nil {
		return nil
	}

	participants := make([]string, 0, len(s.participants))
	for _, k := range s.participants {
		if k.enabled {
			participants = append(participants, k.name)
		}
	}

	return participants
}

// Apply an operation to a state.
func (s *State) Apply(op Operable) (*State, error) {
	if s == nil {
		s = &State{}
	}

	return op.apply(s)
}

func (s *State) findParticipant(name string) (*participant, error) {
	for i, k := range s.participants {
		if k.name == name {
			return &s.participants[i], nil
		}
	}

	return nil, ErrNoParticipant
}
