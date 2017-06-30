package cli

import (
	"io"

	"github.com/nunux-keeper/keeper-cli/api"
	"github.com/spf13/viper"
)

// KeeperCLI Command Line Interface
type KeeperCLI struct {
	APIClient *api.Client
	In        *io.Reader
	Out       *io.Writer
	Err       *io.Writer
}

// NewKeeperCLI returns new CLI instance
func NewKeeperCLI(input io.Reader, output io.Writer, errorOutput io.Writer) (*KeeperCLI, error) {
	endpoint := viper.GetString("endpoint")
	c, err := api.NewAPIClient(endpoint)
	if err != nil {
		return nil, err
	}

	return &KeeperCLI{
		APIClient: c,
		In:        &input,
		Out:       &output,
		Err:       &errorOutput,
	}, nil
}
