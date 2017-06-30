package version

import (
	"fmt"
	"runtime"
	"text/template"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/version"
)

const versionTmpl = `Client:
 Version:      {{.Client.Version}}
 API version:  {{.Client.APIVersion}}
 Go version:   {{.Client.GoVersion}}
 OS/Arch:      {{.Client.Os}}/{{.Client.Arch}}

Server:
 Version:      {{.Server.Version}}
 API version:  {{.Server.APIVersion}}
`

type genericVersion struct {
	Version    string
	APIVersion string
	GoVersion  string
	Os         string
	Arch       string
}

type versionResponse struct {
	Client *genericVersion
	Server *genericVersion
}

func NewCommand(kCli *cli.KeeperCLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the client and server version information",
		RunE: func(cc *cobra.Command, args []string) error {
			return runVersion(kCli, cc)
		},
	}
	return cmd
}

func runVersion(kCli *cli.KeeperCLI, cmd *cobra.Command) error {
	vd := versionResponse{
		Client: &genericVersion{
			Version:    version.App,
			APIVersion: version.Api,
			GoVersion:  runtime.Version(),
			Os:         runtime.GOOS,
			Arch:       runtime.GOARCH,
		},
	}

	resp, err := kCli.APIClient.GetApiInfo()
	if err == nil {
		vd.Server = &genericVersion{
			Version:    resp.Version,
			APIVersion: resp.APIVersion,
		}
	} else {
		fmt.Fprintf(*kCli.Out, "Error: %v", err)
		vd.Server = &genericVersion{
			Version:    "n/a",
			APIVersion: "n/a",
		}
	}

	tmpl, err := template.New("version").Parse(versionTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(*kCli.Out, vd)
	return err
}
