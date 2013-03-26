package trie

import (
	"math"
)

func overlap(a, b []rune) (o []rune) {
	minLen := int(math.Min(float64(len(a)), float64(len(b))))
	o = make([]rune, 0, minLen)
	for i := 0; i < minLen; i++ {
		if a[i] == b[i] {
			o = append(o, a[i])
		}
	}
	return
}
