package trash

import (
	"fmt"
	"text/tabwriter"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newListCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List trash content",
		RunE: func(cc *cobra.Command, args []string) error {
			return runListCommand(kCli, cc)
		},
	}
	return cmd
}

func runListCommand(kCli *cli.KeeperCLI, cmd *cobra.Command) error {

	documents, err := kCli.APIClient.GetGraveyard()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(*kCli.Out, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "#\tTitle\tContent-Type\tOrigin\tDate\t")
	for _, doc := range documents {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", doc.Id, doc.Title, doc.ContentType, doc.Origin, doc.Date)
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}
