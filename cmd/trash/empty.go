package trash

import (
	"fmt"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newEmptyCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "empty",
		Short: "Empty the trash",
		RunE: func(cc *cobra.Command, args []string) error {
			return runEmptyCommand(kCli, cc)
		},
	}
	return cmd
}

func runEmptyCommand(kCli *cli.KeeperCLI, cmd *cobra.Command) error {
	err := kCli.APIClient.EmptyGraveyard()
	if err != nil {
		return err
	}

	fmt.Fprintln(*kCli.Out, "Trash is empty.")
	return nil
}
