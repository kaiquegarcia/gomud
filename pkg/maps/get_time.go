package maps

import (
	"time"
)

func GetTime(m map[string]string, key string, def time.Time, layout string) time.Time {
	valueStr, ok := m[key]
	if !ok {
		return def
	}

	value, err := time.Parse(layout, valueStr)
	if err != nil {
		return def
	}

	return value
}
