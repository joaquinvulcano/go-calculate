package calculator

import "errors"

var (
	// ErrDivisionByZero se devuelve cuando se intenta dividir por cero
	ErrDivisionByZero = errors.New("division by zero")
	// ErrInvalidOperation se devuelve cuando la operación no es válida
	ErrInvalidOperation = errors.New("invalid operation")
)
