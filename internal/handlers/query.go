package handlers

import (
	"net/http"
	"strconv"
)

func queryFloat64(r *http.Request, name string) (value float64, errMsg string, ok bool) {
	raw := r.URL.Query().Get(name)
	if raw == "" {
		return 0, "The parameter '" + name + "' is required and must be numeric.", false
	}
	v, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, "The parameter '" + name + "' is required and must be numeric.", false
	}
	return v, "", true
}
