package error

import "errors"

var(
	ErrNotFound=errors.New("user not found")
	ErrInvalidID    = errors.New("invalid user id")
	ErrValidation   = errors.New("validation failed")
)