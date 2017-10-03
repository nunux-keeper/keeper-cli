package user

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
		Short: "List users",
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

	users, err := kli.API.GetUsers()
	if err != nil {
		return err
	}

	if kli.JSON {
		return common.WriteCmdJSONResponse(users)
	}

	w := new(tabwriter.Writer)
	// Format in tab-separated columns with a tab stop of 8.
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
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
