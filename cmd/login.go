package cmd

import (
	"fmt"

	"github.com/ncarlier/keeper-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type loginOptions struct {
	user     string
	password string
}

var opts loginOptions

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login in to a Nunux Keeper instance.",
	Long: `Login in to a Nunux Keeper instance.
If no server specified by the endpoint flag, the default is used.`,
	RunE: loginRun,
}

func loginRun(cmd *cobra.Command, args []string) error {
	kClient, err := api.NewKeeperAPIClient(viper.GetString("endpoint"))
	if err != nil {
		return err
	}

	infos, err := kClient.Login(opts.user, opts.password)
	if err != nil {
		return err
	}
	err = api.SaveTokenInfos(infos)
	if err != nil {
		return err
	}

	fmt.Printf("User %s logged.\n", opts.user)
	return nil
}

func init() {
	RootCmd.AddCommand(loginCmd)

	flags := loginCmd.Flags()

	flags.StringVarP(&opts.user, "username", "u", "", "Username")
	flags.StringVarP(&opts.password, "password", "p", "", "Password")
}
