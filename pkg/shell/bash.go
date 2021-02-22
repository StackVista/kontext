package shell

import (
	"bytes"
	"fmt"
	"io"

	"github.com/stackvista/kontext/pkg/env"
	"github.com/stackvista/kontext/pkg/hook"
	"github.com/stackvista/kontext/pkg/shell/escape"
)

type bash struct{}

// Bash shell instance
var Bash Shell = bash{}

const bashHook = `
_{{.Name}}_hook() {
  local previous_exit_status=$?;
  trap -- '' SIGINT;
  eval "$("{{.SelfPath}}" init bash)";
  trap - SIGINT;
  return $previous_exit_status;
};
if ! [[ "${PROMPT_COMMAND:-}" =~ _{{.Name}}_hook ]]; then
  PROMPT_COMMAND="_{{.Name}}_hook${PROMPT_COMMAND:+;$PROMPT_COMMAND}"
fi
`

func (sh bash) Hook() (hook.Hook, error) {
	return hook.Hook{Contents: bashHook}, nil
}

func (sh bash) Export(e env.Export) (string, error) {
	buf := bytes.NewBuffer(nil)
	for key, value := range e {
		if value == nil {
			if err := sh.unset(buf, key); err != nil {
				return "", err
			}
		} else {
			if err := sh.export(buf, key, *value); err != nil {
				return "", err
			}
		}
	}
	return buf.String(), nil
}

func (sh bash) Dump(env env.Environment) (string, error) {
	buf := bytes.NewBuffer(nil)
	for key, value := range env {
		err := sh.export(buf, key, value)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func (sh bash) export(buf io.StringWriter, key, value string) error {
	k, err := sh.escape(key)
	if err != nil {
		return err
	}

	v, err := sh.escape(value)
	if err != nil {
		return err
	}

	_, err = buf.WriteString(fmt.Sprintf("export %s=%s;", k, v))
	return err
}

func (sh bash) unset(buf io.StringWriter, key string) error {
	k, err := sh.escape(key)
	if err != nil {
		return err
	}

	_, err = buf.WriteString(fmt.Sprintf("unset %s;", k))
	return err
}

func (sh bash) escape(str string) (string, error) {
	return BashEscape(str)
}

// https://github.com/solidsnack/shell-escape/blob/master/Text/ShellEscape/Bash.hs
/*
A Bash escaped string. The strings are wrapped in @$\'...\'@ if any
bytes within them must be escaped; otherwise, they are left as is.
Newlines and other control characters are represented as ANSI escape
sequences. High bytes are represented as hex codes. Thus Bash escaped
strings will always fit on one line and never contain non-ASCII bytes.
*/
func BashEscape(str string) (string, error) {
	s, esc, err := escape.Escape(BashMapping, str)
	if err != nil {
		return "", err
	}

	if str == "" {
		return "''", nil
	}

	if esc {
		return fmt.Sprintf("$'%s'", s), nil
	}

	return s, nil
}

func BashMapping(char byte) escape.Function {
	switch {
	case char == escape.ACK:
		return escape.Hex
	case char == escape.TAB:
		return escape.Escaped(`\t`)
	case char == escape.LF:
		return escape.Escaped(`\n`)
	case char == escape.CR:
		return escape.Escaped(`\r`)
	case char <= escape.US:
		return escape.Hex
	case char <= escape.AMPERSTAND:
		return escape.Quoted
	case char == escape.SINGLE_QUOTE:
		return escape.Backslash
	case char <= escape.PLUS:
		return escape.Quoted
	case char <= escape.NINE:
		return escape.Literal
	case char <= escape.QUESTION:
		return escape.Quoted
	case char <= escape.LOWERCASE_Z:
		return escape.Literal
	case char == escape.OPEN_BRACKET:
		return escape.Quoted
	case char == escape.BACKSLASH:
		return escape.Backslash
	case char == escape.UNDERSCORE:
		return escape.Literal
	case char <= escape.TILDA:
		return escape.Quoted
	case char == escape.DEL:
		return escape.Hex
	default:
		return escape.Hex
	}
}
