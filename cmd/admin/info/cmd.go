package info

import (
	"text/template"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

var serverInfosTmpl = `Server informations:
 Nb. users     {{.NbUsers}}
 Nb. documents {{.NbDocuments}}
`

func NewCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Get server informations",
		RunE: func(cc *cobra.Command, args []string) error {
			return runGetServerInfos(kCli, cc)
		},
	}
	return cmd
}

func runGetServerInfos(kCli *cli.KeeperCLI, cmd *cobra.Command) error {
	resp, err := kCli.APIClient.GetServerInfos()
	if err != nil {
		return err
	}

	tmpl, err := template.New("serverInfos").Parse(serverInfosTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(*kCli.Out, resp)
	return err
}
