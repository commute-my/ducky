package ors

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	baseUrl *url.URL
	httpCli *http.Client
}

type ClientOption func(*Client)

func WithBaseUrl(baseUrl *url.URL) ClientOption {
	return func(c *Client) {
		c.baseUrl = baseUrl
	}
}

func NewClient(opts ...ClientOption) *Client {
	cli := &Client{
		httpCli: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(cli)
	}

	return cli
}

func (c *Client) ResolveURL(path string) *url.URL {
	return &url.URL{
		Scheme: c.baseUrl.Scheme,
		Host: c.baseUrl.Host,
		Path: fmt.Sprintf("/ors/v2/%s", strings.TrimPrefix(path, "/")),
	}	
}