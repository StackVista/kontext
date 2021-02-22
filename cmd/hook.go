package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stackvista/kontext/pkg/hook"
	"github.com/stackvista/kontext/pkg/shell"
)

func HookCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "hook <shell>",
		Short:              "Returns a shell hook function to automatically switch on changing directories",
		Args:               cobra.ExactArgs(1),
		DisableFlagParsing: true,
		Example: `To register the hook in your shell, on most shells you can simply execute:

	> eval "$(kontext hook <shell>)"

The shell can be autodetected by providing a '-' for the <shell> argument.`,
		RunE: runHook,
	}

	return cmd
}

func runHook(cmd *cobra.Command, args []string) error {
	if IsHelp(cmd, args) {
		return cmd.Usage()
	}

	target := args[0]

	shell, err := shell.DetectShell(target)
	if err != nil {
		return err
	}

	hook, err := shell.Hook()
	if err != nil {
		return err
	}

	hookContext, err := buildHookContext()
	if err != nil {
		return err
	}

	rendered, err := hook.Render(hookContext)
	if err != nil {
		return err
	}

	cmd.Print(rendered)

	return nil
}

// buildHookContext builds the rendering Context for the shell hook
func buildHookContext() (hook.Context, error) {
	selfPath, err := os.Executable()
	if err != nil {
		return hook.Context{}, err
	}

	// Convert Windows path if needed
	selfPath = strings.ReplaceAll(selfPath, "\\", "/")
	ctx := hook.Context{
		SelfPath: selfPath,
		Name:     "kontext",
	}

	return ctx, nil
}
