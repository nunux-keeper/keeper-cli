package cmd

import (
	"errors"
	"io"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/ncarlier/keeper-cli/api"
	cmdutil "github.com/ncarlier/keeper-cli/cmd/util"
)

func NewCmdCreateDocument(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	var doc api.DocumentResponse
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a document",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreateDocument(f, out, cmd, &doc)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&doc.Title, "title", "t", "", "Document title")
	flags.StringVarP(&doc.Content, "content", "c", "", "Document content")
	flags.StringVarP(&doc.Origin, "url", "u", "", "Document URL")
	return cmd
}

func runCreateDocument(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, doc *api.DocumentResponse) error {
	if doc.Title == "" && doc.Content == "" && doc.Origin == "" {
		return errors.New("You have to specify at least a title, a content or an url.")
	}

	c, err := f.Client()
	if err != nil {
		return err
	}

	document, err := c.CreateDocument(doc)
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
