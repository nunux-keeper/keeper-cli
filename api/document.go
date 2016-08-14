package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type DocumentResponse struct {
	Id          string `json:id`
	Title       string `json:title`
	Content     string `json:content`
	ContentType string `json:contentType`
	Origin      string `json:origin`
	Date        string `json:date`
	Owner       string `json:owner`
	Ghost       bool   `json:ghost`
}

type DocumentsResponse struct {
	Documents []DocumentResponse `json:"hits"`
}

func (k *KeeperAPIClient) GetDocuments() ([]DocumentResponse, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", k.Config.Endpoint+"/v2/document", nil)
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

func (k *KeeperAPIClient) GetDocument(docid string) (*DocumentResponse, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", k.Config.Endpoint+"/v2/document/"+docid, nil)
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

	var result DocumentResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}
