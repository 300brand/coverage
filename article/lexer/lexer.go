package lexer

import (
	"strings"
)

var StemmingEnabled = false

func Words(b []byte) []string {
	return strings.Fields(string(Normalize(b)))
}
