package api

import (
	"fmt"
	"net/url"
)

type Client struct {
	Config *Config
}

func NewNunuxKeeperClient(endpoint string) (*Client, error) {
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

	return &Client{
		Config: config,
	}, nil
}
