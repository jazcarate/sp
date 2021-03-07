// Package sourcing contains state and ways to change it
package sourcing

import (
	"errors"
	"fmt"
)

// StateChanger represents one of the operations that can modify the state.
type StateChanger interface {
	apply(*State) (*State, error)
}

// MultiOp Operation: Bundles multiple operations in secuence.
// First error halts the whole operation.
type MultiOp struct{ Ops []StateChanger }

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

// AddParticipant Operation: Adds a new participant to the split with a default split of 0
type AddParticipant struct{ Name string }

func (op AddParticipant) apply(s *State) (*State, error) {
	needle := op.Name
	_, err := s.findParticipant(needle)

	if !errors.Is(err, ErrNoParticipant) {
		return s, &ApplyError{PreviousState: s, Op: op, Err: ErrAlreadyExists}
	}

	s.Participants = append(s.Participants, Participant{Name: needle, Split: 0})

	return s, nil
}

// SplitParticipant Operation: Changes the split of a participant.
type SplitParticipant struct {
	Name     string
	NewSplit int
}

func (op SplitParticipant) apply(s *State) (*State, error) {
	p, err := s.findParticipant(op.Name)
	if err != nil {
		return nil, &ApplyError{PreviousState: s, Op: op, Err: err}
	}

	p.Split = op.NewSplit

	return s, nil
}

// A SigningConfiguration dictates how to verify each operation.
type SigningConfiguration string

func (c SigningConfiguration) String() string {
	return string(c)
}

const (
	// Trust means that no signing required. Default configuration.
	Trust SigningConfiguration = "Trust"
	// All means everyone has to sign off every operation.
	All = "All"
	// Involved menas that only parties involved need to sign.
	Involved = "Involved"
)

// Configure Operation: Changes the current trust configuration.
type Configure struct{ NewConfig SigningConfiguration }

func (op Configure) apply(s *State) (*State, error) {
	s.Configuration = op.NewConfig
	return s, nil
}
