package coverage

import (
	"sort"
	"strings"
)

type Words struct {
	All      []string
	Keywords []string
	Stems    []string
}

func (ws *Words) Add(w Word) {
	ws.All = append(ws.All, w.Word)
	ws.Keywords = uAppend(ws.Keywords, strings.ToLower(w.Word))
	if w.Stem != "" {
		ws.Stems = uAppend(ws.Stems, w.Stem)
	}
}

func (ws *Words) Strings() (s []string) {
	return ws.All
}

func uAppend(a []string, s string) []string {
	for _, as := range a {
		if as == s {
			return a
		}
	}
	a = append(a, s)
	sort.Strings(a)
	return a
}
