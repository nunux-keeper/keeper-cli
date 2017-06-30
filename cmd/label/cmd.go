package label

import (
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "label",
		Short: "Manage labels",
		RunE:  common.ShowHelp(*kCli.Out),
	}
	cmd.AddCommand(
		newCreateCommand(kCli),
		newGetCommand(kCli),
		newListCommand(kCli),
		newRemoveCommand(kCli),
		newRestoreCommand(kCli),
		// newDestroyCommand(kCli),
	)
	return cmd
}
