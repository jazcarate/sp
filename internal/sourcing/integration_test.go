package sourcing_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/jazcarate/sp/internal/sourcing"
	"github.com/stretchr/testify/assert"
)

func TestIntegration_linksToDocWhenNewSP(t *testing.T) {
	var (
		s *sourcing.State
		b bytes.Buffer
	)

	err := s.Markdown(&b)

	assert.Empty(t, err)
	assert.Equal(t, "^# Split Chain", b.String())
	assert.Regexp(t, "^# Split Chain", b.String())
}

func do(input string) (string, error) {
	var (
		s *sourcing.State
		b bytes.Buffer
	)

	op, parseErr := sourcing.Parse(input)
	if parseErr != nil {
		return "", fmt.Errorf("parsing %v: %w", input, parseErr)
	}

	s, stErr := s.Apply(op)
	if stErr != nil {
		return "", fmt.Errorf("applying %v: %w", op, stErr)
	}

	tmplErr := s.Markdown(&b)
	if tmplErr != nil {
		return "", fmt.Errorf("templating %v: %w", op, tmplErr)
	}

	return b.String(), nil
}
