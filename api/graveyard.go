package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

func (k *Client) GetGraveyard() ([]DocumentResponse, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", k.Config.Endpoint+"/v2/graveyard/documents", nil)
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

	var result DocumentsResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return result.Documents, err
}

func (k *Client) EmptyGraveyard() error {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", k.Config.Endpoint+"/v2/graveyard/documents", nil)
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

func (k *Client) DestroyDocument(docid string) error {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", k.Config.Endpoint+"/v2/graveyard/documents/"+docid, nil)
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
