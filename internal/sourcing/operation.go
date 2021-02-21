// Package sourcing contains state and ways to change it
package sourcing

// An Operable represents a change to the state.
type Operable interface {
	apply(State) (State, error)
}

// AddParticipant Operation: Adds a new participant to the split.
type AddParticipant struct{ Name name }

func (op AddParticipant) apply(s State) (State, error) {
	participant := op.Name
	_, ok := s.participants[participant]

	if ok {
		return s, &ApplyError{PreviousState: s, Op: op, Err: ErrAlreadyExists}
	}

	s.participants[participant] = struct{}{}

	return s, nil
}

// RemoveParticipant Operation: Removes a new participant to the split.
type RemoveParticipant struct{ Name name }

func (op RemoveParticipant) apply(s State) (State, error) {
	participant := op.Name
	_, ok := s.participants[participant]

	if !ok {
		return s, &ApplyError{PreviousState: s, Op: op, Err: ErrNoparticipant}
	}

	delete(s.participants, participant)

	return s, nil
}
