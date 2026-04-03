package handlers

import (
	"errors"
	"net/http"

	"github.com/calculator-api-go-en/calculator-api/internal/domain"
	"github.com/calculator-api-go-en/calculator-api/internal/usecases"
)

// Calculator exposes HTTP handlers that delegate to CalculatorUseCase.
type Calculator struct {
	uc usecases.CalculatorUseCase
}

// NewCalculator constructs Calculator handlers with the given use case.
func NewCalculator(uc usecases.CalculatorUseCase) *Calculator {
	return &Calculator{uc: uc}
}

func (h *Calculator) calculate(w http.ResponseWriter, r *http.Request, operation string) {
	a, msg, ok := queryFloat64(r, "a")
	if !ok {
		writeError(w, http.StatusBadRequest, msg)
		return
	}
	b, msg, ok := queryFloat64(r, "b")
	if !ok {
		writeError(w, http.StatusBadRequest, msg)
		return
	}

	result, err := h.uc.Calculate(operation, a, b)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrDivisionByZero):
			writeError(w, http.StatusBadRequest, err.Error())
		case errors.Is(err, domain.ErrInvalidOperation):
			writeError(w, http.StatusBadRequest, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "Internal server error.")
		}
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"operation": operation,
		"a":         a,
		"b":         b,
		"result":    result,
	})
}
