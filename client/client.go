package client

import (
	"crypto/tls"
	"net/http"
)

// Client represents HTTP client which is used call ADP Rest API.
type Client struct {
	Host          string
	User          string
	Password      string
	TaskAccessKey string
	HTTPClient    *http.Client
}

// NewClient is constructor of RestClient
func NewClient(opts ...func(*Client)) *Client {
	cfg := &http.Transport{
		MaxIdleConns:        5,
		MaxConnsPerHost:     5,
		MaxIdleConnsPerHost: 5,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // NOLINT
	}

	client := &Client{
		HTTPClient: &http.Client{Transport: cfg},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// WithHost sets host field
func WithHost(s string) func(*Client) {
	return func(c *Client) {
		c.Host = s
	}
}

// WithUser sets User field
func WithUser(s string) func(*Client) {
	return func(c *Client) {
		c.User = s
	}
}

// WithPassword sets Password field
func WithPassword(s string) func(*Client) {
	return func(c *Client) {
		c.Password = s
	}
}

// WithTaskAccessKey sets TaskAccessKey field
func WithTaskAccessKey(s string) func(*Client) {
	return func(c *Client) {
		c.TaskAccessKey = s
	}
}
