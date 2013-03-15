package merger

import (
	"testing"
	"time"
)

func TestFieldsAdded(t *testing.T) {
	cl := Changelog{}
	cl.Changed("A", "C")
	cl.Changed("B", "C")

	expect := []string{"A", "B", "C"}

	if l := len(cl.Fields()); l != len(expect) {
		t.Errorf("Incorrect field count: %d", l)
	}

	for _, exp := range expect {
		found := false
		for _, f := range cl.Fields() {
			if exp == f {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Field %s not found", exp)
		}
	}
}

func TestChangeFound(t *testing.T) {
	cl := Changelog{}
	cl.Changed("A")

	if _, found := cl.GetChange("A"); !found {
		t.Error("Couldn't find field A")
	}
}

func TestChangeNotFound(t *testing.T) {
	cl := Changelog{}
	cl.Changed("A")

	if _, found := cl.GetChange("B"); found {
		t.Error("Found field B")
	}
}

func TestLastChange(t *testing.T) {
	cl := Changelog{}
	stamps := []int64{2000, 1000, 4000, 3000}
	for _, stamp := range stamps {
		cl.Add("A", time.Unix(stamp, 0))
	}
	if stamp := cl.LastChangeTo("A").Unix(); stamp != stamps[2] {
		t.Errorf("Incorrect timestamp: %d", stamp)
	}
}

func TestLastChangeNotFound(t *testing.T) {
	cl := Changelog{}
	stamps := []int64{2000, 1000, 4000, 3000}
	for _, stamp := range stamps {
		cl.Add("A", time.Unix(stamp, 0))
	}
	if stamp := cl.LastChangeTo("B").Unix(); stamp != (time.Time{}).Unix() {
		t.Errorf("Incorrect timestamp: %d", stamp)
	}
}
