package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type UserResponse struct {
	Id           string `json:"id,omitempty"`
	Uid          string `json:"uid,omitempty"`
	Name         string `json:"name,omitempty"`
	Date         string `json:"date,omitempty"`
	NbDocuments  int    `json:nbDocuments",omitempty"`
	NbLabels     int    `json:nbLabels",omitempty"`
	NbSharing    int    `json:nbSharing",omitempty"`
	StorageUsage int    `json:storageUsage",omitempty"`
}

type ServerInfosResponse struct {
	NbUsers     int `json:nbUsers`
	NbDocuments int `json:nbDocuments`
}

func (k *Client) GetServerInfos() (*ServerInfosResponse, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", k.Config.Endpoint+"/v2/admin/infos", nil)
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

	var result ServerInfosResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}

func (k *Client) GetUsers() ([]UserResponse, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", k.Config.Endpoint+"/v2/admin/users", nil)
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

	var result []UserResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return result, err
}

func (k *Client) GetUser(uid string) (*UserResponse, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", k.Config.Endpoint+"/v2/admin/users/"+uid, nil)
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

	var result UserResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}
