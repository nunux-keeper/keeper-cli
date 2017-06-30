package user

import (
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
		RunE:  common.ShowHelp(*kCli.Out),
	}
	cmd.AddCommand(
		newGetCommand(kCli),
		newListCommand(kCli),
	)
	return cmd
}
