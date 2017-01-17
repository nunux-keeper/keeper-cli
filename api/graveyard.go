package api

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func (k *Client) GetGraveyard() ([]DocumentResponse, error) {
	res, err := k.Get("/v2/graveyard/documents", nil)
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

func (k *Client) EmptyGraveyard() error {
	res, err := k.Delete("/v2/graveyard/documents", nil)
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

func (k *Client) DestroyDocument(docid string) error {
	res, err := k.Delete("/v2/graveyard/documents/"+docid, nil)
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
