package sourcing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParticipant_EmptyStateHasNoParticipants(t *testing.T) {
	s := NewState()

	assert.Empty(t, s.Participants())
}

func TestParticipant_NilState(t *testing.T) {
	var s *State

	assert.Empty(t, s.Participants())
}

func TestParticipant_AddingOne(t *testing.T) {
	s, err := NewState().Apply(AddParticipant("Joe"))

	assert.Empty(t, err)
	assert.ElementsMatch(t, [1]name{"Joe"}, s.Participants())
}

func TestParticipant_AddingAndRemovingHasEmptyParticipants(t *testing.T) {
	s1, err1 := NewState().Apply(AddParticipant("Joe"))
	s2, err2 := s1.Apply(RemoveParticipant("Joe"))

	assert.Empty(t, err1)
	assert.Empty(t, err2)
	assert.Empty(t, s2.Participants())
}

func TestParticipantError_RemovingANonExistantParticipantErrors(t *testing.T) {
	s, err := NewState().Apply(RemoveParticipant("Joe"))

	assert.Empty(t, s.Participants())

	if assert.Error(t, err) {
		assert.Equal(t,
			"apply <sourcing.removeParticipant{participant:\"Joe\"}>: participant <Joe>: that participant does not exist",
			err.Error())
	}
}

func TestParticipantError_AddingDuplicateParticipantsErrors(t *testing.T) {
	s1, _ := NewState().Apply(AddParticipant("Joe"))
	s2, err := s1.Apply(AddParticipant("Joe"))

	assert.ElementsMatch(t, [1]name{"Joe"}, s2.Participants())

	if assert.Error(t, err) {
		assert.Equal(t,
			"apply <sourcing.addParticipant{participant:\"Joe\"}>: participant <Joe>: participant already exists",
			err.Error())
	}
}
