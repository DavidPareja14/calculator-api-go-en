package domain

import "errors"

var (
	// ErrInvalidOperation is returned when the operation string is not supported.
	ErrInvalidOperation = errors.New("the operation is not supported")

	// ErrDivisionByZero is returned when dividing by zero.
	ErrDivisionByZero = errors.New("Division by zero is not allowed.")
)
