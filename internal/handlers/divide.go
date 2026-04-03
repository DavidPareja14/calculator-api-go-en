package handlers

import (
	"net/http"

	"github.com/calculator-api-go-en/calculator-api/internal/domain"
)

// Divide handles GET /divide.
func (h *Calculator) Divide(w http.ResponseWriter, r *http.Request) {
	h.calculate(w, r, string(domain.OperationDivide))
}
