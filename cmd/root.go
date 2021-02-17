package cmd

import (
	"context"
	"fmt"
	"os"

	color "github.com/logrusorgru/aurora/v3"

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
	cmd.AddCommand(VersionCommand())

	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Printf("🎃 %s\n", color.Red(err))
		os.Exit(1)
	}
}
