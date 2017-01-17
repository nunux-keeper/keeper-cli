package api

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type ProfileResponse struct {
	Admin bool   `json:admin`
	Date  string `json:date`
	Hash  string `json:hash`
	Name  string `json:name`
	Uid   string `json:uid`
}

func (k *Client) GetProfile() (*ProfileResponse, error) {
	res, err := k.Get("/v2/profiles/current", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result ProfileResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}
