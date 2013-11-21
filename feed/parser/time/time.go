package time

import (
	"errors"
	"strings"
	"time"
)

type Time string

// time.Kitchen not used.
var formats = []string{
	time.ANSIC,                   // "Mon Jan _2 15:04:05 2006"
	time.UnixDate,                // "Mon Jan _2 15:04:05 MST 2006"
	time.RubyDate,                // "Mon Jan 02 15:04:05 -0700 2006"
	time.RFC822,                  // "02 Jan 06 15:04 MST"
	time.RFC822Z,                 // "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	time.RFC850,                  // "Monday, 02-Jan-06 15:04:05 MST"
	time.RFC1123,                 // "Mon, 02 Jan 2006 15:04:05 MST"
	time.RFC1123Z,                // "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	time.RFC3339,                 // "2006-01-02T15:04:05Z07:00"
	time.RFC3339Nano,             // "2006-01-02T15:04:05.999999999Z07:00"
	time.Stamp,                   // "Jan _2 15:04:05"
	time.StampMilli,              // "Jan _2 15:04:05.000"
	time.StampMicro,              // "Jan _2 15:04:05.000000"
	time.StampNano,               // "Jan _2 15:04:05.000000000"
	"Mon, 2 Jan 2006  15:04:05",  // A.M. Best Company
	"Mon, 02 Jan 2006 15:04 MST", // CNBC
}

// Parses the incoming argument into a time.Time value. Valid formats include
// those found in the time package.
func (t Time) Parse() (ts time.Time, err error) {
	s := strings.TrimSpace(string(t))
	// Support for Bitpipe timestamps using 24:00:00
	s = strings.Replace(s, " 24:", " 00:", 1)
	for _, layout := range formats {
		if ts, err = time.Parse(layout, s); err == nil {
			return
		}
	}
	// As in time.Parse(), return UTC for the first arg, which will come out
	// of the previous calls to time.Parse()
	return ts, errors.New("Could not parse " + s)
}

func (t Time) String() string {
	return string(t)
}

// Returns the time.Time value of the string version of time. If the value is
// invalid, a warning is pushed through the logger and the zero-time in UTC
// returned
func (t Time) Time() (ts time.Time) {
	ts, err := t.Parse()
	if err != nil {
		//logger.Warn(err)
	}
	return
}
