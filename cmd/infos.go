package cmd

import (
	"io"
	"text/template"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

var serverInfosTmpl = `Server informations:
 Nb. users     {{.NbUsers}}
 Nb. documents {{.NbDocuments}}
`

func NewCmdInfos(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "infos",
		Short: "Get server informations (ADMIN)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetServerInfos(f, out, cmd)
		},
	}
	return cmd
}

func runGetServerInfos(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	resp, err := c.GetServerInfos()
	if err != nil {
		return err
	}

	tmpl, err := template.New("serverInfos").Parse(serverInfosTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(out, resp)
	return err
}
