package logger

import (
	"fmt"
	"log"
	"os"
)

type Level int

const (
	_DEBUG Level = iota
	_INFO
	_WARN
	_FATAL
	_PANIC
	flags = log.LstdFlags | log.Lshortfile
)

var loggers = [...]*log.Logger{
	_DEBUG: log.New(os.Stderr, "DEBUG   ", flags),
	_INFO:  log.New(os.Stderr, "INFO    ", flags),
	_WARN:  log.New(os.Stderr, "WARNING ", flags),
	_FATAL: log.New(os.Stderr, "FATAL   ", flags),
	_PANIC: log.New(os.Stderr, "PANIC   ", flags),
}

func Print(level int, v ...interface{}) {
	loggers[level].Output(2, fmt.Sprintln(v...))
}

func Debug(v ...interface{}) {
	loggers[_DEBUG].Output(2, fmt.Sprintln(v...))
}

func Debugf(format string, v ...interface{}) {
	loggers[_DEBUG].Output(2, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	loggers[_WARN].Output(2, fmt.Sprintln(v...))
}

func Warnf(format string, v ...interface{}) {
	loggers[_WARN].Output(2, fmt.Sprintf(format, v...))
}
