package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
)

type AuthRealmResponse struct {
	Realm        string `json:realm`
	TokenService string `json:"token-service"`
}

func GetAuthRealm(authRealm string) (*AuthRealmResponse, error) {
	r, err := http.Get(authRealm)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var result AuthRealmResponse
	err = json.NewDecoder(r.Body).Decode(&result)
	return &result, err
}

type Credentials struct {
	Username string
	Password string
}

func GetOfflineToken(tokenServiceUrl string, creds *Credentials) (*TokenInfos, error) {
	r, err := http.PostForm(tokenServiceUrl+"/token", url.Values{
		"client_id":  {"nunux-keeper-cli"},
		"username":   {creds.Username},
		"password":   {creds.Password},
		"grant_type": {"password"},
		"scope":      {"offline_access"},
	})
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode >= 400 {
		return nil, errors.New(r.Status)
	}
	var result TokenInfos
	err = json.NewDecoder(r.Body).Decode(&result)
	result.TokenService = tokenServiceUrl
	return &result, err
}

func GetAccessToken(config *Config) (string, error) {
	if config.Credentials == nil {
		// return "", errors.New("No credentials. Please login first.")
		return "", nil
	}
	r, err := http.PostForm(config.Credentials.TokenService+"/token", url.Values{
		"grant_type": {"refresh_token"},
		"client_id":  {"nunux-keeper-cli"},
		// "client_secret": {config.ClientSecret},
		"refresh_token": {config.Credentials.RefreshToken},
	})
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	if r.StatusCode >= 400 {
		io.Copy(os.Stderr, r.Body)
		return "", errors.New(r.Status)
	}
	var result TokenInfos
	err = json.NewDecoder(r.Body).Decode(&result)
	return result.AccessToken, err
}
