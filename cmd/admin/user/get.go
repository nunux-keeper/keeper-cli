package user

import (
	"errors"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func newGetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "get (UID)",
		Short: "Get user details",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("User ID required.")
			}
			uid := args[0]
			return runGetCommand(cc, uid)
		},
	}
}

func runGetCommand(cmd *cobra.Command, uid string) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	resp, err := kli.API.GetUser(uid)
	if err != nil {
		return err
	}
	return common.WriteCmdResponse(resp, common.USER, kli.JSON)
}
