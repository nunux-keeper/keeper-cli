package api

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type UserResponse struct {
	Id           string `json:"id"`
	Uid          string `json:"uid"`
	Name         string `json:"name"`
	Date         string `json:"date"`
	NbDocuments  int    `json:"nbDocuments"`
	NbLabels     int    `json:"nbLabels"`
	NbSharing    int    `json:"nbSharing"`
	StorageUsage int    `json:"storageUsage"`
}

type ServerInfosResponse struct {
	NbUsers     int `json:nbUsers`
	NbDocuments int `json:nbDocuments`
}

type JobsInfosResponse struct {
	InactiveCount int `json:"inactiveCount"`
	CompleteCount int `json:"completeCount"`
	ActiveCount   int `json:"activeCount"`
	FailedCount   int `json:"failedCount"`
	WorkTime      int `json:"workTime"`
}

type JobRequest struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type JobResponse struct {
	Id        string      `json:"id"`
	Type      string      `json:"type"`
	Data      interface{} `json:"data"`
	Priority  int         `json:"priority"`
	Progress  string      `json:"progress"`
	State     string      `json:"state"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	Duration  string      `json:"duration"`
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
	res, err := k.Get("/v2/admin/worker/stats", nil)
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

func (k *Client) CreateJob(job *JobRequest) (*JobResponse, error) {
	res, err := k.Post("/v2/admin/worker/job", &job, nil)
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
	res, err := k.Get("/v2/admin/worker/job/"+id, nil)
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
