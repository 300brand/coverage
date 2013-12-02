package filter

import (
	"code.google.com/p/go.net/html"
	"unicode"
)

var DoubleQuote = &unicode.RangeTable{
	R16: []unicode.Range16{
		//{0x0022, 0x0022, 1},
		{0x201c, 0x201f, 1},
		{0xff02, 0xff02, 1},
	},
}

var SingleQuote = &unicode.RangeTable{
	R16: []unicode.Range16{
		//{0x0027, 0x0027, 1},
		{0x2018, 0x201b, 1},
		{0xff07, 0xff07, 1},
	},
}

var table = map[rune][]rune{
	0x0085: []rune("..."), // &hellip; Horizontal Ellipsis
	0x2026: []rune("..."), // Horizontal Ellipsis
	0x0095: []rune{'-'},   // &bull; Bullet
}

func TranslateUnicode(n *html.Node) bool {
	if n.Type == html.TextNode {
		n.Data = TranslateString(n.Data)
	}
	return false
}

func TranslateString(s string) string {
	r := make([]rune, 0, len(s))
	for _, c := range s {
		switch {
		case unicode.Is(DoubleQuote, c):
			r = append(r, '"')
		case unicode.Is(SingleQuote, c):
			r = append(r, '\'')
		case unicode.Is(unicode.Dash, c):
			r = append(r, '-')
		case c > 127:
			r = append(r, table[c]...)
		default:
			r = append(r, c)
		}
	}
	return string(r)
}
