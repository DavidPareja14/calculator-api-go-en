package handlers

import "net/http"

// Health handles GET /health.
func (h *Calculator) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
