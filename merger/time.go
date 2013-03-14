package merger

import (
	"sort"
	"time"
)

type Times []time.Time

var _ sort.Interface = Times{}

func (t Times) Last() time.Time    { return t[len(t)-1] }
func (t Times) Len() int           { return len(t) }
func (t Times) Less(i, j int) bool { return t[i].Before(t[j]) }
func (t Times) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func (t *Times) Add(u time.Time) {
	*t = append(*t, u)
	sort.Sort(t)
}
