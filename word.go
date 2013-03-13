package coverage

type Word struct {
	Word  string
	Stem  string
	Index int
}

func (w Word) String() string {
	return w.Word
}
