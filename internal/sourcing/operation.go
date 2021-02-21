// Package sourcing contains state and ways to change it
package sourcing

import (
	"errors"
	"fmt"
)

// An Operable represents a change to the state.
type Operable interface {
	apply(*State) (*State, error)
}

// MultiOp Operation: Bundles multiple operations in secuence.
// First error halts the whole operation.
type MultiOp struct{ Ops []Operable }

func (mop MultiOp) apply(s *State) (*State, error) {
	var err error

	for i, op := range mop.Ops {
		s, err = s.Apply(op)
		if err != nil {
			return nil, fmt.Errorf("couldn't apply operation #%v: %w", i, err)
		}
	}

	return s, nil
}

// AddParticipant Operation: Adds a new participant to the split.
type AddParticipant struct{ Name string }

func (op AddParticipant) apply(s *State) (*State, error) {
	needle := op.Name
	_, err := s.findParticipant(needle)

	if !errors.Is(err, ErrNoParticipant) {
		return s, &ApplyError{PreviousState: s, Op: op, Err: ErrAlreadyExists}
	}

	s.participants = append(s.participants, participant{name: needle, enabled: false, split: nil})

	return s, nil
}

// EnabbleParticipant Operation: Enables an existing participant to be part of the split.
type EnabbleParticipant struct{ Name string }

func (op EnabbleParticipant) apply(s *State) (*State, error) {
	p, err := s.findParticipant(op.Name)
	if err != nil {
		return nil, &ApplyError{PreviousState: s, Op: op, Err: err}
	}

	p.enabled = true

	return s, nil
}

// RemoveParticipant Operation: Removes a new participant to the split.
type RemoveParticipant struct{ Name string }

func (op RemoveParticipant) apply(s *State) (*State, error) {
	p, err := s.findParticipant(op.Name)
	if err != nil {
		return nil, &ApplyError{PreviousState: s, Op: op, Err: err}
	}

	p.enabled = false

	return s, nil
}
