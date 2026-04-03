package handlers

import (
	"net/http"

	"github.com/calculator-api-go-en/calculator-api/internal/domain"
)

// Subtract handles GET /subtract.
func (h *Calculator) Subtract(w http.ResponseWriter, r *http.Request) {
	h.calculate(w, r, string(domain.OperationSubtract))
}
