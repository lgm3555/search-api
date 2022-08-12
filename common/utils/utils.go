package utils

import (
	"log"
	"os"
)

var Logger = NewLogger()

const (
	Info  = "INFO: "
	Debug = "DEBUG: "
	Trace = "TRACE: "
	Warn  = "WARN: "
	Error = "ERROR: "
)

func NewLogger() *log.Logger {

	logger := log.New(os.Stdout, log.Prefix(), log.LstdFlags)

	return logger
}
