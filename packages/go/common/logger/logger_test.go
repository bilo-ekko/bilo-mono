package logger

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name    string
		context string
	}{
		{"With context", "TestContext"},
		{"With empty context", ""},
		{"With special characters", "Test-Context_123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewLogger(tt.context)
			if logger == nil {
				t.Error("NewLogger() returned nil")
			}
			if logger.context != tt.context {
				t.Errorf("NewLogger() context = %v, want %v", logger.context, tt.context)
			}
		})
	}
}

func TestLoggerFormatMessage(t *testing.T) {
	logger := NewLogger("TestContext")

	tests := []struct {
		name         string
		level        LogLevel
		message      string
		wantContains []string
	}{
		{
			name:         "Debug message",
			level:        DEBUG,
			message:      "test debug",
			wantContains: []string{"DEBUG", "TestContext", "test debug"},
		},
		{
			name:         "Info message",
			level:        INFO,
			message:      "test info",
			wantContains: []string{"INFO", "TestContext", "test info"},
		},
		{
			name:         "Warn message",
			level:        WARN,
			message:      "test warning",
			wantContains: []string{"WARN", "TestContext", "test warning"},
		},
		{
			name:         "Error message",
			level:        ERROR,
			message:      "test error",
			wantContains: []string{"ERROR", "TestContext", "test error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := logger.formatMessage(tt.level, tt.message)

			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("formatMessage() = %v, should contain %v", got, want)
				}
			}
		})
	}
}

func TestLoggerMethods(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	logger := NewLogger("TestLogger")

	tests := []struct {
		name         string
		logFunc      func()
		wantContains []string
	}{
		{
			name: "Debug method",
			logFunc: func() {
				buf.Reset()
				logger.Debug("debug message")
			},
			wantContains: []string{"DEBUG", "TestLogger", "debug message"},
		},
		{
			name: "Info method",
			logFunc: func() {
				buf.Reset()
				logger.Info("info message")
			},
			wantContains: []string{"INFO", "TestLogger", "info message"},
		},
		{
			name: "Warn method",
			logFunc: func() {
				buf.Reset()
				logger.Warn("warn message")
			},
			wantContains: []string{"WARN", "TestLogger", "warn message"},
		},
		{
			name: "Error method without error",
			logFunc: func() {
				buf.Reset()
				logger.Error("error message", nil)
			},
			wantContains: []string{"ERROR", "TestLogger", "error message"},
		},
		{
			name: "Error method with error",
			logFunc: func() {
				buf.Reset()
				logger.Error("error message", errors.New("test error"))
			},
			wantContains: []string{"ERROR", "TestLogger", "error message", "test error"},
		},
		{
			name: "Infof method",
			logFunc: func() {
				buf.Reset()
				logger.Infof("formatted %s %d", "message", 123)
			},
			wantContains: []string{"INFO", "TestLogger", "formatted message 123"},
		},
		{
			name: "Errorf method",
			logFunc: func() {
				buf.Reset()
				logger.Errorf("formatted error: %s", "details")
			},
			wantContains: []string{"ERROR", "TestLogger", "formatted error: details"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute the log function
			tt.logFunc()

			// Get the output
			output := buf.String()

			// Check for expected content
			for _, want := range tt.wantContains {
				if !strings.Contains(output, want) {
					t.Errorf("Log output = %v, should contain %v", output, want)
				}
			}
		})
	}
}

func TestLogLevels(t *testing.T) {
	tests := []struct {
		name  string
		level LogLevel
		want  string
	}{
		{"Debug level", DEBUG, "DEBUG"},
		{"Info level", INFO, "INFO"},
		{"Warn level", WARN, "WARN"},
		{"Error level", ERROR, "ERROR"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.level) != tt.want {
				t.Errorf("LogLevel = %v, want %v", tt.level, tt.want)
			}
		})
	}
}

// Benchmark tests
func BenchmarkLoggerInfo(b *testing.B) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	logger := NewLogger("BenchmarkContext")

	for i := 0; i < b.N; i++ {
		logger.Info("benchmark message")
	}
}

func BenchmarkLoggerInfof(b *testing.B) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	logger := NewLogger("BenchmarkContext")

	for i := 0; i < b.N; i++ {
		logger.Infof("benchmark message %d", i)
	}
}

func BenchmarkLoggerError(b *testing.B) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	logger := NewLogger("BenchmarkContext")
	err := errors.New("test error")

	for i := 0; i < b.N; i++ {
		logger.Error("benchmark error", err)
	}
}
