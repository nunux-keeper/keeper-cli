package document

import (
	"errors"
	"fmt"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "rm (ID)",
		Short: "Remove a document",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Document ID required.")
			}
			docid := args[0]
			return runRemoveCommand(cc, docid)
		},
	}
}

func runRemoveCommand(cmd *cobra.Command, docid string) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	err = kli.API.RemoveDocument(docid)
	if err != nil {
		return err
	}
	fmt.Println("Document removed.")
	return nil
}
