package strtotime

import (
	"testing"
	"time"
)

var now = time.Now()

func TestFormats(t *testing.T) {
	for i, layout := range formats {
		s := now.Format(layout)
		if _, err := Parse(s); err != nil {
			t.Errorf("[%02d] (%s): %s - %s", i, layout, s, err)
		}
	}
}
