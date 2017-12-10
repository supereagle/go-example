package main

import (
	"net/http"
	"io"
)

// customizedClient represents the customized based on standard HTTP client.
type customizedClient struct {
	basicAuth *BasicAuth
	client    *http.Client
}

// basicAuth represents the basic authentication for request.
type BasicAuth struct {
	Username string
	Password string
}

// NewCustomizedClient news a customized client.
func NewCustomizedClient(username, password string) *customizedClient {
	return &customizedClient{
		basicAuth: &BasicAuth{
			Username: username,
			Password: password,
		},
		client: http.DefaultClient,
	}
}

func (c *customizedClient) do(method string, url string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	// Set the auth for the request if needed.
	if c.basicAuth != nil {
		req.SetBasicAuth(c.basicAuth.Username, c.basicAuth.Password)
	}

	return c.client.Do(req)
}
