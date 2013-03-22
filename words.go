package coverage

type Words struct {
	All   []string
	Lower []string
	Stems []string
}

func (ws *Words) Add(w Word) {
	ws.All = append(ws.All, w.Word)
}

func (ws *Words) Strings() (s []string) {
	return ws.All
}
