package sourcing_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jazcarate/sp/internal/sourcing"
)

func TestRoundtrip_RoundtripEmptyState(t *testing.T) {
	var s *sourcing.State = sourcing.NewState()

	asserRoundtrip(t, s)
}

func TestRoundtrip_RoundtripWithAName(t *testing.T) {
	var s *sourcing.State = sourcing.NewState()
	s.Name = "foo"

	asserRoundtrip(t, s)
}

var now = time.Date(2020, time.April, 10, 14, 20, 10, 22, time.UTC) //nolint:gochecknoglobals // use a single global date thought the testing

func TestRoundtrip_Participant(t *testing.T) {
	var s *sourcing.State
	s, _ = s.Apply(sourcing.AddParticipant{Name: "Joe", PublicKey: "3"}, now)

	asserRoundtrip(t, s)
}

func TestRoundtrip_Split(t *testing.T) {
	var s *sourcing.State
	s, _ = s.Apply(sourcing.AddParticipant{Name: "Joe", PublicKey: "3"}, now)
	s, _ = s.Apply(sourcing.SplitParticipant{Name: "Joe", NewSplit: 10}, now)

	asserRoundtrip(t, s)
}

func TestRoundtrip_Configuration(t *testing.T) {
	var s *sourcing.State
	s, _ = s.Apply(sourcing.Configure{NewConfig: sourcing.All}, now)

	asserRoundtrip(t, s)
}

func TestRoundtrip_Spend(t *testing.T) {
	var s *sourcing.State
	s, _ = s.Apply(sourcing.AddParticipant{Name: "Joe", PublicKey: "3"}, now)
	s, _ = s.Apply(sourcing.Spend{Who: "Joe", Amount: 300}, now)

	asserRoundtrip(t, s)
}

func TestRoundtrip_Transfer(t *testing.T) {
	var s *sourcing.State
	s, _ = s.Apply(sourcing.AddParticipant{Name: "Joe", PublicKey: "3"}, now)
	s, _ = s.Apply(sourcing.AddParticipant{Name: "Bob", PublicKey: "1"}, now)
	s, _ = s.Apply(sourcing.Transfer{From: "Joe", To: "Bob", Amount: 300}, now)

	asserRoundtrip(t, s)
}

func TestRoundtrip_MultiOp(t *testing.T) {
	var s *sourcing.State
	s, _ = s.Apply(sourcing.MultiOp{
		Ops: []sourcing.StateChanger{
			sourcing.AddParticipant{Name: "Joe", PublicKey: "3"},
			sourcing.AddParticipant{Name: "Bob", PublicKey: "1"},
			sourcing.Transfer{From: "Joe", To: "Bob", Amount: 300},
		},
	}, now)

	asserRoundtrip(t, s)
}

func asserRoundtrip(t *testing.T, s *sourcing.State) {
	var b bytes.Buffer

	_ = s.Markdown(&b)
	newS, err := sourcing.Parse(&b)

	assert.Empty(t, err)
	assert.Equal(t, s, newS)
}
