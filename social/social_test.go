package social

import (
	"testing"
)

func TestSocial(t *testing.T) {
	s := new(Stats)
	if err := FetchString("http://www.google.com", s); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", s)
}
