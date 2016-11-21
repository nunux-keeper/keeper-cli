package cmd

import (
	"errors"
	"io"
	"text/template"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

var LabelTmpl = `Label:
 Id:    {{.Id}}
 Label: {{.Label}}
 Color: {{.Color}}
 Date:  {{.Date}}
 Ghost: {{.Ghost}}
`

func NewCmdGetLabel(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "label (ID)",
		Short: "Get a label",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Label ID required.")
			}
			docid := args[0]

			return runGetLabel(f, out, cmd, docid)
		},
	}

	return cmd
}

func runGetLabel(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, id string) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	label, err := c.GetLabel(id)
	if err != nil {
		return err
	}

	tmpl, err := template.New("label").Parse(LabelTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(out, label)
	return err
}
