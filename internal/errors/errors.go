package errors

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidAuth  = errors.New("invalid credentials")
	ErrUnauthorized = errors.New("unauthorized")
)
