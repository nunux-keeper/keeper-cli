package profile

import (
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Get current user profile",
		RunE: func(cc *cobra.Command, args []string) error {
			return runProfileCommand(cc)
		},
	}
	return cmd
}

func runProfileCommand(cmd *cobra.Command) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	resp, err := kli.API.GetProfile()
	if err != nil {
		return err
	}
	return common.WriteCmdResponse(resp, common.PROFILE, kli.JSON)
}
