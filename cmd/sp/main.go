// Application entry point
package main

import (
	"os"

	"github.com/jazcarate/sp/internal/sourcing"
)

func main() {
	var s *sourcing.State

	s, err := s.Apply(sourcing.MultiOp{Ops: []sourcing.StateChanger{
		sourcing.AddParticipant{Name: "Joe", PublicKey: "1"},
		sourcing.SplitParticipant{Name: "Joe", NewSplit: 1},
		sourcing.AddParticipant{Name: "Ben", PublicKey: "2"},
		sourcing.SplitParticipant{Name: "Ben", NewSplit: 1},
	}})
	if err != nil {
		panic(err)
	}

	s, err = s.Apply(sourcing.AddParticipant{Name: "Jerry", PublicKey: "3"})
	if err != nil {
		panic(err)
	}

	s, err = s.Apply(sourcing.Transfer{From: "Joe", To: "Ben", Amount: 10})
	if err != nil {
		panic(err)
	}

	err = s.Markdown(os.Stdout)
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
