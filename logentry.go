package coverage

import (
	"fmt"
	"git.300brand.com/coverage/logger"
	"time"
)

type LogEntry struct {
	Time  time.Time
	Level logger.Level
	Msg   string
}

func (l LogEntry) String() {
	return fmt.Sprintf("[%d] %s %s", l.Level, l.Time, l.Msg)
}

type LogEntries []LogEntry
