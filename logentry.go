package coverage

import (
	"git.300brand.com/coverage/logger"
	"time"
)

type LogEntry struct {
	Time  time.Time
	Level logger.Level
	Msg   string
}

type LogEntries []LogEntry
