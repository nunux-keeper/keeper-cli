package api

import (
	"errors"
	"strings"
)

func (k *Client) Login(username string, password string) (*TokenInfos, error) {
	if username = strings.TrimSpace(username); username == "" {
		return nil, errors.New("Username not specified.")
	}
	info, err := k.GetApiInfo()
	if err != nil {
		return nil, err
	}

	realm, err := GetAuthRealm(info.Links.AuthRealm.Href)
	if err != nil {
		return nil, err
	}

	creds := &Credentials{
		Username: username,
		Password: password,
	}

	return GetOfflineToken(realm.TokenService, creds)
}
