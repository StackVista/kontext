package shell

import (
	"fmt"
	"path/filepath"

	"github.com/stackvista/kontext/pkg/env"
	"github.com/stackvista/kontext/pkg/hook"
)

type ErrUnknownShell struct {
	target string
}

func (e ErrUnknownShell) Error() string {
	return fmt.Sprintf("Unknown shell: %s", e.target)
}

// Shell is the interface that represents the interaction with the host shell.
type Shell interface {
	// Hook is the string that gets evaluated into the host shell config and
	// setups direnv as a prompt hook.
	Hook() (hook.Hook, error)

	// Export outputs the ShellExport as an evaluatable string on the host shell
	Export(e Export) (string, error)

	// Dump outputs and evaluatable string that sets the env in the host shell
	Dump(env env.Environment) (string, error)
}

// Export represents environment variables to add and remove on the host
// shell.
type Export map[string]*string

// Add represents the additon of a new environment variable
func (e Export) Add(key, value string) {
	e[key] = &value
}

// Remove represents the removal of a given `key` environment variable.
func (e Export) Remove(key string) {
	e[key] = nil
}

// DetectShell returns a Shell instance from the given target.
//
// target is usually $0 and can also be prefixed by `-`
func DetectShell(target string) (Shell, error) {
	// $0 starts with "-"
	if target[0:1] == "-" {
		target = target[1:]
	}

	if target == "" {
		// No shell entered, try to detect through environment
		e := env.GetEnvironment()
		if s, ok := e["SHELL"]; ok {
			target = s
		}
	}

	target = filepath.Base(target)

	switch target {
	case "bash":
		return Bash, nil
	case "zsh":
		return Zsh, nil
	case "json":
		return JSON, nil
	}

	return nil, ErrUnknownShell{target}
}
