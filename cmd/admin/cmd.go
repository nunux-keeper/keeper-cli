package admin

import (
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/admin/info"
	"github.com/nunux-keeper/keeper-cli/cmd/admin/job"
	"github.com/nunux-keeper/keeper-cli/cmd/admin/user"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admin",
		Short: "Admin commands",
		RunE:  common.ShowHelp(*kCli.Out),
	}
	cmd.AddCommand(
		info.NewCommand(kCli),
		job.NewCommand(kCli),
		user.NewCommand(kCli),
	)
	return cmd
}
