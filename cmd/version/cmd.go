package version

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/nunux-keeper/keeper-cli/cli"
	"github.com/nunux-keeper/keeper-cli/cmd/common"
	"github.com/nunux-keeper/keeper-cli/version"
)

type genericVersion struct {
	Version    string `json:version`
	APIVersion string `json:apiVersion`
	GoVersion  string `json:goVersion`
	Os         string `json:os,omitempty`
	Arch       string `json:arch,omitempty`
}

type versionResponse struct {
	Client *genericVersion
	Server *genericVersion
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the client and server version information",
		RunE: func(cc *cobra.Command, args []string) error {
			return runVersion(cc)
		},
	}
	return cmd
}

func runVersion(cmd *cobra.Command) error {
	vd := versionResponse{
		Client: &genericVersion{
			Version:    version.App,
			APIVersion: version.Api,
			GoVersion:  runtime.Version(),
			Os:         runtime.GOOS,
			Arch:       runtime.GOARCH,
		},
	}

	kli, err := cli.NewKeeperCLI()
	if err != nil {
		return err
	}

	resp, err := kli.API.GetApiInfo()
	if err == nil {
		vd.Server = &genericVersion{
			Version:    resp.Version,
			APIVersion: resp.APIVersion,
		}
	} else {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		vd.Server = &genericVersion{
			Version:    "n/a",
			APIVersion: "n/a",
		}
	}
	return common.WriteCmdResponse(vd, common.VERSION, kli.JSON)
}
