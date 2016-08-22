package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/ncarlier/keeper-cli/api"
	cmdutil "github.com/ncarlier/keeper-cli/cmd/util"
)

type createOptions struct {
	title   string
	content string
	url     string
}

func NewCmdCreateDocument(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	var opts createOptions
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a document",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreateDocument(f, out, cmd, &opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.title, "title", "t", "", "Document title")
	flags.StringVarP(&opts.content, "content", "c", "", "Document content")
	flags.StringVarP(&opts.url, "url", "u", "", "Document URL")
	return cmd
}

func runCreateDocument(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, opts *createOptions) error {
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

	c, err := f.Client()
	if err != nil {
		return err
	}

	doc := &api.DocumentResponse{
		Title:   opts.title,
		Content: content,
		Origin:  opts.url,
	}
	document, err := c.CreateDocument(doc)
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
