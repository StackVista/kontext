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

	// Export outputs the Export as an evaluatable string on the host shell
	Export(e env.Export) (string, error)

	// Dump outputs and evaluatable string that sets the env in the host shell
	Dump(env env.Environment) (string, error)
}

// DetectShell returns a Shell instance from the given target.
//
// target is usually $0 and can also be prefixed by `-`
func DetectShell(target string) (Shell, error) {
	shell := target

	// $0 starts with "-"
	if shell[0:1] == "-" {
		shell = shell[1:]
	}

	if shell == "" {
		s, err := env.FindEnvironment("SHELL")
		if err != nil {
			return nil, ErrUnknownShell{target}
		}

		shell = s
	}

	shell = filepath.Base(shell)

	switch shell {
	case "bash":
		return Bash, nil
	case "zsh":
		return Zsh, nil
	case "json":
		return JSON, nil
	}

	return nil, ErrUnknownShell{target}
}
