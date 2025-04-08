package envs

import (
	"encoding/json"
	"os"
)

func Get[Type any](key string, defaultValue Type) Type {
	stringValue := os.Getenv(key)
	if stringValue == "" {
		return defaultValue
	}

	if _, ok := any(defaultValue).(string); ok {
		return any(stringValue).(Type)
	}

	var parsedValue Type
	err := json.Unmarshal([]byte(stringValue), &parsedValue)
	if err != nil {
		return defaultValue
	}

	return parsedValue
}
