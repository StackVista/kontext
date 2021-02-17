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
		Example: `On most shells you can simply execute:

	> eval "$(kontext hook <shell>)"
		`,
		RunE: runHook,
	}

	return cmd
}

func runHook(cmd *cobra.Command, args []string) error {
	target := args[0]
	// Because flag parsing is disabled for this command, manually check for help flag
	if target == "-h" || target == "--help" {
		return cmd.Usage()
	}

	selfPath, err := os.Executable()
	if err != nil {
		return err
	}

	// Convert Windows path if needed
	selfPath = strings.ReplaceAll(selfPath, "\\", "/")
	ctx := hook.Context{
		SelfPath: selfPath,
		Name:     "kontext",
	}

	shell, err := shell.DetectShell(target)
	if err != nil {
		return err
	}

	hook, err := shell.Hook()
	if err != nil {
		return err
	}

	rendered, err := hook.Render(ctx)
	if err != nil {
		return err
	}

	print(rendered)

	return nil
}
