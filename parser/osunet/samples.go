package osunet

import (
	"context"
	"net/http"
)

func sampleGet() (*http.Request, error) {
	client := NewClient()
	client.auth()
	req, err := client.NewRequest(context.TODO(), "GET", client.baseURL, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req, err
}

func samplePost() (*http.Request, error) {
	client := NewClient()
	req, err := client.NewRequest(context.TODO(), "POST", client.baseURL, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req, err
}
