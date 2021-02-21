package sourcing_test

import (
	"testing"

	"github.com/jazcarate/sp/internal/sourcing"
	"github.com/stretchr/testify/assert"
)

func TestParticipant_EmptyStateHasNoParticipants(t *testing.T) {
	s := sourcing.NewState()

	assert.Empty(t, s.Participants())
}

func TestParticipant_NilState(t *testing.T) {
	var s *sourcing.State

	assert.Empty(t, s.Participants())
}

func TestParticipant_AddingOne(t *testing.T) {
	s, err := sourcing.NewState().Apply(sourcing.AddParticipant{Name: "Joe"})

	assert.Empty(t, err)
	assert.ElementsMatch(t, [1]string{"Joe"}, s.Participants())
}

func TestParticipant_AddingAndRemovingHasEmptyParticipants(t *testing.T) {
	s1, err1 := sourcing.NewState().Apply(sourcing.AddParticipant{Name: "Joe"})
	s2, err2 := s1.Apply(sourcing.RemoveParticipant{Name: "Joe"})

	assert.Empty(t, err1)
	assert.Empty(t, err2)
	assert.Empty(t, s2.Participants())
}

func TestParticipantError_RemovingANonExistantParticipantErrors(t *testing.T) {
	s, err := sourcing.NewState().Apply(sourcing.RemoveParticipant{Name: "Joe"})

	assert.Empty(t, s.Participants())

	if assert.Error(t, err) {
		assert.Equal(t,
			"apply <sourcing.RemoveParticipant{Name:\"Joe\"}>: that participant does not exist",
			err.Error())
	}
}

func TestParticipantError_AddingDuplicateParticipantsErrors(t *testing.T) {
	s1, _ := sourcing.NewState().Apply(sourcing.AddParticipant{Name: "Joe"})
	s2, err := s1.Apply(sourcing.AddParticipant{Name: "Joe"})

	assert.ElementsMatch(t, [1]string{"Joe"}, s2.Participants())

	if assert.Error(t, err) {
		assert.Equal(t,
			"apply <sourcing.AddParticipant{Name:\"Joe\"}>: participant already exists",
			err.Error())
	}
}
