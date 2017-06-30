package trash

import (
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trash",
		Short: "Manage the trash",
		RunE:  common.ShowHelp(*kCli.Out),
	}
	cmd.AddCommand(
		newEmptyCommand(kCli),
		newListCommand(kCli),
	)
	return cmd
}
