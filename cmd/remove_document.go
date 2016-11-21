package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

func NewCmdRemoveDocument(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm (ID)",
		Short: "Remove a document",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Document ID required.")
			}
			docid := args[0]

			return runRemoveDocument(f, out, cmd, docid)
		},
	}
	return cmd
}

func runRemoveDocument(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, docid string) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	err = c.RemoveDocument(docid)
	if err != nil {
		return err
	}

	fmt.Fprintln(out, "Document removed.")
	return nil
}
