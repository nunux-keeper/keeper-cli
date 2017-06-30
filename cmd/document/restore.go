package document

import (
	"errors"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func newRestoreCommand(kCli *cli.KeeperCLI) *cobra.Command {
	return &cobra.Command{
		Use:   "restore (ID)",
		Short: "Restore a deleted document",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Document ID required.")
			}
			docid := args[0]
			return runRestoreCommand(kCli, cc, docid)
		},
	}
}

func runRestoreCommand(kCli *cli.KeeperCLI, cmd *cobra.Command, docid string) error {
	document, err := kCli.APIClient.RestoreDocument(docid)
	if err != nil {
		return err
	}
	return common.WriteDocument(document, *kCli.Out)
}
