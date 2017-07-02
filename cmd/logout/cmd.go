package logout

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/api"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout from a Nunux Keeper instance",
		RunE: func(cc *cobra.Command, args []string) error {
			return runLogoutCommand(cc)
		},
	}
	return cmd
}

func runLogoutCommand(cmd *cobra.Command) error {
	err := api.RemoveTokenInfos()
	if err != nil {
		return err
	}

	fmt.Println("User logged out.")
	return nil
}
