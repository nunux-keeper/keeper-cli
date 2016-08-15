package cmd

import (
	"errors"
	"os"
	"text/template"

	"github.com/ncarlier/keeper-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var docOptions api.DocumentResponse

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a document",
	RunE:  createDocumentRun,
}

func createDocumentRun(cmd *cobra.Command, args []string) error {
	if docOptions.Title == "" && docOptions.Content == "" && docOptions.Origin == "" {
		return errors.New("Use at least a flag to create a valid document.")
	}

	kClient, err := api.NewKeeperAPIClient(viper.GetString("endpoint"))
	if err != nil {
		return err
	}

	document, err := kClient.CreateDocument(&docOptions)
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
	documentCmd.AddCommand(createCmd)

	flags := createCmd.Flags()

	flags.StringVarP(&docOptions.Title, "title", "t", "", "Document title")
	flags.StringVarP(&docOptions.Content, "content", "c", "", "Document content")
	flags.StringVarP(&docOptions.Origin, "url", "u", "", "Document URL")
}
