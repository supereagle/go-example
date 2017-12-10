package main

import (
	"net/http"
	"net/url"
)

// NewProxyClient news a standard HTTP client with a proxy.
func NewProxyClient(username, password string) *http.Client{
	return &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				// Just set the basic auth and directly return nil, then no proxy is used.
				req.SetBasicAuth(jenkinsUsername, jenkinsPassword)
				return nil, nil
			},
		},
	}
}
