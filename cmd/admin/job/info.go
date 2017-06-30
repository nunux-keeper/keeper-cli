package job

import (
	"html/template"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

var jobsInfosTmpl = `Jobs informations:
 Nb. inactive  {{.InactiveCount}}
 Nb. complete  {{.CompleteCount}}
 Nb. active    {{.ActiveCount}}
 Nb. failed    {{.FailedCount}}
 Work time     {{.WorkTime}}
`

func newInfoCommand(kCli *cli.KeeperCLI) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Get jobs informations",
		RunE: func(cc *cobra.Command, args []string) error {
			return runInfoCommand(kCli, cc)
		},
	}
}

func runInfoCommand(kCli *cli.KeeperCLI, cmd *cobra.Command) error {
	resp, err := kCli.APIClient.GetJobsInfos()
	if err != nil {
		return err
	}

	tmpl, err := template.New("jobsInfos").Parse(jobsInfosTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(*kCli.Out, resp)
	return err
}
