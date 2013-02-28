package filter

import (
	"code.google.com/p/go.net/html"
)

type Filters []Filter

func (fs Filters) Any(n *html.Node) bool {
	for _, f := range fs {
		if f(n) {
			return true
		}
	}
	return false
}

func (fs Filters) All(n *html.Node) bool {
	for _, f := range fs {
		if !f(n) {
			return false
		}
	}
	return true
}
