package export

import (
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

// NewCommand Declare new command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Manage exports",
		RunE:  common.ShowHelp(),
	}
	cmd.AddCommand(
		newScheduleCommand(),
		newStatusCommand(),
		newDownloadCommand(),
	)
	return cmd
}
