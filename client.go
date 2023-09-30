package adp

import (
	"crypto/tls"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	Host          string
	Port          int
	User          string
	Password      string
	TaskAccessKey string
	RestyClient   *resty.Client
}

// NewClient creates a new instance of the Client struct.
//
// Returns a pointer to the newly created Client.
func NewClient() *Client {

	return &Client{
		Host:          "localhost",
		Port:          443,
		User:          "adpuser",
		Password:      "",
		TaskAccessKey: "",
	}
}

// WithHost sets the host for the client.
//
// host: the host address to set.
// Returns the updated client with the new host.
func (client *Client) WithHost(host string) *Client {
	client.Host = host
	return client
}

// WithPort sets the port of the client.
//
// port: The port number to be set.
// Returns the updated client with the new port.
func (client *Client) WithPort(port int) *Client {
	client.Port = port
	return client
}

// WithUser sets the user for the Client.
//
// user: The username to set for the Client.
// Returns the updated client with new user.
func (client *Client) WithUser(user string) *Client {
	client.User = user
	return client
}

// WithPassword sets the password for the client.
//
// password: The password to be set.
// Returns the updated client with new user password.
func (client *Client) WithPassword(password string) *Client {
	client.Password = password
	return client
}

// WithTaskAccessKey sets the access key for the client.
//
// accessKey: The access key to be set.
// Returns the updated client with the new access key.
func (client *Client) WithTaskAccessKey(accessKey string) *Client {
	client.TaskAccessKey = accessKey
	return client
}

// Assemble assembles the client with the necessary configuration.
//
// This function sets up the REST client with the base URL, TLS configuration,
// and headers. It also sets the Task-Access-Key header if provided.
// The function returns the assembled client.
func (client *Client) Assemble() *Client {

	r := resty.New().
		SetBaseURL(fmt.Sprintf("https://%s:%d", client.Host, client.Port)).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Auth-Username": client.User,
			"Auth-Password": client.Password,
		})

	if len(client.TaskAccessKey) > 0 {
		r.SetHeader("Task-Access-Key", client.TaskAccessKey)
	}

	client.RestyClient = r

	return client
}
