package shell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBashEscape(t *testing.T) {
	tests := map[string]string{
		"":             `''`,
		"escape'quote": `$'escape\'quote'`,
		"foo\r\n\tbar": `$'foo\r\n\tbar'`,
		"foo bar":      `$'foo bar'`,
		"é":            `$'\xc3\xa9'`,
	}
	for k, v := range tests {
		key, val := k, v
		t.Run("", func(t *testing.T) {
			s, err := BashEscape(key)
			assert.NoError(t, err)
			assert.Equal(t, val, s)
		})
	}
}
