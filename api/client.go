package api

import (
	"fmt"
	"net/url"
)

type KeeperAPIClient struct {
	Config *Config
}

func NewKeeperAPIClient(c *Config) (*KeeperAPIClient, error) {
	_, err := url.Parse(c.ApiRoot)
	if err != nil {
		return nil, fmt.Errorf("invalid servers URL %s: %s", c.ApiRoot, err)
	}
	return &KeeperAPIClient{
		Config: c,
	}, nil
}
