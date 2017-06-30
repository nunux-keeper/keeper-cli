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

type Request struct {
	Method string
	Url    string
	Query  *url.Values
	Body   io.Reader
	Form   *url.Values
}

// Do HTTP request
func (k *Client) Do(r *Request) (*http.Response, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(r.Method, k.Config.Endpoint+r.Url, r.Body)
	if err != nil {
		return nil, err
	}

	if accessToken != "" {
		req.Header.Add("Authorization", "Bearer "+accessToken)
	}

	if r.Method == "POST" && r.Form != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else if r.Method == "POST" || r.Method == "PUT" {
		req.Header.Set("Content-Type", "application/json")
	}

	if r.Query != nil {
		req.URL.RawQuery = r.Query.Encode()
	}

	return client.Do(req)
}

// Get Trigger HTTP GET
func (k *Client) Get(url string, query *url.Values) (*http.Response, error) {
	req := &Request{Method: "GET", Url: url, Query: query}
	return k.Do(req)
}

// Delete Trigger HTTP DELETE
func (k *Client) Delete(url string, query *url.Values) (*http.Response, error) {
	req := &Request{Method: "DELETE", Url: url, Query: query}
	return k.Do(req)
}

// Post Trigger HTTP POST
func (k *Client) Post(url string, query *url.Values, body io.Reader) (*http.Response, error) {
	req := &Request{Method: "POST", Url: url, Query: query, Body: body}
	return k.Do(req)
}

// PostForm Trigger HTTP POST form
func (k *Client) PostForm(url string, query *url.Values, form *url.Values) (*http.Response, error) {
	req := &Request{Method: "POST", Url: url, Query: query, Form: form}
	return k.Do(req)
}

// Put Trigger HTTP PUT
func (k *Client) Put(url string, query *url.Values, body io.Reader) (*http.Response, error) {
	req := &Request{Method: "PUT", Url: url, Query: query, Body: body}
	return k.Do(req)
}

// NewAPIClient returns new API client instance
func NewAPIClient(endpoint string) (*Client, error) {
	_, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("Invalid server endpoint %s: %s", endpoint, err)
	}

	config := &Config{
		Endpoint: endpoint,
	}

	config.Credentials, err = LoadTokenInfos()
	if err != nil {
		return nil, err
	}

	return &Client{
		Config: config,
	}, nil
}
