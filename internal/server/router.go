package server

import (
	"net/http"

	"github.com/calculator-api-go-en/calculator-api/internal/handlers"
	"github.com/calculator-api-go-en/calculator-api/internal/usecases"
	"github.com/go-chi/chi/v5"
)

// NewRouter builds the HTTP handler with Chi routes and injected dependencies.
func NewRouter(uc usecases.CalculatorUseCase) http.Handler {
	r := chi.NewRouter()
	h := handlers.NewCalculator(uc)

	r.Get("/health", h.Health)
	r.Get("/add", h.Add)
	r.Get("/subtract", h.Subtract)
	r.Get("/multiply", h.Multiply)
	r.Get("/divide", h.Divide)

	return r
}
