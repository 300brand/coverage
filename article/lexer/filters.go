package lexer

type filterFunc func(rune) bool

var filterFuncs []filterFunc

func init() {
	filterFuncs = []filterFunc{
		space,
		quote,
		punctuation,
	}
}

func punctuation(r rune) bool {
	switch r {
	case ',':
	case '.':
	case '!':
	case '?':
	case '&':
	default:
		return false
	}
	return true
}

func quote(r rune) bool {
	switch r {
	case '"':
	default:
		return false
	}
	return true
}

func space(r rune) bool {
	switch r {
	case ' ':
	case '\n':
	case '\r':
	case '\t':
	default:
		return false
	}
	return true
}
