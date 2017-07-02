package label

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/cli"
)

func newDestroyCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "destroy (ID)",
		Short: "Remove a label from the trash",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Label ID required.")
			}
			docid := args[0]
			return runDestroyCommand(cc, docid)
		},
	}
}

func runDestroyCommand(cmd *cobra.Command, id string) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	err = kli.API.DestroyLabel(id)
	if err != nil {
		return err
	}
	fmt.Println("Label destroyed.")
	return nil
}
