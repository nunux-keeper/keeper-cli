package cmd

import (
	"errors"
	"fmt"

	"github.com/ncarlier/keeper-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a document",
	RunE:  rmDocumentRun,
}

func rmDocumentRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Document ID required.")
	}
	docid := args[0]

	kClient, err := api.NewKeeperAPIClient(viper.GetString("endpoint"))
	if err != nil {
		return err
	}

	err = kClient.RemoveDocument(docid)
	if err != nil {
		return err
	}

	fmt.Println("Document removed.")
	return nil
}

func init() {
	documentCmd.AddCommand(rmCmd)
}
