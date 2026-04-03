package handlers

import (
	"net/http"

	"github.com/calculator-api-go-en/calculator-api/internal/domain"
)

// Add handles GET /add.
func (h *Calculator) Add(w http.ResponseWriter, r *http.Request) {
	h.calculate(w, r, string(domain.OperationAdd))
}
