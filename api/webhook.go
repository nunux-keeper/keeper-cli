package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
)

// WebhookResponse Webhook object response
type WebhookResponse struct {
	ID               string   `json:"id,omitempty"`
	Owner            string   `json:"owner,omitempty"`
	URL              string   `json:"url,omitempty"`
	Secret           string   `json:"secret,omitempty"`
	Active           bool     `json:"active,omitempty"`
	Events           []string `json:"events,omitempty"`
	Labels           []string `json:"labels,omitempty"`
	CreationDate     string   `json:"cdate,omitempty"`
	ModificationDate string   `json:"mdate,omitempty"`
}

// WebhooksResponse Webhook list response
type WebhooksResponse struct {
	Webhooks []WebhookResponse
}

// GetWebhooks Get webhook list
func (k *Client) GetWebhooks() ([]WebhookResponse, error) {
	res, err := k.Get("/v2/webhooks", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result WebhooksResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return result.Webhooks, err
}

// GetWebhook Get a webhook
func (k *Client) GetWebhook(id string) (*WebhookResponse, error) {
	res, err := k.Get("/v2/webhooks/"+id, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result WebhookResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}

// CreateWebhook Create a webhook
func (k *Client) CreateWebhook(webhook *WebhookResponse) (*WebhookResponse, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(webhook)

	res, err := k.Post("/v2/webhooks/", nil, b)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result WebhookResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}

// UpdateWebhook Update a webhook
func (k *Client) UpdateWebhook(id string, webhook *WebhookResponse) (*WebhookResponse, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(webhook)

	res, err := k.Put("/v2/webhooks/"+id, nil, b)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result WebhookResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}

// RemoveWebhook Delete a webhook
func (k *Client) RemoveWebhook(id string) error {
	res, err := k.Delete("/v2/webhooks/"+id, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return errors.New(res.Status)
	}
	return nil
}
