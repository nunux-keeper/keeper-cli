package cmd

import (
	"fmt"
	"io"
	"path"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// NewKeepCommand creates the `keepctl` command and its nested children.
func NewKeepctlCommand(input io.Reader, output, errorOutput io.Writer) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:       "keepctl",
		Short:     "Nunux Keeper v2 command-line interface",
		Long:      `This CLI allow you to manage yours documents hoste on a Nunux Keeper instance.`,
		Run:       runHelp,
		ValidArgs: []string{"version", "login"},
	}

	cmds.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.keeper/config.yaml)")
	cmds.PersistentFlags().String("endpoint", "https://api.nunux.org/keeper", "API endpoint")

	viper.BindPFlag("endpoint", cmds.PersistentFlags().Lookup("endpoint"))

	cobra.OnInitialize(initConfig)
	// Create new CLI
	cli, err := cli.NewKeeperCLI(input, output, errorOutput)
	if err != nil {
		panic(err)
	}

	// Add commands
	commands.AddCommands(cmds, cli)

	return cmds
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(path.Join("$HOME", ".keeper"))
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
