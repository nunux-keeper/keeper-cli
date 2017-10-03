package webhook

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/api"
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
)

type updateWebhookOptions struct {
	ID     string
	url    string
	secret string
	labels string
	events string
	active bool
}

func newUpdateCommand() *cobra.Command {
	var opts updateWebhookOptions
	cmd := &cobra.Command{
		Use:   "update (ID)",
		Short: "Update a webhook",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("webhook ID required")
			}
			id := args[0]
			return runUpdateCommand(cc, id, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.url, "url", "u", "", "Webhook URL")
	flags.StringVarP(&opts.secret, "secret", "s", "", "Webhook secret")
	flags.BoolVarP(&opts.active, "active", "a", true, "Webhook activation status")
	flags.StringVarP(&opts.labels, "labels", "l", "", "Label filter (comma separated list)")
	flags.StringVarP(&opts.events, "events", "e", "", "Event filter (comma separated list)")
	return cmd
}

func runUpdateCommand(cmd *cobra.Command, id string, opts *updateWebhookOptions) error {
	wh := &api.WebhookResponse{
		URL:    opts.url,
		Secret: opts.secret,
		Active: opts.active,
		Labels: deleteEmpty(strings.Split(opts.labels, ",")),
		Events: deleteEmpty(strings.Split(opts.events, ",")),
	}
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	resp, err := kli.API.UpdateWebhook(id, wh)
	if err != nil {
		return err
	}
	return common.WriteCmdResponse(resp, common.WEBHOOK, kli.JSON)
}
