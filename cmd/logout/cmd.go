package logout

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/api"
	"github.com/nunux-keeper/keeper-cli/cli"
)

func NewCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout from a Nunux Keeper instance",
		RunE: func(cc *cobra.Command, args []string) error {
			return runLogoutCommand(kCli, cc)
		},
	}
	return cmd
}

func runLogoutCommand(kCli *cli.KeeperCLI, cmd *cobra.Command) error {
	err := api.RemoveTokenInfos()
	if err != nil {
		return err
	}

	fmt.Fprintln(*kCli.Out, "User logged out.")
	return nil
}
