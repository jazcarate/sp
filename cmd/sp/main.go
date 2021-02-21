// Application entry point
package main

import (
	"fmt"
	"os"

	"github.com/jazcarate/sp/internal/sourcing"
)

const (
	success = 0
	failure = 1
)

func main() {
	var s *sourcing.State

	s, err := s.Apply(sourcing.MultiOp{Ops: []sourcing.Operable{
		sourcing.AddParticipant{Name: "Joe"},
		sourcing.EnabbleParticipant{Name: "Joe"},
		sourcing.AddParticipant{Name: "Ben"},
		sourcing.RemoveParticipant{Name: "Ben"},
	}})

	if err != nil {
		fmt.Println(err)
		os.Exit(failure)
	} else {
		fmt.Println(s.Markdown())
		os.Exit(success)
	}
}
