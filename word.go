package coverage

type Word struct {
	Word  string
	Stem  string
}

func (w Word) String() string {
	return string(w.Word)
}
