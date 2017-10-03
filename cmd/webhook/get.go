package webhook

import (
	"errors"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func newGetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "get (ID)",
		Short: "Get a webhook",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("webhook ID required")
			}
			id := args[0]

			return runGetCommand(cc, id)
		},
	}
}

func runGetCommand(cmd *cobra.Command, id string) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	resp, err := kli.API.GetWebhook(id)
	if err != nil {
		return err
	}
	return common.WriteCmdResponse(resp, common.WEBHOOK, kli.JSON)
}
