package coverage

import (
	"testing"
)

func TestNewFeed(t *testing.T) {
	p := NewPublication()
	t.Logf("%+v", p)
}
