package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("expected 5, got %f", result)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    if result != 2 {
        t.Errorf("expected 2, got %f", result)
    }
}

func TestMultiply(t *testing.T) {
    result := Multiply(2, 3)
    if result != 6 {
        t.Errorf("expected 6, got %f", result)
    }
}

func TestDivide(t *testing.T) {
    result, err := Divide(6, 3)
    if err != nil || result != 2 {
        t.Errorf("expected 2, got %f, error: %v", result, err)
    }

    _, err = Divide(1, 0)
    if err == nil {
        t.Error("expected error for division by zero, got none")
    }
}
