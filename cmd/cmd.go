package cmd

import (
	"fmt"
	"io"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmdutil "github.com/nunux-keeper/keeper-cli/cmd/util"
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
		Use:                    "keepctl",
		Short:                  "Nunux Keeper v2 command-line interface",
		Long:                   `This CLI allow you to manage yours documents hoste on a Nunux Keeper instance.`,
		Run:                    runHelp,
		ValidArgs:              []string{"version", "login"},
		BashCompletionFunction: bash_completion_func,
	}

	cmds.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.keeper/config.yaml)")
	cmds.PersistentFlags().String("endpoint", "https://api.nunux.org/keeper", "API endpoint")

	viper.BindPFlag("endpoint", cmds.PersistentFlags().Lookup("endpoint"))

	cobra.OnInitialize(initConfig)

	cmds.AddCommand(NewCmdVersion(f, out))
	cmds.AddCommand(NewCmdInfos(f, out))
	cmds.AddCommand(NewCmdLogin(f, out))
	cmds.AddCommand(NewCmdLogout(f, out))
	cmds.AddCommand(NewCmdProfile(f, out))

	// get
	cmd_get := NewCmdGetDocument(f, out)
	cmd_get.AddCommand(NewCmdGetLabel(f, out))
	cmd_get.AddCommand(NewCmdGetUser(f, out))
	cmds.AddCommand(cmd_get)

	// create
	cmd_create := NewCmdCreateDocument(f, out)
	cmd_create.AddCommand(NewCmdCreateLabel(f, out))
	cmds.AddCommand(cmd_create)

	// rm
	cmd_rm := NewCmdRemoveDocument(f, out)
	cmd_rm.AddCommand(NewCmdRemoveLabel(f, out))
	cmds.AddCommand(cmd_rm)

	// restore
	cmd_re := NewCmdRestoreDocument(f, out)
	cmd_re.AddCommand(NewCmdRestoreLabel(f, out))
	cmds.AddCommand(cmd_re)

	// destroy
	cmds.AddCommand(NewCmdDestroyDocument(f, out))

	//ls
	cmd_ls := NewCmdListDocuments(f, out)
	cmd_ls.AddCommand(NewCmdListTrash(f, out))
	cmd_ls.AddCommand(NewCmdListLabels(f, out))
	cmd_ls.AddCommand(NewCmdListUsers(f, out))
	cmds.AddCommand(cmd_ls)

	// empty
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
