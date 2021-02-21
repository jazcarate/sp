package sourcing_test

import (
	"testing"

	"github.com/jazcarate/sp/internal/sourcing"
	"github.com/stretchr/testify/assert"
)

func TestMarkdown_NilState(t *testing.T) {
	var s *sourcing.State

	assert.Equal(t, "🗅 No state", s.Markdown())
}

func TestMarkdown_HidesDisabledParticipants(t *testing.T) {
	var s *sourcing.State

	s, err := s.Apply(sourcing.AddParticipant{Name: "Joe"})

	assert.Empty(t, err)
	assert.Equal(t, "🗅 No state", s.Markdown())
}

func TestMarkdown_WithParticipants(t *testing.T) {
	var s *sourcing.State

	s, err1 := s.Apply(sourcing.AddParticipant{Name: "Joe"})
	s, err2 := s.Apply(sourcing.EnabbleParticipant{Name: "Joe"})
	s, err3 := s.Apply(sourcing.AddParticipant{Name: "Ben"})

	assert.Empty(t, err1)
	assert.Empty(t, err2)
	assert.Empty(t, err3)
	assert.Equal(t, "Participating: Joe", s.Markdown())
}
