package lexer

import (
	"bytes"
	"github.com/300brand/coverage/article/trie"
	"sort"
	"strings"
)

var stopWords *trie.Trie

func Keywords(b []byte) (kw []string) {
	words := Words(bytes.ToLower(b))
	sort.Strings(words)
	lastword := ""
	for _, w := range words {
		if w == lastword {
			continue
		}
		if !stopWords.Has(w) {
			kw = append(kw, w)
		}
		lastword = w
	}
	return
}

func init() {
	stopWords = trie.New()
	// No punctuation as Normalize removes the single quotes
	var words = strings.Fields(`
		a about above after again against all am an and any are arent as at be
		because been before being below between both but by cant cannot could
		couldnt did didnt do does doesnt doing dont down during each few for
		from further had hadnt has hasnt have havent having he hed hell hes her
		here heres hers herself him himself his how hows i id ill im ive if in
		into is isnt it its its itself lets me more most mustnt my myself no nor
		not of off on once only or other ought our ours ourselves out over own
		same shant she shed shell shes should shouldnt so some such than that
		thats the their theirs them themselves then there theres these they
		theyd theyll theyre theyve this those through to too under until up very
		was wasnt we wed well were weve were werent what whats when whens where
		wheres which while who whos whom why whys with wont would wouldnt you
		youd youll youre youve your yours yourself yourselves
	`)
	for _, w := range words {
		stopWords.Add(w)
	}
}
