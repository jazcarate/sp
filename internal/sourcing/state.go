// Package sourcing contains state and ways to change it
package sourcing

import (
	"fmt"
	"time"

	"github.com/jazcarate/sp/internal/trianglem"
)

// A Participant represents a participant in the split.
type Participant struct {
	Name            string
	PublicKey       string
	Split           int
	SplitPercentage int // in centiunits
}

// A LogEvent represents an event in the log.
type LogEvent struct {
	ID        int
	By        string
	Operation StateChanger
	On        int64
	Note      string
	Signature string
	Valid     bool
}

// A State represents the current state.
type State struct {
	Name          string
	Participants  ([]Participant)
	Configuration string
	Balance       *trianglem.M
	Log           ([]LogEvent)
	LastOp        int
}

// NewState constructor with sensible default value.
// https://github.com/jazcarate/sp/issues/7
func NewState() *State {
	return &State{
		Name:          "Split Chain",
		Participants:  nil,
		Configuration: Trust,
		Log:           nil,
		LastOp:        0,
		Balance:       nil,
	}
}

// Apply an operation to a state.
func (s *State) Apply(op StateChanger, now time.Time) (*State, error) {
	if s == nil {
		s = NewState()
	}

	s.LastOp++
	s.Log = append(s.Log, LogEvent{
		ID:        s.LastOp,
		Operation: op,
		On:        now.Unix(),
		Note:      "",
		Signature: "1",
		By:        "joaco",
		Valid:     true,
	})

	s, err := op.apply(s)
	if err != nil {
		return nil, fmt.Errorf("couldn't apply state: %w", err)
	}

	return s, nil
}

func (s *State) findParticipant(name string) (*Participant, int, error) {
	for i, k := range s.Participants {
		if k.Name == name {
			return &s.Participants[i], i, nil
		}
	}

	return nil, -1, ErrNoParticipant
}

const percentaga = 100

func (s *State) readjustSplits() *State {
	var total int
	for _, p := range s.Participants {
		total += p.Split
	}

	for i, p := range s.Participants {
		s.Participants[i].SplitPercentage = (p.Split * percentaga) / total
	}

	return s
}
