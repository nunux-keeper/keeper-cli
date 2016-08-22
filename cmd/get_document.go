package cmd

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/ncarlier/keeper-cli/api"
	cmdutil "github.com/ncarlier/keeper-cli/cmd/util"
)

var DocumentTmpl = `Document:
 Id:          {{.Id}}
 Title:       {{.Title}}
 ContentType: {{.ContentType}}
 Content:     {{.Content}}
 Origin:      {{.Origin}}
 Date:        {{.Date}}
 Ghost:       {{.Ghost}}
`

type getDocumentOptions struct {
	attribute string
}

func getField(d *api.DocumentResponse, field string) string {
	r := reflect.ValueOf(d)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func NewCmdGetDocument(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	var opts getDocumentOptions
	cmd := &cobra.Command{
		Use:   "get (ID)",
		Short: "Get a document (by default) or a label",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Document ID required.")
			}
			docid := args[0]

			return runGetDocument(f, out, cmd, docid, &opts)
		},
		ValidArgs: []string{"create", "rm", "restore", "destroy"},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.attribute, "attr", "a", "", "Attribute selection")

	return cmd
}

func runGetDocument(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, docid string, opts *getDocumentOptions) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	document, err := c.GetDocument(docid)
	if err != nil {
		return err
	}

	if opts.attribute != "" {
		fmt.Fprintln(out, getField(document, opts.attribute))
		return nil
	}

	tmpl, err := template.New("document").Parse(DocumentTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(out, document)
	return err
}
