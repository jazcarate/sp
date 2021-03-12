// Application entry point
package main

import (
	"os"
	"time"

	"github.com/jazcarate/sp/internal/sourcing"
)

func main() {
	var s *sourcing.State

	now := time.Now()

	s, err := s.Apply(sourcing.Configure{NewConfig: sourcing.All}, now)
	if err != nil {
		panic(err)
	}

	err = s.Markdown(os.Stdout)
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
