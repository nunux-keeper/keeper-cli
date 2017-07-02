package job

import (
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func newInfoCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Get jobs informations",
		RunE: func(cc *cobra.Command, args []string) error {
			return runInfoCommand(cc)
		},
	}
}

func runInfoCommand(cmd *cobra.Command) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	resp, err := kli.API.GetJobsInfos()
	if err != nil {
		return err
	}
	return common.WriteCmdResponse(resp, common.JOBS_INFO, kli.JSON)
}
