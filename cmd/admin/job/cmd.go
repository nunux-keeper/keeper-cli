package job

import (
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "job",
		Short: "Manage jobs",
		RunE:  common.ShowHelp(),
	}
	cmd.AddCommand(
		newCreateCommand(),
		newGetCommand(),
		newInfoCommand(),
	)
	return cmd
}
