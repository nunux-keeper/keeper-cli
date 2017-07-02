package label

import (
	"errors"
	"fmt"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "rm (ID)",
		Short: "Remove a label",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Label ID required.")
			}
			docid := args[0]
			return runRemoveCommand(cmd, docid)
		},
	}
}

func runRemoveCommand(cmd *cobra.Command, id string) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	err = kli.API.RemoveLabel(id)
	if err != nil {
		return err
	}
	fmt.Println("Label removed.")
	return nil
}
