package controllers

import (
	"net/http"
)

const (
	BaseApiUrl = "https://data.sfgov.org"
	BaseGeoApiUrl = "https://maps.googleapis.com"
)

type Client struct {
	baseUrl string
	httpClient  *http.Client
}

type Option func(*Client)

func SetHTTPClient(httpClient *http.Client) Option {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

func NewClient(baseUrl string, options ...Option) *Client {
	client := Client {
		baseUrl: baseUrl,
		httpClient: &http.Client {
			// Timeout: 10 * time.Second,
		},
	}

	for i := range options {
		options[i](&client)
	}

	return &client
}
