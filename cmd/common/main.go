package common

import (
	"os"

	"github.com/spf13/cobra"
)

func ShowHelp() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetOutput(os.Stderr)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}
