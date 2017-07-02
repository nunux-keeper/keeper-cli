package trash

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List trash content",
		RunE: func(cc *cobra.Command, args []string) error {
			return runListCommand(cc)
		},
	}
	return cmd
}

func runListCommand(cmd *cobra.Command) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	documents, err := kli.API.GetGraveyard()
	if err != nil {
		return err
	}

	if kli.JSON {
		return common.WriteCmdJsonResponse(documents)
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "#\tTitle\tContent-Type\tOrigin\tDate\t")
	for _, doc := range documents {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", doc.Id, doc.Title, doc.ContentType, doc.Origin, doc.Date)
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}
