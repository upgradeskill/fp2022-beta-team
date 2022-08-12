package errs

import "errors"

var (
	ErrValidPassword = errors.New("not accepted password")
)
