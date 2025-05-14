package factorymethod

import (
	"testing"
)

func doLog(factory LoggerFactory) {
	factory.CreateLogger().Log("This is a log message")
}

func Test_FactoryMethod_Logger(t *testing.T) {
	doLog(&FileLoggerFactory{})
	doLog(&ConsoleLoggerFactory{})
}
