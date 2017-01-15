package cmd

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

func NewCmdListUsers(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "List users",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListUsers(f, out, cmd)
		},
	}
	return cmd
}

func runListUsers(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	users, err := c.GetUsers()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(out, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "#\tUID\tName\tDate\tNb. documents\tNb. labels\tNb. sharing\tStorage usage\t")
	for _, user := range users {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%d\t%d\t%d\t\n",
			user.Id, user.Uid, user.Name, user.Date,
			user.NbDocuments, user.NbLabels, user.NbSharing, user.StorageUsage)
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}
