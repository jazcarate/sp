// Package sourcing contains state and ways to change it
package sourcing

type name = string

// A State represents the current state.
type State struct {
	participants (map[name]struct{})
}

// NewState constructs a new empty state.
func NewState() State {
	return State{
		participants: make(map[name]struct{}),
	}
}

// Participants returns current participant slice from a state.
func (s *State) Participants() []name {
	if s == nil {
		return nil
	}

	participants := make([]name, 0, len(s.participants))
	for k := range s.participants {
		participants = append(participants, k)
	}

	return participants
}

// Apply an operation to a state.
func (s State) Apply(op Operable) (State, error) {
	return op.apply(s)
}
