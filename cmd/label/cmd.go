package label

import (
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "label",
		Short: "Manage labels",
		RunE:  common.ShowHelp(),
	}
	cmd.AddCommand(
		newCreateCommand(),
		newGetCommand(),
		newListCommand(),
		newRemoveCommand(),
		newRestoreCommand(),
		// newDestroyCommand(),
	)
	return cmd
}
