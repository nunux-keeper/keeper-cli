package webhook

import (
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// NewCommand Create new command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhook",
		Short: "Manage webhooks",
		RunE:  common.ShowHelp(),
	}
	cmd.AddCommand(
		newCreateCommand(),
		newUpdateCommand(),
		newGetCommand(),
		newListCommand(),
		newRemoveCommand(),
	)
	return cmd
}
