package trie

import (
	"fmt"
	"strings"
)

const dumpPad = "-"

type Branch struct {
	Branches []*Branch
	Runes    map[rune]int `json:"-"`
	Leaf     []rune       // AKA "Shortcut"
	End      bool
}

func NewBranch() *Branch {
	return &Branch{
		Runes:    map[rune]int{},
		Branches: make([]*Branch, 0, CharLen+5),
	}
}

func (b *Branch) Accommodate(index int) (extended bool) {
	branchLen := len(b.Branches)
	idxLen := index + 1
	if idxLen <= branchLen {
		return
	}
	b.Branches = append(b.Branches, make([]*Branch, idxLen-branchLen)...)
	return true
}

func (b *Branch) Add(entry []rune) {
	// No end previously exists, this branch ends here
	if b.Leaf == nil {
		b.Leaf = entry
		b.End = true
		return
	}

	// At the end of the branch, or added a duplicate word
	if len(b.Leaf) == 0 && len(entry) == 0 {
		return
	}

	// Determine overlap between branch's leaf and incoming entry
	leaf := overlap(b.Leaf, entry)

	// New leaf is smaller, send the rest of the current leaf down to become
	// new branches
	if len(leaf) < len(b.Leaf) {
		tail := b.Leaf[len(leaf):]
		idx := b.MakeIndex(tail[0])
		branch := NewBranch()
		branch.Leaf, b.Leaf = tail[1:], leaf
		// Memory already allocated, just swap
		branch.Branches, b.Branches = b.Branches, branch.Branches
		branch.End, b.End = b.End, branch.End
		branch.Runes = b.Runes
		// Push new branch in
		b.Accommodate(idx)
		b.Branches[idx] = branch
	}

	// New leaf is smaller than the entry, send tail to become new branch
	if len(leaf) < len(entry) {
		tail := entry[len(leaf):]
		idx := b.MakeIndex(tail[0])
		if b.Accommodate(idx) || b.Branches[idx] == nil {
			b.Branches[idx] = NewBranch()
		}
		b.Branches[idx].Add(tail[1:])
	} else {
		b.End = true
	}
}

func (b *Branch) Dump(depth int) (out string) {
	leafLen := len(b.Leaf)
	if leafLen > 0 {
		out += fmt.Sprintf("%s L:%s\n", strings.Repeat(dumpPad, depth), string(b.Leaf))
	}
	if b.End {
		out += fmt.Sprintf("%s $\n", strings.Repeat(dumpPad, depth+leafLen))
	}
	for r, br := range b.Branches {
		if br != nil {
			out += fmt.Sprintf("%s I:%s\n", strings.Repeat(dumpPad, depth+leafLen), string([]rune{rune(r) + 'a'}))
			out += br.Dump(depth + leafLen + 1)
		}
	}
	return
}

func (b *Branch) Index(r rune) (i int) {
	if CharLow <= r && r <= CharHigh {
		return int(r - CharLow)
	}
	i, ok := b.Runes[r]
	if !ok {
		return -1
	}
	return
}

func (b *Branch) MakeIndex(r rune) (i int) {
	i = b.Index(r)
	if i == -1 {
		i = len(b.Runes) + CharLen
		b.Runes[r] = i
	}
	return
}
