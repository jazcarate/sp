package sourcing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jazcarate/sp/internal/sourcing"
)

func TestParticipant_EmptyStateHasNoParticipants(t *testing.T) {
	s := sourcing.NewState()

	assert.Empty(t, s.Participants)
}

func TestParticipant_AddingOneAddsItToTheList(t *testing.T) {
	var s *sourcing.State
	s, err := s.Apply(sourcing.AddParticipant{Name: "Joe", PublicKey: "1"})

	assert.Empty(t, err)
	assert.Len(t, s.Participants, 1)
	assert.Equal(t, s.Participants[0].Name, "Joe")
}

func TestParticipant_EnablingAnUnexistingErrors(t *testing.T) {
	var s *sourcing.State
	_, err := s.Apply(sourcing.SplitParticipant{Name: "Joe", NewSplit: 3})

	if assert.Error(t, err) {
		assert.Equal(t,
			"apply <sourcing.SplitParticipant{Name:\"Joe\", NewSplit:3}>: that participant does not exist",
			err.Error())
	}
}

func TestParticipantError_AddingDuplicateParticipantsErrors(t *testing.T) {
	var s *sourcing.State

	_, err := s.Apply(sourcing.MultiOp{Ops: []sourcing.StateChanger{
		sourcing.AddParticipant{Name: "Joe", PublicKey: "1"},
		sourcing.AddParticipant{Name: "Joe", PublicKey: "2"},
	}})

	if assert.Error(t, err) {
		assert.Equal(t,
			"couldn't apply operation #1: apply <sourcing.AddParticipant{Name:\"Joe\", PublicKey:\"2\"}>: participant already exists",
			err.Error())
	}
}

func TestParticipantBalance_AddingAParticipantStartsWithA0Balance(t *testing.T) {
	var s *sourcing.State

	s, err := s.Apply(sourcing.AddParticipant{Name: "Joe", PublicKey: "1"})
	val, getErr := s.Balance.Get(0, 0)

	assert.Empty(t, err)
	assert.Empty(t, getErr)
	assert.Empty(t, val)
}
