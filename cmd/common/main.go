package common

import (
	"io"

	"github.com/spf13/cobra"
)

func ShowHelp(err io.Writer) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetOutput(err)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}
