package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/ncarlier/keeper-cli/api"
	cmdutil "github.com/ncarlier/keeper-cli/cmd/util"
)

type loginOptions struct {
	user     string
	password string
}

func NewCmdLogin(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	var opts loginOptions
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to a Nunux Keeper instance",
		Long: `Login to a Nunux Keeper instance.
		If no server specified by the endpoint flag, the default is used.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runLogin(f, out, cmd, &opts)
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&opts.user, "username", "u", "", "Username")
	flags.StringVarP(&opts.password, "password", "p", "", "Password")

	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")

	return cmd
}

func runLogin(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, opts *loginOptions) error {
	c, err := f.Client()
	if err != nil {
		return err
	}

	infos, err := c.Login(opts.user, opts.password)
	if err != nil {
		return err
	}
	err = api.SaveTokenInfos(infos)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "User %s logged.\n", opts.user)
	return nil
}
