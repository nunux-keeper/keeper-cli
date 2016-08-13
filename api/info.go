package api

import (
	"encoding/json"
	"net/http"
)

type InfoResponse struct {
	Name        string `json:name`
	Description string `json:description`
	Version     string `json:version`
	Env         string `json:env`
}

func (k *KeeperAPIClient) GetApiInfo() (*InfoResponse, error) {
	r, err := http.Get(k.Config.ApiRoot)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var result InfoResponse
	err = json.NewDecoder(r.Body).Decode(&result)
	return &result, err
}
