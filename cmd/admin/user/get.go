package user

import (
	"errors"
	"text/template"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/spf13/cobra"
)

var UserTmpl = `User:
 Id:            {{.Id}}
 UID:           {{.Uid}}
 Name:          {{.Name}}
 Date:          {{.Date}}
 Nb. documents: {{.NbDocuments}}
 Nb. labels:    {{.NbLabels}}
 Nb. sharing:   {{.NbSharing}}
 Storage usage: {{.StorageUsage}}
`

func newGetCommand(kCli *cli.KeeperCLI) *cobra.Command {
	return &cobra.Command{
		Use:   "get (UID)",
		Short: "Get user details",
		RunE: func(cc *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("User ID required.")
			}
			uid := args[0]
			return runGetCommand(kCli, cc, uid)
		},
	}
}

func runGetCommand(kCli *cli.KeeperCLI, cmd *cobra.Command, uid string) error {
	user, err := kCli.APIClient.GetUser(uid)
	if err != nil {
		return err
	}

	tmpl, err := template.New("user").Parse(UserTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(*kCli.Out, user)
	return err
}
