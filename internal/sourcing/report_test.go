package sourcing_test

import (
	"bytes"
	"testing"

	"github.com/jazcarate/sp/internal/sourcing"
	"github.com/stretchr/testify/assert"
)

func TestMarkdown_NilState(t *testing.T) {
	var (
		s *sourcing.State
		b bytes.Buffer
	)

	err := s.Markdown(&b)

	assert.Empty(t, err)
	assert.Equal(t, "# Split Chain\n\n"+
		"🗅 Starting out a new Split Chain and don't know \"what now?\".\n\n"+
		"No problem! Check the [docs](https://github.com/jazcarate/sp/blob/master/docs/new_sp_now_what.md)\n\n"+
		"## Operations\n"+
		"Current trust configuration: **Trust** [(ℹ)]"+
		"(https://github.com/jazcarate/sp/blob/master/docs/understanding_a_report.md.md#trust)\n\n"+
		"### Log\n"+
		"🌈 Fresh new 🌈\n",
		b.String())
}
