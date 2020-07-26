package logging

import (
	"fmt"
)

// AppLogger defines the functions
type AppLogger interface {
	Logf(format string, args ...interface{})
}

// FakeLogger for testing
type FakeLogger struct{}

// Logf does nothing
func (fl *FakeLogger) Logf(format string, args ...interface{}) {}

// Logger used in the App
type Logger struct {
	Prefix string
}

// Logf prints the log message
func (log *Logger) Logf(format string, args ...interface{}) {
	fmt.Print(log.Prefix)
	fmt.Printf(format, args...)
	fmt.Print("\n")
}

// CreateLogger creates a new Logger
func CreateLogger(prefix string) *Logger {
	return &Logger{
		Prefix: prefix,
	}
}
