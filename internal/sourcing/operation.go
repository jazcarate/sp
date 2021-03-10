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
		s, err = op.apply(s)
		if err != nil {
			return nil, fmt.Errorf("couldn't apply operation #%v: %w", i, err)
		}
	}

	return s, nil
}

// AddParticipant Operation: Adds a new participant to the split with a default split of 0.
type AddParticipant struct {
	Name      string
	PublicKey string
}

func (op AddParticipant) apply(s *State) (*State, error) {
	needle := op.Name
	_, _, err := s.findParticipant(needle)

	if !errors.Is(err, ErrNoParticipant) {
		return s, &ApplyError{PreviousState: s, Op: op, Err: ErrAlreadyExists}
	}

	s.Participants = append(s.Participants, Participant{
		Name:            needle,
		PublicKey:       op.PublicKey,
		Split:           0,
		SplitPercentage: 0,
	})
	s.Balance = s.Balance.Incr()

	return s, nil
}

// SplitParticipant Operation: Changes the split of a participant.
type SplitParticipant struct {
	Name     string
	NewSplit int
}

func (op SplitParticipant) apply(s *State) (*State, error) {
	p, _, err := s.findParticipant(op.Name)
	if err != nil {
		return nil, &ApplyError{PreviousState: s, Op: op, Err: err}
	}

	p.Split = op.NewSplit
	s = s.readjustSplits()

	return s, nil
}

// Transfer Operation: Moves money around.
type Transfer struct {
	From   string
	To     string
	Amount int
}

func (op Transfer) apply(s *State) (*State, error) {
	_, from, errF := s.findParticipant(op.From)
	_, to, errT := s.findParticipant(op.To)

	if errF != nil {
		return nil, &ApplyError{PreviousState: s, Op: op, Err: fmt.Errorf("transfer from: %w", errF)}
	}

	if errT != nil {
		return nil, &ApplyError{PreviousState: s, Op: op, Err: fmt.Errorf("transfer to: %w", errT)}
	}

	rem := op.Amount
	// Cancel the direct debt
	// Start paying other's debts

	for intermediate, debt := range s.Balance.Iterate(to) {
		debt := debt
		if debt < 0 && from != intermediate && to != intermediate {
			val := max(rem, -debt)

			errF = s.Balance.Modify(from, intermediate, func(i int) int { return i - val })

			if errF != nil {
				return nil, &ApplyError{PreviousState: s, Op: op, Err: fmt.Errorf("transfer change from: %w", errF)}
			}

			errT = s.Balance.Modify(intermediate, to, func(i int) int { return i - val })

			if errT != nil {
				return nil, &ApplyError{PreviousState: s, Op: op, Err: fmt.Errorf("transfer change to: %w", errT)}
			}

			rem -= val
		}
	}

	err := s.Balance.Modify(from, to, func(i int) int { return i - rem })
	if err != nil {
		return nil, &ApplyError{PreviousState: s, Op: op, Err: fmt.Errorf("transfer change remainder: %w", err)}
	}

	return s, nil
}

func max(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// A SigningConfiguration dictates how to verify each operation.
const (
	// Trust means that no signing required. Default configuration.
	Trust string = "Trust"
	// All means everyone has to sign off every operation.
	All = "All"
	// Involved menas that only parties involved need to sign.
	Involved = "Involved"
)

// Configure Operation: Changes the current trust configuration.
type Configure struct{ NewConfig string }

func (op Configure) apply(s *State) (*State, error) {
	s.Configuration = op.NewConfig
	return s, nil
}

// Spend Operation: Looks at the split and sucks money from everyoneo else.
type Spend struct {
	Who    string
	Amount int
}

func (op Spend) apply(s *State) (*State, error) {
	_, x, err := s.findParticipant(op.Who)
	if err != nil {
		return nil, &ApplyError{PreviousState: s, Op: op, Err: err}
	}

	for y, p := range s.Participants {
		p1 := p

		if x == y {
			continue
		}

		err = s.Balance.Modify(x, y, func(x int) int { return x - (op.Amount * p1.SplitPercentage / 100) })

		if err != nil {
			return nil, &ApplyError{PreviousState: s, Op: op, Err: err}
		}
	}

	return s, nil
}
