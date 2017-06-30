package document

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/cli"
)

func newDestroyCommand(kCli *cli.KeeperCLI) *cobra.Command {
	return &cobra.Command{
		Use:   "destroy (ID)",
		Short: "Remove a document from the trash",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Document ID required.")
			}
			docid := args[0]
			return runDestroyCommand(kCli, cc, docid)
		},
	}
}

func runDestroyCommand(kCli *cli.KeeperCLI, cmd *cobra.Command, docid string) error {
	err := kCli.APIClient.DestroyDocument(docid)
	if err != nil {
		return err
	}
	fmt.Fprintln(*kCli.Out, "Document destroyed.")
	return nil
}
