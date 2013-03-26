// Derived from information gathered here:
// http://www.geekm.ag/Archive/Quickly_identifying_phrases_in_text_using_Go_(golang)
//
// And sample code here: github.com/rjohnsondev/go-trie
//
// Rewritten to make a little more idiomatic and easier to understand
package trie

const (
	CharLow  = 'a'
	CharHigh = 'z'
	CharLen  = CharHigh - CharLow + 1
)

type Trie struct {
	Trunk *Branch
}

func New() *Trie {
	t := &Trie{}
	t.Trunk = NewBranch()
	return t
}

func (t *Trie) Add(entry string) {
	t.Trunk.Add([]rune(entry))
}

func (t *Trie) Dump() (out string) {
	return t.Trunk.Dump(0)
}

func (t *Trie) Has(entry string) (found bool) {
	b := t.Trunk
	runes := []rune(entry)

	// it's <= here to ensure we get to the cheat comparison nil on a valid path
	for i := 0; i <= len(runes); i++ {
		// Match Leaf, if it exists
		for j := 0; j < len(b.Leaf); j++ {
			if i >= len(runes) {
				return false
			}
			if b.Leaf[j] != runes[i] {
				return false
			}
			i++
		}
		// We have reassembled the word, completely
		if i == len(runes) {
			return b.End
		}
		idx := b.Index(runes[i])
		// No branch
		if len(b.Branches) < idx || b.Branches[idx] == nil {
			return false
		}
		b = b.Branches[idx]
	}
	return true
}
