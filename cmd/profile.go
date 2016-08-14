package cmd

import (
	"os"
	"text/template"

	"github.com/ncarlier/keeper-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var profileTmpl = `Profile:
 UID:   {{.Uid}}
 Name:  {{.Name}}
 Date:  {{.Date}}
 Admin: {{.Admin}}
`

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Get current user profile.",
	RunE:  profileRun,
}

func profileRun(cmd *cobra.Command, args []string) error {
	kClient, err := api.NewKeeperAPIClient(viper.GetString("endpoint"))
	if err != nil {
		return err
	}

	userProfile, err := kClient.GetProfile()
	if err != nil {
		return err
	}

	tmpl, err := template.New("profile").Parse(profileTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, userProfile)
	return err
}

func init() {
	RootCmd.AddCommand(profileCmd)
}
