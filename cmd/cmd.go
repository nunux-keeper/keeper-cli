package cmd

import (
	"fmt"
	"io"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmdutil "github.com/ncarlier/keeper-cli/cmd/util"
)

const bash_completion_func = `# call keepctl get $1,
__keepctl_parse_get() {
    local keepctl_output out
    if keepctl_output=$(keepctl ls --no-headers" 2>/dev/null); then
        out=($(echo "${keepctl_output}" | awk '{print $1}'))
        COMPREPLY=( $( compgen -W "${out[*]}" -- "$cur" ) )
    fi
}

__keepctl_get_resource() {
    __keepctl_parse_get ${nouns[${#nouns[@]} -1]}
    if [[ $? -eq 0 ]]; then
        return 0
    fi
}

__custom_func() {
    case ${last_command} in
        keepctl_get | keepctl_remove | keepctl_restore | keepctl_destroy)
            __keepctl_get_resource
            return
            ;;
        *)
            ;;
    esac
}
`

var cfgFile string

// NewKeepCommand creates the `keepctl` command and its nested children.
func NewKeepctlCommand(f *cmdutil.Factory, in io.Reader, out, err io.Writer) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   "keepctl",
		Short: "Nunux Keeper v2 command-line interface",
		Long:  `This CLI allow you to manage yours documents hoste on a Nunux Keeper instance.`,
		Run:   runHelp,
		BashCompletionFunction: bash_completion_func,
		ValidArgs:              []string{"version", "login"},
	}

	cmds.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.keeper/config.yaml)")
	cmds.PersistentFlags().String("endpoint", "https://api.nunux.org/keeper", "API endpoint")

	viper.BindPFlag("endpoint", cmds.PersistentFlags().Lookup("endpoint"))

	cobra.OnInitialize(initConfig)

	cmds.AddCommand(NewCmdVersion(f, out))
	cmds.AddCommand(NewCmdLogin(f, out))
	cmds.AddCommand(NewCmdLogout(f, out))
	cmds.AddCommand(NewCmdProfile(f, out))
	cmds.AddCommand(NewCmdGetDocument(f, out))
	cmds.AddCommand(NewCmdCreateDocument(f, out))
	cmds.AddCommand(NewCmdRemoveDocument(f, out))
	cmds.AddCommand(NewCmdRestoreDocument(f, out))
	cmds.AddCommand(NewCmdDestroyDocument(f, out))
	cmdList := NewCmdListDocuments(f, out)
	cmdList.AddCommand(NewCmdListTrash(f, out))
	cmdList.AddCommand(NewCmdListLabels(f, out))
	cmds.AddCommand(cmdList)
	cmds.AddCommand(NewCmdEmptyTrash(f, out))

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
