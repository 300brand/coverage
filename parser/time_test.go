package parser

import (
	"testing"
	"time"
)

var now = time.Now()

func TestFormats(t *testing.T) {
	for i, layout := range formats {
		ts := Time(now.Format(layout))
		if _, err := ts.Parse(); err != nil {
			t.Errorf("[%02d] (%s): %s - %s", i, layout, ts, err)
		}
	}
}
