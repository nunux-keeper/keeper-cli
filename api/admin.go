package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
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

type JobsInfosResponse struct {
	InactiveCount int `json:"inactiveCount,omitempty"`
	CompleteCount int `json:"completeCount,omitempty"`
	ActiveCount   int `json:"activeCount,omitempty"`
	FailedCount   int `json:"failedCount,omitempty"`
	WorkTime      int `json:"workTime,omitempty"`
}

type JobResponse struct {
	Id        string      `json:"id,omitempty"`
	Type      string      `json:"type",omitempty`
	Data      interface{} `json:"data",omitempty`
	Priority  int         `json:"priority",omitempty`
	Progress  string      `json:"progress",omitempty`
	State     string      `json:"state",omitempty`
	CreatedAt string      `json:"created_at",omitempty`
	UpdatedAt string      `json:"updated_at",omitempty`
	Duration  string      `json:"duration",omitempty`
}

func (k *Client) GetServerInfos() (*ServerInfosResponse, error) {
	res, err := k.Get("/v2/admin/infos", nil)
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
	res, err := k.Get("/v2/admin/users", nil)
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
	res, err := k.Get("/v2/admin/users/"+uid, nil)
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

func (k *Client) GetJobsInfos() (*JobsInfosResponse, error) {
	res, err := k.Get("/v2/admin/kue/stats", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result JobsInfosResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}

func (k *Client) CreateJob(name string, params url.Values) (*JobResponse, error) {
	res, err := k.Post("/v2/admin/jobs/"+name, &params, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result JobResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}

func (k *Client) GetJob(id string) (*JobResponse, error) {
	res, err := k.Get("/v2/admin/kue/job/"+id, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		io.Copy(os.Stderr, res.Body)
		return nil, errors.New(res.Status)
	}

	var result JobResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	return &result, err
}
