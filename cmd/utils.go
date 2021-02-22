package cmd

import (
	"github.com/spf13/cobra"
)

func IsHelp(cmd *cobra.Command, args []string) bool {
	a := args[0]
	return cmd.DisableFlagParsing && (a == "-h" || a == "--help")
}
