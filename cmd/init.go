package cmd

import (
	"github.com/spf13/cobra"
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
		return err
	}

	konfigPath, err := konfig.FindKontextConfig()
	if err != nil {
		return err
	}

	exp, err := konfig.BuildKontextExport(konfigPath)
	if err != nil {
		return err
	}

	p, err := sh.Export(exp)
	if err != nil {
		return err
	}

	cmd.Print(p)

	return nil
}
