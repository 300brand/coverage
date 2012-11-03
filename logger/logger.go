package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	DEBUG = iota
	WARN
	FATAL
	PANIC
)

var (
	flags    = log.LstdFlags | log.Lshortfile
	loggers  = make(map[int]*log.Logger, len(prefixes))
	prefixes = map[int]string{
		DEBUG: "DEBUG   ",
		WARN:  "WARNING ",
		FATAL: "FATAL   ",
		PANIC: "PANIC   ",
	}
)

func Debug(v ...interface{}) {
	getLogger(DEBUG).Output(2, fmt.Sprintln(v...))
}

func Debugf(format string, v ...interface{}) {
	getLogger(DEBUG).Output(2, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	getLogger(WARN).Output(2, fmt.Sprintln(v...))
}

func Warnf(format string, v ...interface{}) {
	getLogger(WARN).Output(2, fmt.Sprintf(format, v...))
}

func getLogger(level int) *log.Logger {
	if _, ok := loggers[level]; !ok {
		loggers[level] = log.New(os.Stderr, prefixes[level], flags)
	}
	return loggers[level]
}
