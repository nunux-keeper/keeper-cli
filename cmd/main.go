package cmd

import (
	"fmt"
	"path"

	"github.com/nunux-keeper/keeper-cli/cmd/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:       "keepctl",
		Short:     "Nunux Keeper v2 command-line interface",
		Long:      `This CLI allow you to manage yours documents hoste on a Nunux Keeper instance.`,
		ValidArgs: []string{"version", "login"},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	initViper()

	viper.SetDefault("json", false)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.keepctl/config.yaml)")
	rootCmd.PersistentFlags().String("endpoint", "https://api.nunux.org/keeper", "API endpoint")
	rootCmd.PersistentFlags().Bool("json", false, "Output as JSON")

	viper.BindPFlag("endpoint", rootCmd.PersistentFlags().Lookup("endpoint"))
	viper.BindPFlag("json", rootCmd.PersistentFlags().Lookup("json"))

	// Add commands
	commands.AddCommands(rootCmd)
}

func initViper() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".keepctl" (without extension).
		viper.AddConfigPath(path.Join("$HOME", ".keepctl"))
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
