package cmd

import (
	"errors"
	"os"
	"text/template"

	"github.com/ncarlier/keeper-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var documentTmpl = `Document:
 ID:          {{.Id}}
 Title:       {{.Title}}
 ContentType: {{.ContentType}}
 Content:     {{.Content}}
 Origin:      {{.Origin}}
 Date:        {{.Date}}
 Ghost:       {{.Ghost}}
`

// documentCmd represents the document command
var documentCmd = &cobra.Command{
	Use:   "document DOCID",
	Short: "Get a document",
	RunE:  documentRun,
}

func documentRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("Document ID required.")
	}
	docid := args[0]

	kClient, err := api.NewKeeperAPIClient(viper.GetString("endpoint"))
	if err != nil {
		return err
	}

	document, err := kClient.GetDocument(docid)
	if err != nil {
		return err
	}

	tmpl, err := template.New("document").Parse(documentTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, document)
	return err

}

func init() {
	RootCmd.AddCommand(documentCmd)
}
