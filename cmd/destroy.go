package cmd

import (
	"errors"
	"fmt"

	"github.com/ncarlier/keeper-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy document from the graveyard",
	RunE:  destroyDocumentRun,
}

func destroyDocumentRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Document ID required.")
	}
	docid := args[0]

	kClient, err := api.NewKeeperAPIClient(viper.GetString("endpoint"))
	if err != nil {
		return err
	}

	err = kClient.DestroyDocument(docid)
	if err != nil {
		return err
	}

	fmt.Println("Document destroyed.")
	return nil
}

func init() {
	documentCmd.AddCommand(destroyCmd)
}
