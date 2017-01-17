package cmd

import (
	"errors"
	"io"
	"text/template"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
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

func NewCmdGetUser(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user (UID)",
		Short: "Get user details (ADMIN)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("User ID required.")
			}
			uid := args[0]

			return runGetUser(f, out, cmd, uid)
		},
	}

	return cmd
}

func runGetUser(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, uid string) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	user, err := c.GetUser(uid)
	if err != nil {
		return err
	}

	tmpl, err := template.New("user").Parse(UserTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(out, user)
	return err
}
