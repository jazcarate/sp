// Application entry point
package main

import (
	"os"
	"time"

	"github.com/jazcarate/sp/internal/sourcing"
)

func main() {
	var s *sourcing.State

	s, err := s.Apply(sourcing.MultiOp{
		Ops: []sourcing.StateChanger{
			sourcing.AddParticipant{Name: "Joe", PublicKey: "3"},
			sourcing.AddParticipant{Name: "Bob", PublicKey: "1"},
			sourcing.Transfer{From: "Joe", To: "Bob", Amount: 300},
		},
	}, time.Now())
	if err != nil {
		panic(err)
	}

	err = s.Markdown(os.Stdout)
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
