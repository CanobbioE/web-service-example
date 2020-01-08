package infrastructure

import (
	"io"
	"log"
)

// BoringLogger is a wrapper for the standard logger.
type BoringLogger struct {
	logger *log.Logger
}

// Log prints the given message.
func (bl BoringLogger) Log(msg string) {
	bl.logger.Println(msg)
}

// NewBoringLogger returns a new instance of the boring logger.
func NewBoringLogger(out io.Writer) *BoringLogger {
	return &BoringLogger{
		logger: log.New(out, "[BoringLogger]", log.LstdFlags),
	}
}
