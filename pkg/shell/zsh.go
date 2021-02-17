package shell

import (
	"bytes"
	"fmt"
	"io"

	"github.com/stackvista/kontext/pkg/env"
	"github.com/stackvista/kontext/pkg/hook"
)

// ZSH is a singleton instance of ZSH_T
type zsh struct{}

// Zsh adds support for the venerable Z shell.
var Zsh Shell = zsh{}

const zshHook = `
_kontext_hook() {
  trap -- '' SIGINT;
  eval "$("{{.SelfPath}}" export zsh)";
  trap - SIGINT;
}
typeset -ag precmd_functions;
if [[ -z ${precmd_functions[(r)_kontext_hook]} ]]; then
  precmd_functions=( _kontext_hook ${precmd_functions[@]} )
fi
typeset -ag chpwd_functions;
if [[ -z ${chpwd_functions[(r)_kontext_hook]} ]]; then
  chpwd_functions=( _kontext_hook ${chpwd_functions[@]} )
fi
`

func (sh zsh) Hook() (hook.Hook, error) {
	return hook.Hook{Contents: zshHook}, nil
}

func (sh zsh) Export(e Export) (string, error) {
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

func (sh zsh) Dump(env env.Environment) (string, error) {
	buf := bytes.NewBuffer(nil)
	for key, value := range env {
		err := sh.export(buf, key, value)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func (sh zsh) export(buf io.StringWriter, key, value string) error {
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

func (sh zsh) unset(buf io.StringWriter, key string) error {
	k, err := sh.escape(key)
	if err != nil {
		return err
	}

	_, err = buf.WriteString(fmt.Sprintf("unset %s;", k))
	return err
}

func (sh zsh) escape(str string) (string, error) {
	return BashEscape(str)
}
