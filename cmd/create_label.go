package cmd

import (
	"errors"
	"io"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/ncarlier/keeper-cli/api"
	cmdutil "github.com/ncarlier/keeper-cli/cmd/util"
)

type createLabelOptions struct {
	label string
	color string
}

func NewCmdCreateLabel(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	var opts createLabelOptions
	cmd := &cobra.Command{
		Use:   "label",
		Short: "Create a label",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreateLabel(f, out, cmd, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.label, "label", "l", "", "Label name")
	flags.StringVarP(&opts.color, "color", "c", "", "Label color")
	return cmd
}

func runCreateLabel(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, opts *createLabelOptions) error {
	if opts.label == "" || opts.color == "" {
		return errors.New("You have to specify at least a label and a color.")
	}

	c, err := f.Client()
	if err != nil {
		return err
	}

	lbl := &api.LabelResponse{
		Label: opts.label,
		Color: opts.color,
	}
	label, err := c.CreateLabel(lbl)
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
