package task

import (
	"bytes"
	"encoding/json"
	"strconv"
)

func unquoteJSONOutput(s *string) {
	unescaped, err := strconv.Unquote(*s)
	if err == nil {
		*s = unescaped
	}
}

func Prettify(s string) string {
	buf := new(bytes.Buffer)
	json.Indent(buf, []byte(s), "", "  ")
	return buf.String()
}
