package cmd

import (
	"io"
	"text/template"

	"github.com/spf13/cobra"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

var profileTmpl = `Profile:
 UID:   {{.Uid}}
 Name:  {{.Name}}
 Date:  {{.Date}}
 Admin: {{.Admin}}
`

func NewCmdProfile(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Get current user profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetProfile(f, out, cmd)
		},
	}
	return cmd
}

func runGetProfile(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	userProfile, err := c.GetProfile()
	if err != nil {
		return err
	}

	tmpl, err := template.New("profile").Parse(profileTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(out, userProfile)
	return err
}
