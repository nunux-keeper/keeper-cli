package webhook

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List webhooks",
		RunE: func(cc *cobra.Command, args []string) error {
			return runListCommand(cc)
		},
	}
}

func runListCommand(cmd *cobra.Command) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	webhooks, err := kli.API.GetWebhooks()
	if err != nil {
		return err
	}

	if kli.JSON {
		return common.WriteCmdJSONResponse(webhooks)
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "#\tURL\tSecret\tEvents\tLabels\tDate")
	for _, webhook := range webhooks {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", webhook.ID, webhook.URL, webhook.Secret, webhook.Events, webhook.Labels, webhook.ModificationDate)
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}
