package webhook

import (
	"errors"
	"fmt"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "rm (ID)",
		Short: "Remove a webhook",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("webhook ID required")
			}
			id := args[0]
			return runRemoveCommand(cmd, id)
		},
	}
}

func runRemoveCommand(cmd *cobra.Command, id string) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	err = kli.API.RemoveWebhook(id)
	if err != nil {
		return err
	}
	fmt.Println("webhook removed.")
	return nil
}
