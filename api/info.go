package api

import (
	"encoding/json"
	"net/http"
)

type Href struct {
	Href string `json:href`
}

type InfoResponse struct {
	Name        string `json:name`
	Description string `json:description`
	Version     string `json:version`
	APIVersion  string `json:apiVersion`
	Env         string `json:env`
	Links       struct {
		AuthRealm *Href `json:"auth-realm"`
	} `json:"_links"`
}

func (k *KeeperAPIClient) GetApiInfo() (*InfoResponse, error) {
	r, err := http.Get(k.Config.Endpoint)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var result InfoResponse
	err = json.NewDecoder(r.Body).Decode(&result)
	return &result, err
}
