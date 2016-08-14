package api

import (
	"fmt"
	"net/url"
)

type KeeperAPIClient struct {
	Config *Config
}

func NewKeeperAPIClient(endpoint string) (*KeeperAPIClient, error) {
	_, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("Invalid server endpoint %s: %s", endpoint, err)
	}

	creds, err := LoadTokenInfos()
	if err != nil {
		return nil, err
	}

	config := &Config{
		Endpoint:    endpoint,
		Credentials: creds,
	}

	return &KeeperAPIClient{
		Config: config,
	}, nil
}
