package logger

import (
	"io"
	"log"
)

var (
	loggingEnabled = false
)

func DisableLogging() {
	loggingEnabled = false
}

func EnableLogging(w io.Writer) {
	loggingEnabled = true
	log.SetOutput(w)
}
