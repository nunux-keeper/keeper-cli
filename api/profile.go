package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type ProfileResponse struct {
	Admin bool   `json:admin`
	Date  string `json:date`
	Hash  string `json:hash`
	Name  string `json:name`
	Uid   string `json:uid`
}

func (k *KeeperAPIClient) GetProfile() (*ProfileResponse, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", k.Config.Endpoint+"/v2/profile", nil)
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

	var result ProfileResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}
