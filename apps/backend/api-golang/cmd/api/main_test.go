package main

import (
	"testing"

	"github.com/bilo-mono/packages/common/calculator"
	"github.com/bilo-mono/packages/common/logger"
)

// TestCalculatorIntegration tests that the calculator package works correctly
func TestCalculatorIntegration(t *testing.T) {
	tests := []struct {
		name       string
		input      calculator.CalculateXInput
		wantResult float64
	}{
		{
			name: "Basic calculation with default values",
			input: calculator.CalculateXInput{
				Value: 10.0,
			},
			wantResult: 30.0, // (10 * 2) + 10
		},
		{
			name: "Calculation with custom multiplier",
			input: calculator.CalculateXInput{
				Value:      100.0,
				Multiplier: ptrFloat64(1.5),
				Offset:     ptrFloat64(25.0),
			},
			wantResult: 175.0, // (100 * 1.5) + 25
		},
		{
			name: "Calculation with zero value",
			input: calculator.CalculateXInput{
				Value:      0,
				Multiplier: ptrFloat64(5.0),
				Offset:     ptrFloat64(10.0),
			},
			wantResult: 10.0, // (0 * 5) + 10
		},
		{
			name: "Calculation with negative values",
			input: calculator.CalculateXInput{
				Value:      -10.0,
				Multiplier: ptrFloat64(2.0),
				Offset:     ptrFloat64(-5.0),
			},
			wantResult: -25.0, // (-10 * 2) + (-5)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculator.CalculateX(tt.input)
			if result.Result != tt.wantResult {
				t.Errorf("CalculateX() = %v, want %v", result.Result, tt.wantResult)
			}
			if result.Formula == "" {
				t.Error("CalculateX() formula should not be empty")
			}
			if result.Timestamp.IsZero() {
				t.Error("CalculateX() timestamp should not be zero")
			}
		})
	}
}

// TestCalculatorBatch tests batch calculation functionality
func TestCalculatorBatch(t *testing.T) {
	values := []float64{10.0, 20.0, 30.0}
	multiplier := ptrFloat64(2.0)
	offset := ptrFloat64(5.0)

	results := calculator.CalculateXBatch(values, multiplier, offset)

	if len(results) != len(values) {
		t.Fatalf("Expected %d results, got %d", len(values), len(results))
	}

	expectedResults := []float64{25.0, 45.0, 65.0}
	for i, result := range results {
		if result.Result != expectedResults[i] {
			t.Errorf("Batch result[%d] = %v, want %v", i, result.Result, expectedResults[i])
		}
	}
}

// TestLoggerIntegration tests that the logger package works correctly
func TestLoggerIntegration(t *testing.T) {
	// Create a logger - this is primarily a smoke test
	// as we can't easily capture log output without more infrastructure
	logger := logger.NewLogger("TestContext")

	// These should not panic
	t.Run("Info logging", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Info() panicked: %v", r)
			}
		}()
		logger.Info("Test info message")
	})

	t.Run("Debug logging", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Debug() panicked: %v", r)
			}
		}()
		logger.Debug("Test debug message")
	})

	t.Run("Warn logging", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Warn() panicked: %v", r)
			}
		}()
		logger.Warn("Test warning message")
	})

	t.Run("Error logging", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Error() panicked: %v", r)
			}
		}()
		logger.Error("Test error message", nil)
	})

	t.Run("Formatted logging", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Infof() panicked: %v", r)
			}
		}()
		logger.Infof("Test formatted message: %s %d", "test", 123)
	})
}

// TestPtrFloat64 tests the helper function
func TestPtrFloat64(t *testing.T) {
	value := 42.5
	ptr := ptrFloat64(value)

	if ptr == nil {
		t.Error("ptrFloat64 returned nil")
	}
	if *ptr != value {
		t.Errorf("ptrFloat64() = %v, want %v", *ptr, value)
	}
}

// TestSharedPackagesAccessibility verifies packages can be imported
func TestSharedPackagesAccessibility(t *testing.T) {
	t.Run("Calculator package is accessible", func(t *testing.T) {
		result := calculator.CalculateX(calculator.CalculateXInput{Value: 1.0})
		if result.Result == 0 && result.Formula == "" {
			t.Error("Calculator package may not be functioning correctly")
		}
	})

	t.Run("Logger package is accessible", func(t *testing.T) {
		log := logger.NewLogger("Test")
		if log == nil {
			t.Error("Logger package failed to create logger instance")
		}
	})
}
