package export

import (
	"os"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newStatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get status of an export",
		RunE: func(cc *cobra.Command, args []string) error {
			return runStatusCommand(cc)
		},
	}
}

func runStatusCommand(cmd *cobra.Command) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	err = kli.API.GetExportStatus(os.Stdout)
	if err != nil {
		return err
	}
	return nil
}
