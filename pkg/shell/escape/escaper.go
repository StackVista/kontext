package escape

import (
	"bytes"
)

type Mapping func(b byte) Function

// Escape takes a mapping and the input string, and escapes the characters of the string where needed.
// It returns the escaped string and a boolean indicated whether the string contained any characters that needed escaping
// Or it returns an error if an error occurred during the escaping
func Escape(mapping Mapping, s string) (string, bool, error) {
	if s == "" {
		return "", false, nil
	}
	in := []byte(s)
	out := bytes.NewBuffer(nil)
	esc := false

	for _, char := range in {
		f := mapping(char)

		b, err := f(out, char)
		if err != nil {
			return "", esc, err
		}
		esc = b || esc
	}

	return out.String(), esc, nil
}
