package adp

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// unquoteJSONOutput unquotes a JSON output string.
func UnquoteJSONOutput(s *string) {
	unescaped, err := strconv.Unquote(*s)
	if err == nil {
		*s = unescaped
	}
}

// Prettify is a function that takes a string as input and returns a prettified version of the string in JSON format.
func Prettify(s string) string {
	buf := new(bytes.Buffer)
	json.Indent(buf, []byte(s), "", "  ")
	return buf.String()
}

// StructToString converts a struct to a string representation using JSON encoding.
func StructToString(i interface{}) (string, error) {
	b, err := json.Marshal(i)
	return string(b), err
}

func NormalizeEntityName(name string) string {
	var result string

	for _, c := range name {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			result += string(c)
		} else {
			if result == "" || result[len(result)-1] != '_' {
				result += "_"
			}
		}
	}
	return result
}
