package job

import (
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "job",
		Short: "Manage jobs",
		RunE:  common.ShowHelp(*kCli.Out),
	}
	cmd.AddCommand(
		newCreateCommand(kCli),
		newGetCommand(kCli),
		newInfoCommand(kCli),
	)
	return cmd
}
