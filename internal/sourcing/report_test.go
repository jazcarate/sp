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

func TestMarkdown_NewState(t *testing.T) {
	s := sourcing.NewState()

	assert.Equal(t, "🗅 No state", s.Markdown())
}

func TestMarkdown_WithParticipants(t *testing.T) {
	s, _ := sourcing.NewState().Apply(sourcing.AddParticipant{Name: "Joe"})

	assert.Equal(t, "Participating: Joe", s.Markdown())
}
