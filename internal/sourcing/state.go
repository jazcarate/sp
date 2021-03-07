// Package sourcing contains state and ways to change it
package sourcing

import (
	"time"

	"github.com/jazcarate/sp/internal/trianglem"
)

// A Participant represents a participant in the split.
type Participant struct {
	Name            string
	Split           int
	SplitPercentage int // in centiunits
}

// A LogEvent represents an event in the log.
type LogEvent struct {
	By        string
	Operation StateChanger
	On        time.Time
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

// TODO Can we make the zero value?
// NewState constructor with sensible default value.
func NewState() *State {
	return &State{
		Name:          "Split Chain",
		Participants:  nil,
		Configuration: Trust,
		Log:           nil,
		LastOp:        -1,
		Balance:       nil,
	}
}

// Apply an operation to a state.
func (s *State) Apply(op StateChanger) (*State, error) {
	if s == nil {
		s = NewState()
	}

	s.Log = append(s.Log, LogEvent{
		Operation: op,
		On:        time.Now(),
		Note:      "",
		Signature: "1",     // TODO
		By:        "joaco", // TODO
		Valid:     true,    // TODO
	})
	s.LastOp++

	return op.apply(s)
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
