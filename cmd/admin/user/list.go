package user

import (
	"fmt"
	"text/tabwriter"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newListCommand(kCli *cli.KeeperCLI) *cobra.Command {
	return &cobra.Command{
		Use:   "ls",
		Short: "List users",
		RunE: func(cc *cobra.Command, args []string) error {
			return runListCommand(kCli, cc)
		},
	}
}

func runListCommand(kCli *cli.KeeperCLI, cmd *cobra.Command) error {
	users, err := kCli.APIClient.GetUsers()
	if err != nil {
		return err
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(*kCli.Out, 0, 8, 1, '\t', 0)
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
