package logger

import (
	"fmt"
	"log"
	"time"
)

type LogLevel string

const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
)

type Logger struct {
	context string
}

// NewLogger creates a new logger instance with a given context
func NewLogger(context string) *Logger {
	return &Logger{
		context: context,
	}
}

func (l *Logger) formatMessage(level LogLevel, message string) string {
	timestamp := time.Now().Format(time.RFC3339)
	return fmt.Sprintf("[%s] [%s] [%s] %s", timestamp, level, l.context, message)
}

// Debug logs a debug message
func (l *Logger) Debug(message string) {
	log.Println(l.formatMessage(DEBUG, message))
}

// Info logs an info message
func (l *Logger) Info(message string) {
	log.Println(l.formatMessage(INFO, message))
}

// Warn logs a warning message
func (l *Logger) Warn(message string) {
	log.Println(l.formatMessage(WARN, message))
}

// Error logs an error message
func (l *Logger) Error(message string, err error) {
	fullMessage := message
	if err != nil {
		fullMessage = fmt.Sprintf("%s - %v", message, err)
	}
	log.Println(l.formatMessage(ERROR, fullMessage))
}

// Infof logs an info message with formatting
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

// Errorf logs an error message with formatting
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...), nil)
}
