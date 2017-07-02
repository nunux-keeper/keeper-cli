package admin

import (
	"github.com/nunux-keeper/keeper-cli/cmd/admin/info"
	"github.com/nunux-keeper/keeper-cli/cmd/admin/job"
	"github.com/nunux-keeper/keeper-cli/cmd/admin/user"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "admin",
		Short: "Admin commands",
		RunE:  common.ShowHelp(),
	}
	cmd.AddCommand(
		info.NewCommand(),
		job.NewCommand(),
		user.NewCommand(),
	)
	return cmd
}
