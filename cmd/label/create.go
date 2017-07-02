package label

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/api"
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
)

type createLabelOptions struct {
	label string
	color string
}

func newCreateCommand() *cobra.Command {
	var opts createLabelOptions
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a label",
		RunE: func(cc *cobra.Command, args []string) error {
			return runCreateCommand(cc, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.label, "label", "l", "", "Label name")
	flags.StringVarP(&opts.color, "color", "c", "", "Label color")
	return cmd
}

func runCreateCommand(cmd *cobra.Command, opts *createLabelOptions) error {
	if opts.label == "" || opts.color == "" {
		return errors.New("You have to specify at least a label and a color.")
	}

	lbl := &api.LabelResponse{
		Label: opts.label,
		Color: opts.color,
	}
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	resp, err := kli.API.CreateLabel(lbl)
	if err != nil {
		return err
	}
	return common.WriteCmdResponse(resp, common.LABEL, kli.JSON)
}
