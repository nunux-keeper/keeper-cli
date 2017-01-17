package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Config *Config
}

func (k *Client) Do(method string, url string, query *url.Values, body io.Reader) (*http.Response, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, k.Config.Endpoint+url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	if method == "POST" || method == "PUT" {
		req.Header.Set("Content-Type", "application/json")
	}

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	return client.Do(req)
}

func (k *Client) Get(url string, query *url.Values) (*http.Response, error) {
	return k.Do("GET", url, query, nil)
}

func (k *Client) Delete(url string, query *url.Values) (*http.Response, error) {
	return k.Do("DELETE", url, query, nil)
}

func (k *Client) Post(url string, body io.Reader) (*http.Response, error) {
	return k.Do("POST", url, nil, body)
}

func (k *Client) Put(url string, body io.Reader) (*http.Response, error) {
	return k.Do("PUT", url, nil, body)
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
