package errs

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrCheckConstraint = errors.New("check constraint")
)
