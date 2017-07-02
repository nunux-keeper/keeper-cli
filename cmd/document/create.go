package document

import (
	"bufio"
	"bytes"
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/api"
	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
)

type createOptions struct {
	title   string
	content string
	url     string
}

func newCreateCommand() *cobra.Command {
	var opts createOptions
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a document",
		RunE: func(cc *cobra.Command, args []string) error {
			return runCreateCommand(cc, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.title, "title", "t", "", "Document title")
	flags.StringVarP(&opts.content, "content", "c", "", "Document content")
	flags.StringVarP(&opts.url, "url", "u", "", "Document URL")
	return cmd
}

func runCreateCommand(cmd *cobra.Command, opts *createOptions) error {
	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	var content string
	fi, _ := os.Stdin.Stat()
	if fi.Mode()&os.ModeNamedPipe == 0 {
		content = opts.content
	} else {
		reader := bufio.NewReader(os.Stdin)
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		content = buf.String()
	}

	if opts.title == "" && content == "" && opts.url == "" {
		return errors.New("You have to specify at least a title, a content or an url.")
	}

	doc := &api.DocumentResponse{
		Title:   opts.title,
		Content: content,
		Origin:  opts.url,
	}
	resp, err := kli.API.CreateDocument(doc)
	if err != nil {
		return err
	}
	return common.WriteCmdResponse(resp, common.DOCUMENT, kli.JSON)
}
