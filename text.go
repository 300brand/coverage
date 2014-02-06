package coverage

type Text struct {
	Words Words
	HTML  []byte   // HTML for article (Page 1)
	Pages [][]byte // HTML for additional pages (i:0 == Page 2)
	Body  Body
}
