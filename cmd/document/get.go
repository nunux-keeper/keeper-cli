package document

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/api"
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
)

type getDocumentOptions struct {
	attribute string
}

func getField(d *api.DocumentResponse, field string) string {
	r := reflect.ValueOf(d)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func newGetCommand(kCli *cli.KeeperCLI) *cobra.Command {
	var opts getDocumentOptions
	cmd := &cobra.Command{
		Use:   "get (ID)",
		Short: "Get a document",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Document ID required.")
			}
			docid := args[0]
			return runGetCommand(kCli, cc, docid, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.attribute, "attr", "a", "", "Attribute selection")
	return cmd
}

func runGetCommand(kCli *cli.KeeperCLI, cmd *cobra.Command, docid string, opts *getDocumentOptions) error {
	document, err := kCli.APIClient.GetDocument(docid)
	if err != nil {
		return err
	}

	if opts.attribute != "" {
		fmt.Fprintln(*kCli.Out, getField(document, opts.attribute))
		return nil
	}
	return common.WriteDocument(document, *kCli.Out)
}
