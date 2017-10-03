package label

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
		Short: "List labels",
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

	labels, err := kli.API.GetLabels()
	if err != nil {
		return err
	}

	if kli.JSON {
		return common.WriteCmdJSONResponse(labels)
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "#\tLabel\tColor\tDate\t")
	for _, label := range labels {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", label.Id, label.Label, label.Color, label.Date)
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}
