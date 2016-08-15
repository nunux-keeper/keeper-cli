package cmd

import (
	"errors"
	"os"
	"text/template"

	"github.com/ncarlier/keeper-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a deleted document",
	RunE:  restoreDocumentRun,
}

func restoreDocumentRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Document ID required.")
	}
	docid := args[0]

	kClient, err := api.NewKeeperAPIClient(viper.GetString("endpoint"))
	if err != nil {
		return err
	}

	document, err := kClient.RestoreDocument(docid)
	if err != nil {
		return err
	}

	tmpl, err := template.New("document").Parse(DocumentTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, document)
	return err
}

func init() {
	documentCmd.AddCommand(restoreCmd)
}
