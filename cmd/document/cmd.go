package document

import (
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "document",
		Aliases: []string{"doc"},
		Short:   "Manage documents",
		RunE:    common.ShowHelp(),
	}
	cmd.AddCommand(
		newCreateCommand(),
		newGetCommand(),
		newListCommand(),
		newRemoveCommand(),
		newRestoreCommand(),
		newDestroyCommand(),
	)
	return cmd
}
