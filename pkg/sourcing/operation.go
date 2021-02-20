// Package sourcing contains state and ways to change it
package sourcing

import (
	"fmt"
)

// An Operable represents a change to the state.
type Operable interface {
	apply(State) (State, error)
}

// AddParticipant Operation: Adds a new participant to the split.
func AddParticipant(p name) Operable {
	return addParticipant{p}
}

type addParticipant struct {
	participant name
}

func (op addParticipant) apply(s State) (State, error) {
	_, ok := s.participants[op.participant]
	if ok {
		return wrap(s,
			fmt.Sprintf("%#v", op),
			participantError(op.participant, errAlreadyExists),
		)
	}

	s.participants[op.participant] = struct{}{}

	return s, nil
}

type removeParticipant struct {
	participant name
}

func (op removeParticipant) apply(s State) (State, error) {
	_, ok := s.participants[op.participant]
	if !ok {
		return wrap(s,
			fmt.Sprintf("%#v", op),
			participantError(op.participant, errNoparticipant),
		)
	}

	delete(s.participants, op.participant)

	return s, nil
}

// RemoveParticipant Operation: Removes a new participant to the split.
func RemoveParticipant(p name) Operable {
	return removeParticipant{p}
}
