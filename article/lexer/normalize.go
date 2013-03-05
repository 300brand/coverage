package lexer

// Pre-processes text to swap out certain characters for spaces or nothing
func Normalize(in []byte) (out []byte) {
	out = make([]byte, 0, len(in))
	for _, b := range in {
		switch {
		case punctuation(b) || quote(b):
			// skip
		case space(b) || hyphenation(b):
			out = append(out, ' ')
		default:
			out = append(out, b)
		}
	}
	return
}

func hyphenation(r byte) bool {
	switch r {
	case '-':
	default:
		return false
	}
	return true
}

func punctuation(r byte) bool {
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

func quote(r byte) bool {
	switch r {
	case '"':
	default:
		return false
	}
	return true
}

func space(r byte) bool {
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
