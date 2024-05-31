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
