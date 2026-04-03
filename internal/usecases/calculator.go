package usecases

import (
	"github.com/calculator-api-go-en/calculator-api/internal/domain"
)

// CalculatorUseCase performs arithmetic calculations for supported operations.
type CalculatorUseCase interface {
	Calculate(operation string, a, b float64) (float64, error)
}

// calculator implements CalculatorUseCase.
type calculator struct{}

// NewCalculator returns a CalculatorUseCase implementation.
func NewCalculator() CalculatorUseCase {
	return &calculator{}
}

// Calculate runs the requested operation on a and b.
func (c *calculator) Calculate(operation string, a, b float64) (float64, error) {
	op, err := domain.ParseOperation(operation)
	if err != nil {
		return 0, err
	}

	switch op {
	case domain.OperationAdd:
		return a + b, nil
	case domain.OperationSubtract:
		return a - b, nil
	case domain.OperationMultiply:
		return a * b, nil
	case domain.OperationDivide:
		if b == 0 {
			return 0, domain.ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, domain.ErrInvalidOperation
	}
}
