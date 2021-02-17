package shell

import (
	"encoding/json"
	"errors"

	"github.com/stackvista/kontext/pkg/env"
	"github.com/stackvista/kontext/pkg/hook"
)

// jsonShell is not a real shell
type jsonShell struct{}

// JSON is not really a shell but it fits. Useful to add support to editor and
// other external tools that understand JSON as a format.
var JSON Shell = jsonShell{}

func (sh jsonShell) Hook() (hook.Hook, error) {
	return hook.Hook{}, errors.New("this feature is not supported")
}

func (sh jsonShell) Export(e Export) (string, error) {
	out, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (sh jsonShell) Dump(env env.Environment) (string, error) {
	out, err := json.MarshalIndent(env, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}
