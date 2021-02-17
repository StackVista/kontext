package escape

import (
	"bytes"
	"fmt"
)

type Source func(b byte) bool
type Function func(buf *bytes.Buffer, b byte) (bool, error)

var _ Function = Literal
var _ Function = Quoted
var _ Function = Backslash
var _ Function = Hex

func B(expected byte) Source {
	return func(b byte) bool {
		return b == expected
	}
}

// BR compares the incoming byte as b > from && b <= to
func BR(from byte, to byte) Source {
	return func(b byte) bool {
		return b > from && b <= to
	}
}

func Literal(buf *bytes.Buffer, b byte) (bool, error) {
	return false, buf.WriteByte(b)
}

func Quoted(buf *bytes.Buffer, b byte) (bool, error) {
	return true, buf.WriteByte(b)
}

func Backslash(buf *bytes.Buffer, b byte) (bool, error) {
	_, err := buf.Write([]byte{'\\', b})
	return true, err
}

func Hex(buf *bytes.Buffer, b byte) (bool, error) {
	_, err := buf.WriteString(fmt.Sprintf("\\x%02x", b))
	return true, err
}

func Escaped(s string) Function {
	return func(buf *bytes.Buffer, b byte) (bool, error) {
		_, err := buf.WriteString(s)
		return true, err
	}
}
