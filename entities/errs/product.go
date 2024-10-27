package errs

import "errors"

var (
	// product not found
	ErrProductNotFound = errors.New("product not found")
	// errors.New("insufficient stock")
	ErrorInsufficientStock = errors.New("insufficient stock")
)
