package label

import (
	"errors"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func newGetCommand(kCli *cli.KeeperCLI) *cobra.Command {
	return &cobra.Command{
		Use:   "get (ID)",
		Short: "Get a label",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Label ID required.")
			}
			docid := args[0]

			return runGetCommand(kCli, cc, docid)
		},
	}
}

func runGetCommand(kCli *cli.KeeperCLI, cmd *cobra.Command, id string) error {
	label, err := kCli.APIClient.GetLabel(id)
	if err != nil {
		return err
	}
	return common.WriteLabel(label, *kCli.Out)
}
