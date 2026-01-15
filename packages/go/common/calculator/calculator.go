package calculator

import (
	"fmt"
	"time"
)

// CalculateXInput represents the input parameters for the calculation
type CalculateXInput struct {
	Value      float64
	Multiplier *float64
	Offset     *float64
}

// CalculateXOutput represents the result of the calculation
type CalculateXOutput struct {
	Result    float64
	Formula   string
	Timestamp time.Time
}

// CalculateX performs a calculation based on the input parameters
// Formula: (value * multiplier) + offset
func CalculateX(input CalculateXInput) CalculateXOutput {
	multiplier := 2.0
	if input.Multiplier != nil {
		multiplier = *input.Multiplier
	}

	offset := 10.0
	if input.Offset != nil {
		offset = *input.Offset
	}

	result := (input.Value * multiplier) + offset
	formula := fmt.Sprintf("(%.2f * %.2f) + %.2f = %.2f", input.Value, multiplier, offset, result)

	return CalculateXOutput{
		Result:    result,
		Formula:   formula,
		Timestamp: time.Now(),
	}
}

// CalculateXBatch performs batch calculations for multiple values
func CalculateXBatch(values []float64, multiplier *float64, offset *float64) []CalculateXOutput {
	results := make([]CalculateXOutput, len(values))
	for i, value := range values {
		results[i] = CalculateX(CalculateXInput{
			Value:      value,
			Multiplier: multiplier,
			Offset:     offset,
		})
	}
	return results
}
