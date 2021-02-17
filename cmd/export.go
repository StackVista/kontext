package cmd

import "github.com/spf13/cobra"

func ExportCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "export <shell>",
		Args:               cobra.ExactArgs(1),
		DisableFlagParsing: true,
		RunE:               runExport,
	}
}

func runExport(cmd *cobra.Command, args []string) error {
	return nil
}
