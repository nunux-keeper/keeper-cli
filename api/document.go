package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"strconv"
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

func (k *Client) GetDocuments(query string, order string, size int, from int) ([]DocumentResponse, error) {
	q := make(url.Values)
	q.Add("size", strconv.Itoa(size))
	q.Add("from", strconv.Itoa(from))
	q.Add("order", order)
	if query != "" {
		q.Add("q", query)
	}

	res, err := k.Get("/v2/documents", &q)
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

func (k *Client) GetDocument(docid string) (*DocumentResponse, error) {
	res, err := k.Get("/v2/documents/"+docid, nil)
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

func (k *Client) CreateDocument(doc *DocumentResponse) (*DocumentResponse, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(doc)
	fmt.Fprintf(os.Stdout, "Posting: %s", b)

	res, err := k.Post("/v2/documents", b)
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

func (k *Client) RemoveDocument(docid string) error {
	res, err := k.Delete("/v2/documents/"+docid, nil)
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

func (k *Client) RestoreDocument(docid string) (*DocumentResponse, error) {
	res, err := k.Put("/v2/graveyard/documents/"+docid, nil)
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
