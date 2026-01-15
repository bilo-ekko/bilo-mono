# Common Go Package

Shared Go utilities for bilo-mono projects.

## Features

### Logger

A simple logging utility with different log levels:

```go
import "github.com/bilo-mono/packages/common/logger"

log := logger.NewLogger("MyApp")
log.Info("Application started")
log.Warn("This is a warning")
log.Error("An error occurred", err)
log.Infof("User %s logged in", username)
```

### Calculator

A utility package that performs calculations based on input parameters:

```go
import "github.com/bilo-mono/packages/common/calculator"

// Single calculation
result := calculator.CalculateX(calculator.CalculateXInput{
    Value:      5,
    Multiplier: &[]float64{3}[0],
    Offset:     &[]float64{10}[0],
})
fmt.Println(result.Formula) // (5.00 * 3.00) + 10.00 = 25.00

// Batch calculation
multiplier := 2.0
offset := 5.0
results := calculator.CalculateXBatch([]float64{1, 2, 3}, &multiplier, &offset)
```

## Installation

This package is part of the bilo-mono monorepo. To use it in your Go application:

```bash
# From your Go app directory
go mod edit -replace github.com/bilo-mono/packages/common=../../packages/go/common
go mod tidy
```

## Development

```bash
# Test the package
go test ./...

# Build the package
go build ./...
```
