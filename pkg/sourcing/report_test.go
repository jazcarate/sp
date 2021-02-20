package sourcing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkdown_NilState(t *testing.T) {
	var s *State

	assert.Equal(t, "🗅 No state", s.Markdown())
}

func TestMarkdown_NewState(t *testing.T) {
	s := NewState()

	assert.Equal(t, "🗅 No state", s.Markdown())
}

func TestMarkdown_WithParticipants(t *testing.T) {
	s, _ := NewState().Apply(AddParticipant("Joe"))

	assert.Equal(t, "Participating: Joe", s.Markdown())
}
