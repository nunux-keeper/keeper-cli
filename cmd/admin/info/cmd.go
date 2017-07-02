package info

import (
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Get server informations",
		RunE: func(cc *cobra.Command, args []string) error {
			return runGetServerInfos(cc)
		},
	}
	return cmd
}

func runGetServerInfos(cmd *cobra.Command) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}
	resp, err := kli.API.GetServerInfos()
	if err != nil {
		return err
	}
	return common.WriteCmdResponse(resp, common.SERVER_INFOS, kli.JSON)
}
