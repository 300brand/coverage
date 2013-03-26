package trie

import (
	"testing"
)

var (
	a = []rune("testing")
	b = []rune("tests")
	c = []rune("test")
)

func TestLongShort(t *testing.T) {
	o := overlap(a, b)
	if s := string(o); s != "test" {
		t.Errorf("Expected test, got: %s", s)
	}
}

func TestShortLong(t *testing.T) {
	o := overlap(b, a)
	if s := string(o); s != "test" {
		t.Errorf("Expected test, got: %s", s)
	}
}

func TestASubB(t *testing.T) {
	o := overlap(c, a)
	if s := string(o); s != "test" {
		t.Errorf("Expected test, got: %s", s)
	}
}

func TestBSubA(t *testing.T) {
	o := overlap(a, c)
	if s := string(o); s != "test" {
		t.Errorf("Expected test, got: %s", s)
	}
}
