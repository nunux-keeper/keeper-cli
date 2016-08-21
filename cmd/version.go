package cmd

import (
	"fmt"
	"io"
	"runtime"
	"text/template"

	"github.com/spf13/cobra"

	cmdutil "github.com/ncarlier/keeper-cli/cmd/util"
	"github.com/ncarlier/keeper-cli/version"
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

func NewCmdVersion(f *cmdutil.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the client and server version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVersion(f, out, cmd)
		},
	}
	return cmd
}

func runVersion(f *cmdutil.Factory, out io.Writer, cmd *cobra.Command) error {
	vd := VersionResponse{
		Client: &Version{
			Version:    version.App,
			APIVersion: version.Api,
			GoVersion:  runtime.Version(),
			Os:         runtime.GOOS,
			Arch:       runtime.GOARCH,
		},
	}

	c, err := f.Client()
	if err != nil {
		return err
	}

	serverInfo, err := c.GetApiInfo()
	if err == nil {
		vd.Server = &Version{
			Version:    serverInfo.Version,
			APIVersion: serverInfo.APIVersion,
		}
	} else {
		fmt.Fprintf(out, "Error: %v", err)
		vd.Server = &Version{
			Version:    "n/a",
			APIVersion: "n/a",
		}
	}

	tmpl, err := template.New("version").Parse(versionTmpl)
	if err != nil {
		return err
	}
	err = tmpl.Execute(out, vd)
	return err
}
