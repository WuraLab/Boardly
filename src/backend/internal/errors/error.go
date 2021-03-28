package errors

type StatusError struct {
	ErrMessage string
}

// Error implements the Error interface.
func (e StatusError) Error() string {
	return e.ErrMessage
}

func NewBadRequest(msg string) 	StatusError{
  return StatusError{ErrMessage: msg}
}