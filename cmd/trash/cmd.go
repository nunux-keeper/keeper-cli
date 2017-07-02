package trash

import (
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trash",
		Short: "Manage the trash",
		RunE:  common.ShowHelp(),
	}
	cmd.AddCommand(
		newEmptyCommand(),
		newListCommand(),
	)
	return cmd
}
