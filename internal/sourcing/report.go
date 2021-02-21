// Package sourcing contains ways to report a state
package sourcing

import (
	"fmt"
	"strings"
)

// Markdown converts a state to a markdown report.
func (s *State) Markdown() string {
	participants := s.Participants()
	if s == nil || len(participants) == 0 {
		return "🗅 No state"
	}

	return fmt.Sprintf("Participating: %v", strings.Join(participants, ", "))
}
