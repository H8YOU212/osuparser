package osunet

import "net/http"

type Client struct {
	baseURL string
	token   token

	http *http.Client

	// userService    *UsersService
	// beatmapService *BeatmapsService
}

type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiredIn   int    `json:"expires_in"`
}

type authdata struct {
	ClientID     int	`json:"client_id"`
	ClientSecret string	`json:"client_secret"`
	Granttype    string	`json:"grant_type"`
	Scope        string	`json:"scope"`
}

