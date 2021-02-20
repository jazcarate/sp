// Package sourcing contains ways to report a state
package sourcing

import (
	"fmt"
	"strings"
)

func (s *State) isEmpty() bool {
	return len(s.participants) == 0
}

// Markdown converts a state to a markdown report.
func (s *State) Markdown() string {
	if s == nil || s.isEmpty() {
		return "🗅 No state"
	}

	return fmt.Sprintf("Participating: %v", strings.Join(s.Participants(), ", "))
}
