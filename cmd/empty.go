package cmd

import (
	"fmt"

	"github.com/ncarlier/keeper-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// emptyCmd represents the empty command
var emptyCmd = &cobra.Command{
	Use:   "empty",
	Short: "Empty the graveyard",
	RunE:  emptyGraveyardRun,
}

func emptyGraveyardRun(cmd *cobra.Command, args []string) error {
	kClient, err := api.NewKeeperAPIClient(viper.GetString("endpoint"))
	if err != nil {
		return err
	}

	err = kClient.EmptyGraveyard()
	if err != nil {
		return err
	}

	fmt.Println("Graveyard is empty.")
	return nil
}

func init() {
	graveyardCmd.AddCommand(emptyCmd)
}
