package cmd

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/spf13/cobra"

	cmdutil "github.com/ncarlier/keeper-cli/cmd/util"
)

func NewCmdListTrash(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trash",
		Short: "List trash content",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListTrash(f, out, cmd)
		},
	}
	return cmd
}

func runListTrash(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	documents, err := c.GetGraveyard()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(out, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "#\tTitle\tContent-Type\tOrigin\tDate\t")
	for _, doc := range documents {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", doc.Id, doc.Title, doc.ContentType, doc.Origin, doc.Date)
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}
