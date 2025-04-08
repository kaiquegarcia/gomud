package maps

import "encoding/json"

func Get[T any](m map[string]string, key string, def T) T {
	value, ok := m[key]
	if !ok {
		return def
	}

	if _, ok := any(def).(string); ok {
		return any(value).(T)
	}

	var output T
	err := json.Unmarshal([]byte(value), &output)
	if err != nil {
		return def
	}

	return output
}
