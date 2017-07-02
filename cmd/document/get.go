package document

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
)

func newGetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "get (ID)",
		Short: "Get a document",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Document ID required.")
			}
			docid := args[0]
			return runGetCommand(cc, docid)
		},
	}
}

func runGetCommand(cmd *cobra.Command, docid string) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	resp, err := kli.API.GetDocument(docid)
	if err != nil {
		return err
	}

	return common.WriteCmdResponse(resp, common.DOCUMENT, kli.JSON)
}
