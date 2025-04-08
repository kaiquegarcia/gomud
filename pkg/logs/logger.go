package logs

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	Logger interface {
		With(fields ...FieldFunc) Logger
		Debug(message string, fields ...FieldFunc)
		Info(message string, fields ...FieldFunc)
		Warn(message string, fields ...FieldFunc)
		Error(message string, fields ...FieldFunc)
	}

	logger struct {
		minLevelToReport Level
		embed            []FieldFunc
	}
)

func NewLogger(minLevelToReport Level) Logger {
	return &logger{
		minLevelToReport: minLevelToReport,
		embed:            make([]FieldFunc, 0),
	}
}

func (l *logger) With(fields ...FieldFunc) Logger {
	return &logger{
		minLevelToReport: l.minLevelToReport,
		embed:            append(l.embed, fields...),
	}
}

func (l *logger) Debug(message string, fields ...FieldFunc) {
	l.log(LevelDebug, message, fields)
}

func (l *logger) Info(message string, fields ...FieldFunc) {
	l.log(LevelInfo, message, fields)
}

func (l *logger) Warn(message string, fields ...FieldFunc) {
	l.log(LevelWarning, message, fields)
}

func (l *logger) Error(message string, fields ...FieldFunc) {
	l.log(LevelError, message, fields)
}

func (l *logger) log(level Level, message string, fields []FieldFunc) {
	if level < l.minLevelToReport {
		return
	}

	adt := ""
	fieldFuncs := append(l.embed, fields...)
	if len(fieldFuncs) > 0 {
		fieldMap := make(FieldMap)
		for _, f := range fieldFuncs {
			f(fieldMap)
		}

		bytes, err := json.Marshal(fieldMap)
		if err != nil {
			fmt.Printf("ERROR ENCODING LOG FIELDs: %s\n", err)
		} else {
			adt = fmt.Sprintf("\n%s", string(bytes))
		}
	}

	fmt.Printf("%s [%s]: %s%s\n", time.Now().Format(time.RFC3339), level.Label(), message, adt)
}
