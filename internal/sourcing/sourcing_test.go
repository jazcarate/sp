package sourcing_test

import (
	"testing"

	"github.com/jazcarate/sp/internal/sourcing"
	"github.com/stretchr/testify/assert"
)

func TestParticipant_EmptyStateHasNoParticipants(t *testing.T) {
	var s *sourcing.State

	assert.Empty(t, s.Participants())
}

func TestParticipant_NilState(t *testing.T) {
	var s *sourcing.State

	assert.Empty(t, s.Participants())
}

func TestParticipant_AddingOne(t *testing.T) {
	var s *sourcing.State
	s, err := s.Apply(sourcing.MultiOp{Ops: []sourcing.Operable{
		sourcing.AddParticipant{Name: "Joe"},
		sourcing.EnabbleParticipant{Name: "Joe"},
	}})

	assert.Empty(t, err)
	assert.ElementsMatch(t, [1]string{"Joe"}, s.Participants())
}

func TestParticipant_AddingAndRemovingHasEmptyParticipants(t *testing.T) {
	var s *sourcing.State

	s, err := s.Apply(sourcing.MultiOp{Ops: []sourcing.Operable{
		sourcing.AddParticipant{Name: "Joe"},
		sourcing.RemoveParticipant{Name: "Joe"},
	}})

	assert.Empty(t, err)
	assert.Empty(t, s.Participants())
}

func TestParticipantError_RemovingANonExistantParticipantErrors(t *testing.T) {
	var s *sourcing.State
	s, err := s.Apply(sourcing.RemoveParticipant{Name: "Joe"})

	assert.Empty(t, s.Participants())

	if assert.Error(t, err) {
		assert.Equal(t,
			"apply <sourcing.RemoveParticipant{Name:\"Joe\"}>: that participant does not exist",
			err.Error())
	}
}

func TestParticipantError_AddingDuplicateParticipantsErrors(t *testing.T) {
	var s *sourcing.State

	_, err := s.Apply(sourcing.MultiOp{Ops: []sourcing.Operable{
		sourcing.AddParticipant{Name: "Joe"},
		sourcing.AddParticipant{Name: "Joe"},
	}})

	if assert.Error(t, err) {
		assert.Equal(t,
			"couldn't apply operation #1: apply <sourcing.AddParticipant{Name:\"Joe\"}>: participant already exists",
			err.Error())
	}
}
