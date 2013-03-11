package logger

import (
	"code.google.com/p/log4go"
)

type Entries []log4go.LogRecord

func (es *Entries) Add(r log4go.LogRecord) {
	*es = append(*es, r)
}
