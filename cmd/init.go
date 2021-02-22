package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stackvista/kontext/pkg/env"
	"github.com/stackvista/kontext/pkg/konfig"
	"github.com/stackvista/kontext/pkg/shell"
)

func InitCommand() *cobra.Command {
	return &cobra.Command{
		Use:                "init <shell>",
		Args:               cobra.ExactArgs(1),
		DisableFlagParsing: true,
		RunE:               runInit,
	}
}

func runInit(cmd *cobra.Command, args []string) error {
	if IsHelp(cmd, args) {
		return cmd.Usage()
	}

	arg := args[0]

	sh, err := shell.DetectShell(arg)
	if err != nil {
		return nil
	}

	var exp env.Export
	konfigPath, err := konfig.FindKontextConfig()
	if err != nil {
		exp = konfig.UnsetKontextExport()
	} else {
		exp, err = konfig.BuildKontextExport(konfigPath)
		if err != nil {
			return nil
		}
	}

	p, err := sh.Export(exp)
	if err != nil {
		return nil
	}

	cmd.Print(p)

	return nil
}
