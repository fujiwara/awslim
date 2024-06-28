package sdkclient

import (
	"encoding/json"
	"strings"
)

func marshalJSON(v interface{}, camelCase bool) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	if !camelCase {
		return json.Marshal(v)
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	switch b[0] {
	case '{':
		m := map[string]any{}
		if err := json.Unmarshal(b, &m); err != nil {
			return nil, err
		}
		walkMap(m, toCamelCase)
		return json.Marshal(m)
	case '[':
		a := []any{}
		if err := json.Unmarshal(b, &a); err != nil {
			return nil, err
		}
		walkArray(a, toCamelCase)
		return json.Marshal(a)
	default:
		return b, nil
	}
}

func toCamelCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func walkMap(m map[string]interface{}, fn func(string) string) {
	for key, value := range m {
		delete(m, key)
		newKey := key
		if fn != nil {
			newKey = fn(key)
		}
		if value != nil {
			m[newKey] = value
		}
		switch value := value.(type) {
		case map[string]any:
			walkMap(value, fn)
		case []interface{}:
			walkArray(value, fn)
		default:
		}
	}
}

func walkArray(a []interface{}, fn func(string) string) {
	for _, value := range a {
		switch value := value.(type) {
		case map[string]interface{}:
			walkMap(value, fn)
		case []interface{}:
			walkArray(value, fn)
		default:
		}
	}
}
