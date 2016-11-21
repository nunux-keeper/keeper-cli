package util

import (
	"github.com/spf13/viper"

	"github.com/nunux-keeper/keeper-cli/api"
)

type Factory struct {
	// Returns a client for accessing Nunux Keeper resources or an error.
	Client func() (*api.Client, error)
}

// NewFactory creates a factory with the default Keeper resources defined
func NewFactory() *Factory {
	return &Factory{
		Client: func() (*api.Client, error) {
			endpoint := viper.GetString("endpoint")
			return api.NewNunuxKeeperClient(endpoint)
		},
	}
}
