package handlers

import (
	"net/http"

	"github.com/calculator-api-go-en/calculator-api/internal/domain"
)

// Multiply handles GET /multiply.
func (h *Calculator) Multiply(w http.ResponseWriter, r *http.Request) {
	h.calculate(w, r, string(domain.OperationMultiply))
}
