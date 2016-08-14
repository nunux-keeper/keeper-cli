package api

type TokenInfos struct {
	TokenService     string `json:"token_service"`
	TokenType        string `json:"token_type"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

type Config struct {
	Endpoint     string
	ClientId     string
	ClientSecret string
	Credentials  *TokenInfos
}
