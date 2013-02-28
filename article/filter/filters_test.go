package filter

import (
	"code.google.com/p/go.net/html"
	"testing"
)

var node = &html.Node{}

func TestAllFalse(t *testing.T) {
	filters := []Filters{
		Filters{f1},
		Filters{f1, t1},
		Filters{t1, f1},
	}
	for i, f := range filters {
		if f.All(node) {
			t.Errorf("[%d] Expected false", i)
		}
	}
}

func TestAllTrue(t *testing.T) {
	filters := []Filters{
		Filters{t1},
		Filters{t1, t2},
	}
	for i, f := range filters {
		if !f.All(node) {
			t.Errorf("[%d] Expected true", i)
		}
	}
}

func TestAnyFalse(t *testing.T) {
	filters := []Filters{
		Filters{f1},
		Filters{f1, f2},
	}
	for i, f := range filters {
		if f.Any(node) {
			t.Errorf("[%d] Expected false", i)
		}
	}
}

func TestAnyTrue(t *testing.T) {
	filters := []Filters{
		Filters{t1},
		Filters{t1, f1},
		Filters{f1, t1},
	}
	for i, f := range filters {
		if !f.Any(node) {
			t.Errorf("[%d] Expected true", i)
		}
	}
}

func f1(n *html.Node) bool { return false }
func f2(n *html.Node) bool { return false }
func t1(n *html.Node) bool { return true }
func t2(n *html.Node) bool { return true }
