package logs

type Level uint

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarning
	LevelError
)

func (lv Level) Label() string {
	switch lv {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	}

	return "undefined"
}
