package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "kontext",
		Short:         "Automatic Kubernetes Context switcher",
		SilenceErrors: true,
	}

	return cmd
}

func Execute(ctx context.Context) {
	cmd := RootCommand()
	cmd.AddCommand(HookCommand())
	cmd.AddCommand(InitCommand())
	cmd.AddCommand(VersionCommand())

	cmd.SetOut(os.Stdout)

	if err := cmd.ExecuteContext(ctx); err != nil {
		// fmt.Printf("🎃 %s\n", color.Red(err))
		os.Exit(1)
	}
}
