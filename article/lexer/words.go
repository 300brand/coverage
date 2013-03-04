package lexer

type Words []Word

func (ws *Words) Add(w Word) {
	*ws = append(*ws, w)
}

func (ws *Words) Strings() (s []string) {
	for _, w := range *ws {
		s = append(s, w.Word)
	}
	return
}
