package merger

import (
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	stamps := []int64{1000, 5000, 2000, 4000}
	ordered := []int64{1000, 2000, 4000, 5000}
	times := Times{}
	for _, stamp := range stamps {
		times.Add(time.Unix(stamp, 0))
	}
	for i, stamp := range times {
		if ordered[i] != stamp.Unix() {
			t.Errorf("[%d] %d", i, stamp.Unix())
		}
	}
}

func TestLast(t *testing.T) {
	stamps := []int64{1000, 5000, 2000, 4000}
	times := Times{}
	for _, stamp := range stamps {
		times.Add(time.Unix(stamp, 0))
	}
	if times.Last().Unix() != stamps[1] {
		t.Error("Wrong time at the end")
		t.Errorf("%v", times)
	}
}
