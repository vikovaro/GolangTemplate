package errors

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidAuth           = errors.New("invalid credentials")
	ErrUnauthorized          = errors.New("unauthorized")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")

	ErrInvalidID = errors.New("invalid id")

	ErrAuthorizationHeaderRequired = errors.New("authorization header required")
	ErrInvalidAuthorizationHeader  = errors.New("invalid authorization header")
	ErrInvalidOrExpiredToken       = errors.New("invalid or expired token")

	ErrRoleNotFoundInToken = errors.New("role not found in token")
	ErrForbidden           = errors.New("forbidden")

	ErrForbiddenCannotEditOtherUsersData = errors.New("forbidden: cannot edit other user's data")

	ErrFailedToHashPassword = errors.New("failed to hash password")

	ErrInternalServerError = errors.New("internal server error")
)
