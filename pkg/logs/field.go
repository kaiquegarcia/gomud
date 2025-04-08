package logs

import "github.com/pingcap/errors"

type (
	FieldMap  map[string]any
	FieldFunc func(m FieldMap)
)

func Field(key string, val any) FieldFunc {
	return func(m FieldMap) {
		m[key] = val
	}
}

func Error(err error) FieldFunc {
	return func(m FieldMap) {
		if err == nil {
			return
		}

		m["error"] = err.Error()
		m["stack_trace"] = errors.ErrorStack(err)
	}
}
