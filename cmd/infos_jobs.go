package cmd

import (
	"html/template"
	"io"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

var jobsInfosTmpl = `Jobs informations:
 Nb. inactive  {{.InactiveCount}}
 Nb. complete  {{.CompleteCount}}
 Nb. active    {{.ActiveCount}}
 Nb. failed    {{.FailedCount}}
 Work time     {{.WorkTime}}
`

func NewCmdInfosJobs(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jobs",
		Short: "Get jobs informations (ADMIN)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetJobsInfos(f, out, cmd)
		},
	}
	return cmd
}

func runGetJobsInfos(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	resp, err := c.GetJobsInfos()
	if err != nil {
		return err
	}

	tmpl, err := template.New("jobsInfos").Parse(jobsInfosTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(out, resp)
	return err
}
