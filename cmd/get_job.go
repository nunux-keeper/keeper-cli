package cmd

import (
	"errors"
	"io"
	"text/template"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

var JobDetailsTmpl = `Job:
 Id:         {{.Id}}
 Type:       {{.Type}}
 Priority:   {{.Priority}}
 Progress:   {{.Progress}}
 State:      {{.State}}
 Created at  {{.CreatedAt}}
 Updated at: {{.UpdatedAt}}
 Duration:   {{.Duration}}
 Params:     {{.Data}}
`

func NewCmdGetJob(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "job <id>",
		Short: "Get job details (ADMIN)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Job ID required.")
			}
			id := args[0]

			return runGetJob(f, out, cmd, id)
		},
	}

	return cmd
}

func runGetJob(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, id string) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	res, err := c.GetJob(id)
	if err != nil {
		return err
	}

	tmpl, err := template.New("jobDetails").Parse(JobDetailsTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(out, res)
	return err
}
