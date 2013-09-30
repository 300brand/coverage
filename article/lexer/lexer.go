package lexer

import (
	"github.com/jbaikge/logger"
	"strings"
)

var StemmingEnabled = false

func Words(b []byte) []string {
	logger.Trace.Print("Words: called")
	return strings.Fields(string(Normalize(b)))
}
