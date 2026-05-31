package osunet

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const (
	uri = "https://osu.ppy.sh/"
)

func NewClient() *Client {
	c := &Client{
		baseURL: uri,
		http:    http.DefaultClient,
	}
	return c
}

func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) NewRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token.AccessToken)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *Client) auth() {
	var authdata authdata

	authdata.ClientID, _ = strconv.Atoi(os.Getenv("clientid"))
	authdata.ClientSecret = os.Getenv("clientsecret")
	authdata.Granttype = "client_credentials"
	authdata.Scope = "public"

	data, err := json.Marshal(authdata)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", c.baseURL+"oauth/token", bytes.NewReader(data))
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body))
		return
	}

	_ = json.Unmarshal(body, &c.token)
}
