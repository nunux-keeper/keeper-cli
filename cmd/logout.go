package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/api"
	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

func NewCmdLogout(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout from a Nunux Keeper instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runLogout(out, cmd)
		},
	}
	return cmd
}

func runLogout(out io.Writer, cmd *cobra.Command) error {
	err := api.RemoveTokenInfos()
	if err != nil {
		return err
	}

	fmt.Fprintln(out, "User logged out.")
	return nil
}
