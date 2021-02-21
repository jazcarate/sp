// Package sourcing contains state and ways to change it
package sourcing

import (
	"errors"
	"fmt"
)

// A ApplyError represents an error when applying Op to a State; defined by the inner Error.
type ApplyError struct {
	PreviousState *State
	Op            Operable
	Err           error
}

func (w *ApplyError) Error() string {
	return fmt.Sprintf("apply <%#v>: %s", w.Op, w.Err)
}

var (
	// ErrAlreadyExists represents an error when trying to add a participant already in the pool.
	ErrAlreadyExists = errors.New("participant already exists")
	// ErrNoparticipant represents an error when trying to remove a participant that is not in the pool.
	ErrNoParticipant = errors.New("that participant does not exist")
)
