package usecases

import (
	"errors"
	"math"
	"testing"

	"github.com/calculator-api-go-en/calculator-api/internal/domain"
)

func TestCalculator_Calculate_Add(t *testing.T) {
	c := NewCalculator()
	got, err := c.Calculate(string(domain.OperationAdd), 10, 5)
	if err != nil {
		t.Fatalf("Calculate: %v", err)
	}
	if got != 15 {
		t.Fatalf("result = %v, want 15", got)
	}
}

func TestCalculator_Calculate_Subtract(t *testing.T) {
	c := NewCalculator()
	got, err := c.Calculate(string(domain.OperationSubtract), 10, 3)
	if err != nil {
		t.Fatalf("Calculate: %v", err)
	}
	if got != 7 {
		t.Fatalf("result = %v, want 7", got)
	}
}

func TestCalculator_Calculate_Multiply(t *testing.T) {
	c := NewCalculator()
	got, err := c.Calculate(string(domain.OperationMultiply), 6, 7)
	if err != nil {
		t.Fatalf("Calculate: %v", err)
	}
	if got != 42 {
		t.Fatalf("result = %v, want 42", got)
	}
}

func TestCalculator_Calculate_Divide(t *testing.T) {
	c := NewCalculator()
	got, err := c.Calculate(string(domain.OperationDivide), 20, 4)
	if err != nil {
		t.Fatalf("Calculate: %v", err)
	}
	if got != 5 {
		t.Fatalf("result = %v, want 5", got)
	}
}

func TestCalculator_Calculate_DivisionByZero(t *testing.T) {
	c := NewCalculator()
	_, err := c.Calculate(string(domain.OperationDivide), 1, 0)
	if err == nil {
		t.Fatal("expected error for division by zero")
	}
	if !errors.Is(err, domain.ErrDivisionByZero) {
		t.Fatalf("errors.Is division by zero: got %v", err)
	}
}

func TestCalculator_Calculate_InvalidOperation(t *testing.T) {
	c := NewCalculator()
	_, err := c.Calculate("modulo", 1, 1)
	if err == nil {
		t.Fatal("expected error for invalid operation")
	}
	if !errors.Is(err, domain.ErrInvalidOperation) {
		t.Fatalf("errors.Is invalid operation: got %v", err)
	}
}

func TestCalculator_Calculate_FloatOperands(t *testing.T) {
	c := NewCalculator()
	got, err := c.Calculate(string(domain.OperationAdd), 0.1, 0.2)
	if err != nil {
		t.Fatalf("Calculate: %v", err)
	}
	if math.Abs(got-0.3) > 1e-9 {
		t.Fatalf("result = %v, want ~0.3", got)
	}
}
