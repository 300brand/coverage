package logger

import (
	"fmt"
	"time"
)

type Entry struct {
	Time    time.Time
	Level   Level
	Message string
}

func (e Entry) LogString() string {
	return fmt.Sprintf("%s %s", e.Level, e.Message)
}
