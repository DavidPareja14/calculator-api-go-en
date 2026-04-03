package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/calculator-api-go-en/calculator-api/internal/usecases"
)

func TestRouter_Add_Success(t *testing.T) {
	uc := usecases.NewCalculator()
	handler := NewRouter(uc)

	req := httptest.NewRequest(http.MethodGet, "/add?a=10&b=5", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusOK)
	}

	var body map[string]any
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body["operation"] != "add" || body["result"] != float64(15) {
		t.Fatalf("body = %#v", body)
	}
}

func TestRouter_Add_InvalidParameter(t *testing.T) {
	uc := usecases.NewCalculator()
	handler := NewRouter(uc)

	req := httptest.NewRequest(http.MethodGet, "/add?a=not-a-number&b=2", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusBadRequest)
	}

	var body struct {
		Error string `json:"error"`
	}
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body.Error == "" {
		t.Fatal("expected error message")
	}
}

func TestRouter_Divide_DivisionByZero(t *testing.T) {
	uc := usecases.NewCalculator()
	handler := NewRouter(uc)

	req := httptest.NewRequest(http.MethodGet, "/divide?a=1&b=0", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusBadRequest)
	}

	var body struct {
		Error string `json:"error"`
	}
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	want := "Division by zero is not allowed."
	if body.Error != want {
		t.Fatalf("error = %q, want %q", body.Error, want)
	}
}
