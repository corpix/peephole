package nsq

const (
	DebugLevel Level = iota
	InfoLevel
	ErrorLevel
	FatalLevel
)

// Level is a logger level.
type Level uint8

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBU"
	case InfoLevel:
		return "INFO"
	case ErrorLevel:
		return "ERRO"
	case FatalLevel:
		return "FATA"
	default:
		return "INFO"
	}
}
