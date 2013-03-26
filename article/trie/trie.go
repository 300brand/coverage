// Derived from information gathered here:
// http://www.geekm.ag/Archive/Quickly_identifying_phrases_in_text_using_Go_(golang)
//
// And sample code here: github.com/rjohnsondev/go-trie
//
// Rewritten to make a little more idiomatic and easier to understand
package trie

import (
	"fmt"
)

const (
	CharLow  = 'a'
	CharHigh = 'z'
	CharLen  = CharHigh - CharLow + 1
)

type Trie struct {
	Len   int
	Trunk *Branch
}

func New() *Trie {
	t := &Trie{}
	t.Trunk = NewBranch()
	return t
}

func (t *Trie) Add(entry string) {
	t.Len++
	t.Trunk.Add([]rune(entry))
}

func (t *Trie) Dump() (out string) {
	out = fmt.Sprintf("Words: %d\n", t.Len)
	out += t.Trunk.Dump(0)
	return
}

func (t *Trie) Has(entry string) (found bool) {
	/*
		b := t.Trunk

		// it's <= here to ensure we get to the cheat comparison nil on a valid path
		for i := 0; i <= len(entry); i++ {
			// Match Leaf, if it exists
		}
	*/
	return
}
