package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type LabelResponse struct {
	Id    string `json:id`
	Label string `json:label`
	Color string `json:color`
	Date  string `json:date`
	Owner string `json:owner`
	Ghost bool   `json:ghost`
}

type LabelsResponse struct {
	Labels []LabelResponse
}

func (k *KeeperAPIClient) GetLabels() ([]LabelResponse, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", k.Config.Endpoint+"/v2/label", nil)
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
