package time

import (
	"fmt"
	"testing"
	"time"
)

var (
	now   = time.Now()
	times []Time
)

func init() {
	for _, layout := range formats {
		times = append(times, Time(now.Format(layout)))
	}
	times = append(times, "Wed, 12 Jun 2013 24:00:00 EDT")
}

func TestFormats(t *testing.T) {
	for i, ts := range times {
		if _, err := ts.Parse(); err != nil {
			t.Errorf("[%02d] (%s): %s", i, ts, err)
		}
	}
}

func TestPaddedTime(t *testing.T) {
	for i, ts := range times {
		ts = Time(fmt.Sprintf("\n \t   %s      \n   ", ts))
		if _, err := ts.Parse(); err != nil {
			t.Errorf("[%02d] (%s):  %s", i, ts, err)
		}
	}
}
