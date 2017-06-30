package document

import (
	"errors"
	"fmt"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newRemoveCommand(kCli *cli.KeeperCLI) *cobra.Command {
	return &cobra.Command{
		Use:   "rm (ID)",
		Short: "Remove a document",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Document ID required.")
			}
			docid := args[0]
			return runRemoveCommand(kCli, cc, docid)
		},
	}
}

func runRemoveCommand(kCli *cli.KeeperCLI, cmd *cobra.Command, docid string) error {
	err := kCli.APIClient.RemoveDocument(docid)
	if err != nil {
		return err
	}
	fmt.Fprintln(*kCli.Out, "Document removed.")
	return nil
}
