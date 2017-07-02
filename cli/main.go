package cli

import (
	"github.com/nunux-keeper/keeper-cli/api"
	"github.com/spf13/viper"
)

// KeeperCLI Command Line Interface
type KeeperCLI struct {
	API  *api.Client
	JSON bool
}

// NewKeeperCLI returns new CLI instance
func NewKeeperCLI() (*KeeperCLI, error) {
	endpoint := viper.GetString("endpoint")
	c, err := api.NewAPIClient(endpoint)
	if err != nil {
		return nil, err
	}
	json := viper.GetBool("json")

	return &KeeperCLI{
		API:  c,
		JSON: json,
	}, nil
}
