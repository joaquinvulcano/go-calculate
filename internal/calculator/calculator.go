package calculator

// Operation representa una operación matemática
type Operation struct {
	Operation string  `json:"operation"`
	Num1      float64 `json:"num1"`
	Num2      float64 `json:"num2"`
}

// Result representa el resultado de una operación
type Result struct {
	Result float64 `json:"result"`
}

// Calculate realiza la operación matemática solicitada
func Calculate(op Operation) (float64, error) {
	switch op.Operation {
	case "add":
		return op.Num1 + op.Num2, nil
	case "subtract":
		return op.Num1 - op.Num2, nil
	case "multiply":
		return op.Num1 * op.Num2, nil
	case "divide":
		if op.Num2 == 0 {
			return 0, ErrDivisionByZero
		}
		return op.Num1 / op.Num2, nil
	default:
		return 0, ErrInvalidOperation
	}
}

// IsValidOperation verifica si la operación es válida
func IsValidOperation(op string) bool {
	switch op {
	case "add", "subtract", "multiply", "divide":
		return true
	default:
		return false
	}
}
