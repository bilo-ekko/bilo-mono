package calculator

import (
	"testing"
)

func TestCalculateX(t *testing.T) {
	tests := []struct {
		name       string
		input      CalculateXInput
		wantResult float64
	}{
		{
			name: "Default multiplier and offset",
			input: CalculateXInput{
				Value: 10.0,
			},
			wantResult: 30.0, // (10 * 2) + 10
		},
		{
			name: "Custom multiplier only",
			input: CalculateXInput{
				Value:      10.0,
				Multiplier: ptrFloat64(3.0),
			},
			wantResult: 40.0, // (10 * 3) + 10
		},
		{
			name: "Custom offset only",
			input: CalculateXInput{
				Value:  10.0,
				Offset: ptrFloat64(5.0),
			},
			wantResult: 25.0, // (10 * 2) + 5
		},
		{
			name: "Custom multiplier and offset",
			input: CalculateXInput{
				Value:      100.0,
				Multiplier: ptrFloat64(1.5),
				Offset:     ptrFloat64(25.0),
			},
			wantResult: 175.0, // (100 * 1.5) + 25
		},
		{
			name: "Zero value",
			input: CalculateXInput{
				Value: 0,
			},
			wantResult: 10.0, // (0 * 2) + 10
		},
		{
			name: "Negative value",
			input: CalculateXInput{
				Value:      -10.0,
				Multiplier: ptrFloat64(2.0),
				Offset:     ptrFloat64(5.0),
			},
			wantResult: -15.0, // (-10 * 2) + 5
		},
		{
			name: "Zero multiplier",
			input: CalculateXInput{
				Value:      100.0,
				Multiplier: ptrFloat64(0),
				Offset:     ptrFloat64(50.0),
			},
			wantResult: 50.0, // (100 * 0) + 50
		},
		{
			name: "Fractional values",
			input: CalculateXInput{
				Value:      2.5,
				Multiplier: ptrFloat64(3.2),
				Offset:     ptrFloat64(1.8),
			},
			wantResult: 9.8, // (2.5 * 3.2) + 1.8
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateX(tt.input)

			// Check result
			if got.Result != tt.wantResult {
				t.Errorf("CalculateX() Result = %v, want %v", got.Result, tt.wantResult)
			}

			// Check formula is not empty
			if got.Formula == "" {
				t.Error("CalculateX() Formula should not be empty")
			}

			// Check timestamp is set
			if got.Timestamp.IsZero() {
				t.Error("CalculateX() Timestamp should not be zero")
			}
		})
	}
}

func TestCalculateXBatch(t *testing.T) {
	tests := []struct {
		name           string
		values         []float64
		multiplier     *float64
		offset         *float64
		expectedLength int
		expectedFirst  float64
		expectedLast   float64
	}{
		{
			name:           "Batch with default values",
			values:         []float64{10.0, 20.0, 30.0},
			multiplier:     nil,
			offset:         nil,
			expectedLength: 3,
			expectedFirst:  30.0, // (10 * 2) + 10
			expectedLast:   70.0, // (30 * 2) + 10
		},
		{
			name:           "Batch with custom multiplier and offset",
			values:         []float64{5.0, 10.0, 15.0},
			multiplier:     ptrFloat64(3.0),
			offset:         ptrFloat64(2.0),
			expectedLength: 3,
			expectedFirst:  17.0, // (5 * 3) + 2
			expectedLast:   47.0, // (15 * 3) + 2
		},
		{
			name:           "Empty batch",
			values:         []float64{},
			multiplier:     ptrFloat64(2.0),
			offset:         ptrFloat64(5.0),
			expectedLength: 0,
		},
		{
			name:           "Single value batch",
			values:         []float64{42.0},
			multiplier:     ptrFloat64(1.0),
			offset:         ptrFloat64(0),
			expectedLength: 1,
			expectedFirst:  42.0, // (42 * 1) + 0
			expectedLast:   42.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateXBatch(tt.values, tt.multiplier, tt.offset)

			if len(got) != tt.expectedLength {
				t.Errorf("CalculateXBatch() length = %v, want %v", len(got), tt.expectedLength)
			}

			if tt.expectedLength > 0 {
				if got[0].Result != tt.expectedFirst {
					t.Errorf("CalculateXBatch() first result = %v, want %v", got[0].Result, tt.expectedFirst)
				}
				if got[len(got)-1].Result != tt.expectedLast {
					t.Errorf("CalculateXBatch() last result = %v, want %v", got[len(got)-1].Result, tt.expectedLast)
				}

				// Verify all results have formula and timestamp
				for i, result := range got {
					if result.Formula == "" {
						t.Errorf("CalculateXBatch() result[%d] Formula should not be empty", i)
					}
					if result.Timestamp.IsZero() {
						t.Errorf("CalculateXBatch() result[%d] Timestamp should not be zero", i)
					}
				}
			}
		})
	}
}

// Helper function to create pointer to float64
func ptrFloat64(v float64) *float64 {
	return &v
}

// Benchmark tests
func BenchmarkCalculateX(b *testing.B) {
	input := CalculateXInput{
		Value:      100.0,
		Multiplier: ptrFloat64(1.5),
		Offset:     ptrFloat64(25.0),
	}

	for i := 0; i < b.N; i++ {
		CalculateX(input)
	}
}

func BenchmarkCalculateXBatch(b *testing.B) {
	values := []float64{10.0, 20.0, 30.0, 40.0, 50.0}
	multiplier := ptrFloat64(2.0)
	offset := ptrFloat64(5.0)

	for i := 0; i < b.N; i++ {
		CalculateXBatch(values, multiplier, offset)
	}
}
