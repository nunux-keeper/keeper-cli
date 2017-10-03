package webhook

import (
	"errors"
	"net/url"
	"strings"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/api"
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
)

type createWebhookOptions struct {
	secret string
	labels string
	events string
	active bool
}

func newCreateCommand() *cobra.Command {
	var opts createWebhookOptions
	cmd := &cobra.Command{
		Use:   "create (URL)",
		Short: "Create a webhook",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("http URL required")
			}
			link := args[0]
			u, err := url.ParseRequestURI(link)
			if err != nil {
				return errors.New("invalid http URL")
			}

			return runCreateCommand(cc, u, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.secret, "secret", "s", "", "Webhook secret")
	flags.BoolVarP(&opts.active, "active", "a", true, "Webhook activation status")
	flags.StringVarP(&opts.labels, "labels", "l", "", "Label filter (comma separated list)")
	flags.StringVarP(&opts.events, "events", "e", "", "Event filter (comma separated list)")
	return cmd
}

func runCreateCommand(cmd *cobra.Command, link *url.URL, opts *createWebhookOptions) error {
	wh := &api.WebhookResponse{
		URL:    link.String(),
		Secret: opts.secret,
		Active: opts.active,
		Labels: deleteEmpty(strings.Split(opts.labels, ",")),
		Events: deleteEmpty(strings.Split(opts.events, ",")),
	}
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	resp, err := kli.API.CreateWebhook(wh)
	if err != nil {
		return err
	}
	return common.WriteCmdResponse(resp, common.WEBHOOK, kli.JSON)
}
