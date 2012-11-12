package parser

import (
	"errors"
	"git.300brand.com/coverage/logger"
	"time"
)

type Time string

var formats = []string{
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822Z,
	time.RFC822,
	time.RFC850,
	time.RFC1123Z,
	time.RFC1123,
	time.RFC3339Nano,
	time.RFC3339,
	// time.Kitchen,
	time.StampNano,
	time.StampMicro,
	time.StampMilli,
	time.Stamp,
}

func (t Time) Time() (ts time.Time) {
	ts, err := Parse(string(t))
	if err != nil {
		logger.Warn(err)
	}
	return
}

func Parse(s string) (t time.Time, err error) {
	for _, layout := range formats {
		if t, err = time.Parse(layout, s); err == nil {
			return
		}
	}
	// As in time.Parse(), return UTC for the first arg, which will come out
	// of the previous calls to time.Parse()
	return t, errors.New("Could not parse " + s)
}
