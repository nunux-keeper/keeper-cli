package label

import (
	"fmt"
	"text/tabwriter"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newListCommand(kCli *cli.KeeperCLI) *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List labels",
		RunE: func(cc *cobra.Command, args []string) error {
			return runListCommand(kCli, cc)
		},
	}
}

func runListCommand(kCli *cli.KeeperCLI, cmd *cobra.Command) error {
	labels, err := kCli.APIClient.GetLabels()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(*kCli.Out, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "#\tLabel\tColor\tDate\t")
	for _, label := range labels {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", label.Id, label.Label, label.Color, label.Date)
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}
