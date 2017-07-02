package trash

import (
	"fmt"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newEmptyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "empty",
		Short: "Empty the trash",
		RunE: func(cc *cobra.Command, args []string) error {
			return runEmptyCommand(cc)
		},
	}
	return cmd
}

func runEmptyCommand(cmd *cobra.Command) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	err = kli.API.EmptyGraveyard()
	if err != nil {
		return err
	}

	fmt.Println("Trash is empty.")
	return nil
}
