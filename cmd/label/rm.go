package label

import (
	"errors"
	"fmt"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newRemoveCommand(kCli *cli.KeeperCLI) *cobra.Command {
	return &cobra.Command{
		Use:   "rm (ID)",
		Short: "Remove a label",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Label ID required.")
			}
			docid := args[0]
			return runRemoveCommand(kCli, cmd, docid)
		},
	}
}

func runRemoveCommand(kCli *cli.KeeperCLI, cmd *cobra.Command, id string) error {
	err := kCli.APIClient.RemoveLabel(id)
	if err != nil {
		return err
	}
	fmt.Fprintln(*kCli.Out, "Label removed.")
	return nil
}
