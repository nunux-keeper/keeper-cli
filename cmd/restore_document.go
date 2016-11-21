package cmd

import (
	"errors"
	"io"
	"text/template"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

func NewCmdRestoreDocument(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restore (ID)",
		Short: "Restore a deleted document",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Document ID required.")
			}
			docid := args[0]

			return runRestoreDocument(f, out, cmd, docid)
		},
	}
	return cmd
}

func runRestoreDocument(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, docid string) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	document, err := c.RestoreDocument(docid)
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
