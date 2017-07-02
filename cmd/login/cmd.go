package login

import (
	"errors"
	"fmt"

	"github.com/bgentry/speakeasy"
	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/api"
	"github.com/nunux-keeper/keeper-cli/cli"
)

type loginOptions struct {
	passwordInteractive bool
}

func NewCommand() *cobra.Command {
	var opts loginOptions
	cmd := &cobra.Command{
		Use:   "login <uid>",
		Short: "Login to a Nunux Keeper instance",
		Long: `Login to a Nunux Keeper instance.
		If no server specified by the endpoint flag, the default is used.`,
		RunE: func(cc *cobra.Command, args []string) error {
			return runLoginCommand(cc, &opts, args)
		},
	}
	flags := cmd.Flags()
	flags.BoolVar(&opts.passwordInteractive, "interactive", true, "If true, read password from stdin instead of interactive terminal")

	return cmd
}

func runLoginCommand(cmd *cobra.Command, opts *loginOptions, args []string) error {
	if len(args) != 1 {
		return errors.New("UID required")
	}

	user := args[0]
	var password string
	var err error

	if !opts.passwordInteractive {
		fmt.Scanf("%s", &password)
	} else {
		password, err = readPasswordInteractive(args[0])
		if err != nil {
			return err
		}
	}

	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	infos, err := kli.API.Login(user, password)
	if err != nil {
		return err
	}
	err = api.SaveTokenInfos(infos)
	if err != nil {
		return err
	}

	fmt.Printf("User %s logged.\n", user)
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
