// Application entry point
package main

import (
	"os"

	"github.com/jazcarate/sp/internal/sourcing"
)

func main() {
	var s *sourcing.State

	s, err := s.Apply(sourcing.MultiOp{Ops: []sourcing.StateChanger{
		sourcing.AddParticipant{Name: "Joe"},
		sourcing.SplitParticipant{Name: "Joe", NewSplit: 1},
		sourcing.AddParticipant{Name: "Ben"},
		sourcing.SplitParticipant{Name: "Ben", NewSplit: 1},
	}})
	if err != nil {
		panic(err)
	}

	err = s.Markdown(os.Stdout)
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
