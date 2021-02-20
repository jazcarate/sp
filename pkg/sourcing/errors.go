// Package sourcing contains state and ways to change it
package sourcing

import (
	"errors"
	"fmt"
)

// A ApplyError represents an error when applying Op to a State; defined by the inner Error.
type ApplyError struct {
	PreviousState State
	Op            string
	Err           error
}

func (w *ApplyError) Error() string {
	return fmt.Sprintf("apply <%s>: %s", w.Op, w.Err)
}

// A ParticipantError represents an error when dealing with adding or removing participants.
type ParticipantError struct {
	ProblematicParticipant name
	Err                    error
}

func (w *ParticipantError) Error() string {
	return fmt.Sprintf("participant <%s>: %s", w.ProblematicParticipant, w.Err)
}

var (
	errAlreadyExists = errors.New("participant already exists")
	errNoparticipant = errors.New("that participant does not exist")
)

func wrap(state State, op string, err error) (State, *ApplyError) {
	return state, &ApplyError{
		PreviousState: state,
		Op:            op,
		Err:           err,
	}
}

func participantError(p name, err error) *ParticipantError {
	return &ParticipantError{
		ProblematicParticipant: p,
		Err:                    err,
	}
}
