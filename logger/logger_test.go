package logger

import (
	"testing"
)

func TestDebug(t *testing.T) {
	Debug("Testing DEBUG level")
	Debugf("Const: %d", _DEBUG)
}

func TestWarn(t *testing.T) {
	Warn("Testing WARN level")
	Warnf("Const: %d", _WARN)
}
