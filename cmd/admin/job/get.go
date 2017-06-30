package job

import (
	"errors"
	"text/template"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
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

func newGetCommand(kCli *cli.KeeperCLI) *cobra.Command {
	return &cobra.Command{
		Use:   "get <id>",
		Short: "Get job details",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Job ID required.")
			}
			id := args[0]
			return runGetCommand(kCli, cc, id)
		},
	}
}

func runGetCommand(kCli *cli.KeeperCLI, cmd *cobra.Command, id string) error {
	res, err := kCli.APIClient.GetJob(id)
	if err != nil {
		return err
	}

	tmpl, err := template.New("jobDetails").Parse(JobDetailsTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(*kCli.Out, res)
	return err
}
