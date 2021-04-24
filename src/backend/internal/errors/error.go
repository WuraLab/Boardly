package errors

import (
	"github.com/pkg/errors"
)

type StatusError struct {
	ErrMessage string
}

// Error implements the Error interface.
func (e StatusError) Error() string {
	return e.ErrMessage
}

func Is(errorA, errorB error) bool {
	return errors.Is(errorA, errorB)
}

func NewBadRequest(msg string) StatusError {
	return StatusError{ErrMessage: msg}
}
