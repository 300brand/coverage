package logger

type Level int

const (
	Ldebug Level = 1 << iota
	Lservice
	Lerror
)

func (l Level) String() string {
	switch l {
	case Ldebug:
		return "DEBUG"
	case Lservice:
		return "SERVICE"
	case Lerror:
		return "ERROR"
	}
	return "UNKNOWN"
}
