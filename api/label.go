package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", k.Config.Endpoint+"/v2/labels", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	res, err := client.Do(req)
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
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", k.Config.Endpoint+"/v2/labels/"+id, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	res, err := client.Do(req)
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
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(label)
	fmt.Fprintf(os.Stdout, "Posting: %s", b)

	client := &http.Client{}
	req, err := http.NewRequest("POST", k.Config.Endpoint+"/v2/labels", b)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
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
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", k.Config.Endpoint+"/v2/labels/"+id, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	res, err := client.Do(req)
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
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", k.Config.Endpoint+"/v2/graveyard/labels/"+id, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	res, err := client.Do(req)
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
