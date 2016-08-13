package cmd

import (
	"log"
	"os"
	"runtime"
	"text/template"

	"github.com/ncarlier/keeper-cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionTmpl = `Client:
 Version:      {{.Client.Version}}
 API version:  {{.Client.APIVersion}}
 Go version:   {{.Client.GoVersion}}
 OS/Arch:      {{.Client.Os}}/{{.Client.Arch}}

Server:
 Version:      {{.Server.Version}}
 API version:  {{.Server.APIVersion}}`

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print client and server versions",
	Run:   versionRun,
}

type Version struct {
	Version    string
	APIVersion string
	GoVersion  string
	Os         string
	Arch       string
}

type VersionResponse struct {
	Client *Version
	Server *Version
}

func versionRun(cmd *cobra.Command, args []string) {
	vd := VersionResponse{
		Client: &Version{
			Version:    "1.0",
			APIVersion: "2.0",
			GoVersion:  runtime.Version(),
			Os:         runtime.GOOS,
			Arch:       runtime.GOARCH,
		},
	}

	config := &api.Config{
		ApiRoot: viper.GetString("api_uri"),
	}

	kClient, err := api.NewKeeperAPIClient(config)
	if err != nil {
		panic(err)
	}

	serverInfo, err := kClient.GetApiInfo()
	if err == nil {
		vd.Server = &Version{
			Version:    serverInfo.Version,
			APIVersion: serverInfo.Version,
		}
	} else {
		log.Println(err)

		vd.Server = &Version{
			Version:    "n/a",
			APIVersion: "n/a",
		}
	}

	tmpl, err := template.New("version").Parse(versionTmpl)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, vd)
	if err != nil {
		panic(err)
	}
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
