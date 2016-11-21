package cmd

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/spf13/cobra"

	cmdutil "github.com/ncarlier/keeper-cli/cmd/util"
)

type listOptions struct {
	noHeaders bool
	inverted  bool
	query     string
	size      int
	from      int
}

func NewCmdListDocuments(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	var opts listOptions
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List documents (by default), labels or trash content",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListDocuments(f, out, cmd, &opts)
		},
		ValidArgs: []string{"labels", "trash"},
	}
	flags := cmd.Flags()
	flags.BoolVar(&opts.noHeaders, "no-headers", false, "Hide headers")
	flags.BoolVar(&opts.inverted, "invert", false, "Invert order (from oldest)")
	flags.StringVarP(&opts.query, "query", "q", "", "Query search")
	flags.IntVarP(&opts.size, "size", "s", 50, "Result size limit")
	flags.IntVarP(&opts.from, "from", "f", 0, "Result begin index")

	return cmd
}

func runListDocuments(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, opts *listOptions) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	order := "desc"
	if opts.inverted {
		order = "asc"
	}

	documents, err := c.GetDocuments(opts.query, order, opts.size, opts.from)
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(out, 0, 8, 1, '\t', 0)
	if !opts.noHeaders {
		fmt.Fprintln(w, "#\tTitle\tContent-Type\tOrigin\tDate\t")
	}
	for _, doc := range documents {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", doc.Id, doc.Title, doc.ContentType, doc.Origin, doc.Date)
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}
