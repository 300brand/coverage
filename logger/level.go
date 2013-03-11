package logger

type Level int

const (
	Ldebug Level = iota << 1
	Lerror
)

func (l Level) String() string {
	switch l {
	case Ldebug:
		return "DEBUG"
	case Lerror:
		return "ERROR"
	}
	return "UNKNOWN"
}
