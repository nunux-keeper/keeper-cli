package user

import (
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
		RunE:  common.ShowHelp(),
	}
	cmd.AddCommand(
		newGetCommand(),
		newListCommand(),
	)
	return cmd
}
