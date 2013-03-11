package logger

import (
	"code.google.com/p/log4go"
)

type LogEntries []log4go.LogRecord

func (es *LogEntries) Add(r log4go.LogRecord) {
	*es = append(*es, r)
}
