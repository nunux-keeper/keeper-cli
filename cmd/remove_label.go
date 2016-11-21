package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

func NewCmdRemoveLabel(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "label (ID)",
		Short: "Remove a label",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Label ID required.")
			}
			docid := args[0]

			return runRemoveLabel(f, out, cmd, docid)
		},
	}
	return cmd
}

func runRemoveLabel(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, id string) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	err = c.RemoveLabel(id)
	if err != nil {
		return err
	}

	fmt.Fprintln(out, "Label removed.")
	return nil
}
