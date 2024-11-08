package logger

import (
	"log"
	"os"
	"sync"
)

// Logger is the singleton struct for the logger.
type Logger struct {
	*log.Logger
}

var (
	instance *Logger
	once     sync.Once
)

// GetLogger provides a global access point to the Logger instance.
func GetLogger() *Logger {
	once.Do(func() {
		// Initialize the logger instance only once
		instance = &Logger{
			Logger: log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	})
	return instance
}

// Info logs an informational message
func (l *Logger) Info(message string) {
	l.Println("[INFO]", message)
}

// Error logs an error message
func (l *Logger) Error(message string) {
	l.Println("[ERROR]", message)
}
