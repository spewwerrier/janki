package jankilog

import "errors"

var (
	ErrApiMultipleUsers     = errors.New("Multiple error exists")
	ErrApiUserNoExist       = errors.New("User does not exists")
	ErrApiIncorrectPassword = errors.New("Incorrect password for given username")
	ErrApiIncorrectSession  = errors.New("Incorrect session details")
)

var (
	ErrDbInternalErr = errors.New("Database error in server")
	ErrDbQueryError  = errors.New("Database failed to preapre the query")
	ErrDbExecError   = errors.New("Database cannot execute the sql")
)

var (
	ErrKnobExists   = errors.New("Knob already exists")
	ErrNoKnobExists = errors.New("Knob does not exists")
)
