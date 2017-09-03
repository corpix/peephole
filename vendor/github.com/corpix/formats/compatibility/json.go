package compatibility

import (
	"fmt"
)

func jsonArray(in []interface{}) []interface{} {
	res := make([]interface{}, len(in))
	for i, v := range in {
		res[i] = JSON(v)
	}
	return res
}

func jsonMap(in map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range in {
		res[fmt.Sprintf("%v", k)] = JSON(v)
	}
	return res
}

// JSON force types in the structs to be compatible with JSON marshaler.
// Mitigating https://github.com/go-yaml/yaml/issues/139
func JSON(v interface{}) interface{} {
	switch v := v.(type) {
	case []interface{}:
		return jsonArray(v)
	case map[interface{}]interface{}:
		return jsonMap(v)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}
