package factorymethod

import "fmt"

// Logger is an interface that defines the methods for logging
type Logger interface {
	Log(message string)
}

// FileLogger is a concrete implementation of Logger that logs to a file
type FileLogger struct{}

func (f *FileLogger) Log(message string) {
	// Implementation for logging to a file
	fmt.Println("Log to file: " + message)
}

// ConsoleLogger is a concrete implementation of Logger that logs to the console
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	// Implementation for logging to the console
	fmt.Println("Log to console: " + message)
}

// TODO: Add more logger types as needed

// ----------------------------------------------------------------------------------------------------------------

// LoggerFactory is an interface that defines the method for creating loggers
type LoggerFactory interface {
	CreateLogger() Logger
}

// FileLoggerFactory is a concrete implementation of LoggerFactory that creates FileLogger
type FileLoggerFactory struct{}

func (f *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

// ConsoleLoggerFactory is a concrete implementation of LoggerFactory that creates ConsoleLogger
type ConsoleLoggerFactory struct{}

func (c *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}

// TODO: Add more logger factories as needed
