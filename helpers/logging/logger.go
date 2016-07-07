// Package logging provides primitives for configuring and creating logs.
package logging

import (
	"log"
	"os"
)

var logger *log.Logger

// GetLogger returns the active logging object for creating new messages.
func GetLogger() *log.Logger {
	if logger == nil {
		logger = log.New(os.Stderr, "", log.LstdFlags)
	}
	return logger
}
