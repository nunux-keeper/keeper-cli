package cmd

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

func NewCmdListLabels(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "labels",
		Short: "List labels",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListLabels(f, out, cmd)
		},
	}
	return cmd
}

func runListLabels(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	labels, err := c.GetLabels()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(out, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "#\tLabel\tColor\tDate\t")
	for _, label := range labels {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", label.Id, label.Label, label.Color, label.Date)
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}
