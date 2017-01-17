package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type LabelResponse struct {
	Id    string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Color string `json:"color,omitempty"`
	Date  string `json:"date,omitempty"`
	Owner string `json:"owner,omitempty"`
	Ghost bool   `json:"ghost,omitempty"`
}

type LabelsResponse struct {
	Labels []LabelResponse
}

func (k *Client) GetLabels() ([]LabelResponse, error) {
	res, err := k.Get("/v2/labels", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result LabelsResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return result.Labels, err
}

func (k *Client) GetLabel(id string) (*LabelResponse, error) {
	res, err := k.Get("/v2/labels/"+id, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result LabelResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}

func (k *Client) CreateLabel(label *LabelResponse) (*LabelResponse, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(label)
	fmt.Fprintf(os.Stdout, "Posting: %s", b)

	res, err := k.Post("/v2/labels/", b)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result LabelResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}

func (k *Client) RemoveLabel(id string) error {
	res, err := k.Delete("/v2/labels/"+id, nil)
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

func (k *Client) RestoreLabel(id string) (*LabelResponse, error) {
	res, err := k.Put("/v2/graveyard/labels/"+id, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result LabelResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}
