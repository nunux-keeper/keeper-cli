package profile

import (
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Get current user profile",
		RunE: func(cc *cobra.Command, args []string) error {
			return runProfileCommand(kCli, cc)
		},
	}
	return cmd
}

func runProfileCommand(kCli *cli.KeeperCLI, cmd *cobra.Command) error {
	resp, err := kCli.APIClient.GetProfile()
	if err != nil {
		return err
	}
	return common.WriteProfile(resp, *kCli.Out)
}
