package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	cmdutil "github.com/ncarlier/keeper-cli/cmd/util"
)

func NewCmdEmptyTrash(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "empty",
		Short: "Empty the trash",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runEmptyTrash(f, out, cmd)
		},
	}
	return cmd
}

func runEmptyTrash(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	err = c.EmptyGraveyard()
	if err != nil {
		return err
	}

	fmt.Fprintln(out, "Trash is empty.")
	return nil
}
