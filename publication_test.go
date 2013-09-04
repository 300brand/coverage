package coverage

import (
	"net/url"
	"testing"
)

func TestNewFeed(t *testing.T) {
	p := NewPublication()
	t.Logf("%+v", p)
}
