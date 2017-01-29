package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/bgentry/speakeasy"
	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/api"
	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
)

type loginOptions struct {
	passwordInteractive bool
}

func NewCmdLogin(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	var opts loginOptions
	cmd := &cobra.Command{
		Use:   "login <uid>",
		Short: "Login to a Nunux Keeper instance",
		Long: `Login to a Nunux Keeper instance.
		If no server specified by the endpoint flag, the default is used.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return loginCommandFunc(f, out, cmd, &opts, args)
		},
	}
	flags := cmd.Flags()
	flags.BoolVar(&opts.passwordInteractive, "interactive", true, "If true, read password from stdin instead of interactive terminal")

	return cmd
}

func loginCommandFunc(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command, opts *loginOptions, args []string) error {
	if len(args) != 1 {
		return errors.New("UID required")
	}

	user := args[0]
	var password string

	c, err := f.Client()
	if err != nil {
		return err
	}

	if !opts.passwordInteractive {
		fmt.Scanf("%s", &password)
	} else {
		password, err = readPasswordInteractive(args[0])
		if err != nil {
			return err
		}
	}

	infos, err := c.Login(user, password)
	if err != nil {
		return err
	}
	err = api.SaveTokenInfos(infos)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "User %s logged.\n", user)
	return nil
}

func readPasswordInteractive(name string) (string, error) {
	prompt := fmt.Sprintf("Password of %s: ", name)
	password, err := speakeasy.Ask(prompt)
	if err != nil {
		return "", fmt.Errorf("failed to ask password: %v", err)
	}

	if len(password) == 0 {
		return "", errors.New("empty password")
	}

	return password, nil
}
