package sourcing_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jazcarate/sp/internal/sourcing"
	"github.com/jazcarate/sp/internal/trianglem"
)

func TestParticipant_EmptyStateHasNoParticipants(t *testing.T) {
	s := sourcing.NewState()

	assert.Empty(t, s.Participants)
}

func TestParticipant_AddingOneAddsItToTheList(t *testing.T) {
	var s *sourcing.State
	s, err := s.Apply(sourcing.AddParticipant{Name: "Joe", PublicKey: "1"}, time.Now())

	assert.Empty(t, err)
	assert.Len(t, s.Participants, 1)
	assert.Equal(t, s.Participants[0].Name, "Joe")
}

func TestParticipant_EnablingAnUnexistingErrors(t *testing.T) {
	var s *sourcing.State
	_, err := s.Apply(sourcing.SplitParticipant{Name: "Joe", NewSplit: 3}, time.Now())

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
	}}, time.Now())

	if assert.Error(t, err) {
		assert.Equal(t,
			"couldn't apply operation #1: apply <sourcing.AddParticipant{Name:\"Joe\", PublicKey:\"2\"}>: participant already exists",
			err.Error())
	}
}

func TestParticipantBalance_AddingAParticipantStartsWithA0Balance(t *testing.T) {
	var s *sourcing.State

	s, err := s.Apply(sourcing.AddParticipant{Name: "Joe", PublicKey: "1"}, time.Now())
	val := s.Balance.Get(0, 0)

	assert.Empty(t, err)
	assert.Empty(t, val)
}

func TestTransfer_ErrorWhenParticipantDoesNotExistTo(t *testing.T) {
	var s *sourcing.State
	s, _ = s.Apply(sourcing.AddParticipant{Name: "Joe", PublicKey: "1"}, time.Now())

	_, err := s.Apply(sourcing.Transfer{From: "Joe", To: "Ben", Amount: 10}, time.Now())

	if assert.Error(t, err) {
		assert.Equal(t,
			"apply <sourcing.Transfer{From:\"Joe\", To:\"Ben\", Amount:10}>: transfer to: that participant does not exist",
			err.Error())
	}
}

func TestTransfer_ErrorWhenParticipantDoesNotExistFrom(t *testing.T) {
	var s *sourcing.State

	_, err := s.Apply(sourcing.Transfer{From: "Joe", To: "Ben", Amount: 10}, time.Now())

	if assert.Error(t, err) {
		assert.Equal(t,
			"apply <sourcing.Transfer{From:\"Joe\", To:\"Ben\", Amount:10}>: transfer from: that participant does not exist",
			err.Error())
	}
}

func TestTransfer_ChangesBalance(t *testing.T) {
	var s *sourcing.State

	s, _ = s.Apply(sourcing.MultiOp{Ops: []sourcing.StateChanger{
		sourcing.AddParticipant{Name: "Joe", PublicKey: "1"},
		sourcing.AddParticipant{Name: "Ben", PublicKey: "2"},
	}}, time.Now())

	s, err := s.Apply(sourcing.Transfer{From: "Joe", To: "Ben", Amount: 10}, time.Now())
	val := s.Balance.Get(1, 0)

	assert.Empty(t, err)
	assert.Equal(t, 10, val)
}

func TestTransfer_TransferBack(t *testing.T) {
	var s *sourcing.State

	s, _ = s.Apply(sourcing.MultiOp{Ops: []sourcing.StateChanger{
		sourcing.AddParticipant{Name: "Joe", PublicKey: "1"},
		sourcing.AddParticipant{Name: "Ben", PublicKey: "2"},
	}}, time.Now())

	s, _ = s.Apply(sourcing.Transfer{From: "Joe", To: "Ben", Amount: 10}, time.Now())
	s, err := s.Apply(sourcing.Transfer{From: "Ben", To: "Joe", Amount: 10}, time.Now())
	val := s.Balance.Get(1, 0)

	assert.Empty(t, err)
	assert.Equal(t, 0, val)
}

func TestTransfer_TransitiveBalance(t *testing.T) {
	var s *sourcing.State

	s, _ = s.Apply(sourcing.MultiOp{Ops: []sourcing.StateChanger{
		sourcing.AddParticipant{Name: "Joe", PublicKey: "1"},
		sourcing.AddParticipant{Name: "Ben", PublicKey: "2"},
		sourcing.AddParticipant{Name: "Bob", PublicKey: "3"},
	}}, time.Now())

	const (
		joe = iota
		ben
		bob
	)

	s, _ = s.Apply(sourcing.Transfer{From: "Joe", To: "Ben", Amount: 10}, time.Now())
	s, _ = s.Apply(sourcing.Transfer{From: "Bob", To: "Joe", Amount: 7}, time.Now())

	assertBalance(t, s.Balance, 7, ben, bob)
	assertBalance(t, s.Balance, 3, ben, joe)
	assertBalance(t, s.Balance, 0, bob, joe)
}

func TestTransfer_TransitiveBalancePartial(t *testing.T) {
	t.Skip("Algorithm is not correctly implemented")

	var s *sourcing.State

	s, _ = s.Apply(sourcing.MultiOp{Ops: []sourcing.StateChanger{
		sourcing.AddParticipant{Name: "Joe", PublicKey: "1"},
		sourcing.AddParticipant{Name: "Ben", PublicKey: "2"},
		sourcing.AddParticipant{Name: "Bob", PublicKey: "3"},
	}}, time.Now())

	const (
		joe = iota
		ben
		bob
	)

	s, _ = s.Apply(sourcing.Transfer{From: "Bob", To: "Joe", Amount: 7}, time.Now())
	s, _ = s.Apply(sourcing.Transfer{From: "Ben", To: "Joe", Amount: 6}, time.Now())
	s, _ = s.Apply(sourcing.Transfer{From: "Joe", To: "Ben", Amount: 10}, time.Now())

	assertBalance(t, s.Balance, 4, ben, bob)
	assertBalance(t, s.Balance, 3, ben, joe)
	assertBalance(t, s.Balance, 0, bob, joe)
}

func assertBalance(t *testing.T, m *trianglem.M, expected, from, to int) {
	val := m.Get(from, to)

	assert.Equal(t, expected, val)
}
