// Application entry point
package main

import (
	"fmt"
	"os"

	"github.com/jazcarate/sp/pkg/sourcing"
)

func multiApply(ops []sourcing.Operable) (*sourcing.State, error) {
	var err error

	s := sourcing.NewState()

	for _, op := range ops {
		s, err = s.Apply(op)
		if err != nil {
			return nil, fmt.Errorf("coudn't apply: %w", err)
		}
	}

	return &s, nil
}

const (
	success = 0
	failure = 1
)

func main() {
	s, err := multiApply([]sourcing.Operable{
		sourcing.AddParticipant("Joe"),
		sourcing.AddParticipant("Ben"),
		sourcing.RemoveParticipant("Ben"),
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(failure)
	} else {
		fmt.Println(s.Markdown())
		os.Exit(success)
	}
}
