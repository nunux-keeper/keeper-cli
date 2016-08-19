package cmd

import (
	"errors"
	"io"
	"text/template"

	"github.com/spf13/cobra"

	cmdutil "github.com/ncarlier/keeper-cli/cmd/util"
)

var DocumentTmpl = `Document:
 ID:          {{.Id}}
 Title:       {{.Title}}
 ContentType: {{.ContentType}}
 Content:     {{.Content}}
 Origin:      {{.Origin}}
 Date:        {{.Date}}
 Ghost:       {{.Ghost}}
`

func NewCmdGetDocument(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get (ID)",
		Short: "Get a document",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Document ID required.")
			}
			docid := args[0]

			return runGetDocument(f, out, cmd, docid)
		},
		ValidArgs: []string{"create", "rm", "restore", "destroy"},
	}
	return cmd
}

func runGetDocument(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, docid string) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	document, err := c.GetDocument(docid)
	if err != nil {
		return err
	}

	tmpl, err := template.New("document").Parse(DocumentTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(out, document)
	return err
}
