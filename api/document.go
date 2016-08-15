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

type DocumentResponse struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Content     string `json:"content,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Date        string `json:"date,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Ghost       bool   `json:"ghost,omitempty"`
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

func (k *KeeperAPIClient) CreateDocument(doc *DocumentResponse) (*DocumentResponse, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(doc)
	fmt.Fprintf(os.Stdout, "Posting: %s", b)

	client := &http.Client{}
	req, err := http.NewRequest("POST", k.Config.Endpoint+"/v2/document", b)
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

	var result DocumentResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}

func (k *KeeperAPIClient) RemoveDocument(docid string) error {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", k.Config.Endpoint+"/v2/document/"+docid, nil)
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

func (k *KeeperAPIClient) RestoreDocument(docid string) (*DocumentResponse, error) {
	accessToken, err := GetAccessToken(k.Config)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", k.Config.Endpoint+"/v2/document/"+docid+"/restore", nil)
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
