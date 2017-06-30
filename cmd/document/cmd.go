package document

import (
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "document",
		Aliases: []string{"doc"},
		Short:   "Manage documents",
		RunE:    common.ShowHelp(*kCli.Out),
	}
	cmd.AddCommand(
		newCreateCommand(kCli),
		newGetCommand(kCli),
		newListCommand(kCli),
		newRemoveCommand(kCli),
		newRestoreCommand(kCli),
		newDestroyCommand(kCli),
	)
	return cmd
}
