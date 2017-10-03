package export

import (
	"fmt"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

func newScheduleCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "schedule",
		Short: "Schedule a new export",
		RunE: func(cc *cobra.Command, args []string) error {
			return runScheduleCommand(cc)
		},
	}
}

func runScheduleCommand(cmd *cobra.Command) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}
	err = kli.API.ScheduleExport()
	if err != nil {
		return err
	}
	fmt.Println("Export scheduled.")
	return nil
}
