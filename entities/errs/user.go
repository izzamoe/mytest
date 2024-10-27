package errs

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrPasswordMismatch   = errors.New("passwords do not match")
	ErrorUserExists       = errors.New("user with this email already exists")
)
