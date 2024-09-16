package jankilog

import "errors"

var ErrApiMultipleUsers = errors.New("Multiple error exists")
var ErrApiUserNoExist = errors.New("User does not exists")
var ErrApiIncorrectPassword = errors.New("Incorrect password for given username")
var ErrApiIncorrectSession = errors.New("Incorrect session details")

var ErrDbInternalErr = errors.New("Database error in server")
var ErrDbQueryError = errors.New("Database failed to preapre the query")
var ErrDbExecError = errors.New("Database cannot execute the sql")
