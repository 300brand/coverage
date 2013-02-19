package logger

import (
	"fmt"
	"time"
)

type Entry struct {
	Time  time.Time
	Level Level
	Msg   string
}

type Entries []Entry

func (l Entry) String() string {
	return fmt.Sprintf("[%d] %s - %s", l.Level, l.Time, l.Msg)
}
