package export

import (
	"errors"
	"os"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

type downloadExportOptions struct {
	output string
}

func newDownloadCommand() *cobra.Command {
	var opts downloadExportOptions
	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download an export",
		RunE: func(cc *cobra.Command, args []string) error {
			return runDownloadCommand(cc, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.output, "output", "o", "export.zip", "Export file location")
	return cmd
}

func runDownloadCommand(cmd *cobra.Command, opts *downloadExportOptions) error {
	if opts.output == "" {
		return errors.New("you have to specify at least the export location")
	}

	// Create the file
	out, err := os.Create(opts.output)
	if err != nil {
		return err
	}
	defer out.Close()

	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	err = kli.API.DownloadExport(out)
	if err != nil {
		return err
	}
	return nil
}
