package lexer

import (
	"git.300brand.com/coverage"
	"github.com/rookii/paicehusk"
	"strings"
)

var StemmingEnabled = false

func GetWords(b []byte) []string {
	return strings.Fields(string(Normalize(b)))
}
