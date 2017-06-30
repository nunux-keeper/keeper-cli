package api

import "encoding/json"

type Href struct {
	Href string `json:href`
}

type ApiInfoResponse struct {
	Name        string `json:name`
	Description string `json:description`
	Version     string `json:version`
	APIVersion  string `json:apiVersion`
	Env         string `json:env`
	Links       struct {
		AuthRealm *Href `json:"auth-realm"`
	} `json:"_links"`
}

func (k *Client) GetApiInfo() (*ApiInfoResponse, error) {
	r, err := k.Get("/", nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var result ApiInfoResponse
	err = json.NewDecoder(r.Body).Decode(&result)
	return &result, err
}
