package adp

import (
	"crypto/tls"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	DEFAULT_DOMAIN = "localhost"
	DEFAULT_PORT   = 443

	DEFAULT_USER = "adpuser"
)

type ClientBuilder struct {
	*Client
}

type Client struct {
	Domain        string
	Port          int
	User          string
	Password      string
	TaskAccessKey string
	RestyClient   *resty.Client
}

func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		&Client{
			Domain:        DEFAULT_DOMAIN,
			Port:          DEFAULT_PORT,
			User:          DEFAULT_USER,
			Password:      "",
			TaskAccessKey: "",
		},
	}
}

func (b *ClientBuilder) WithDomain(domain string) *ClientBuilder {
	b.Domain = domain
	return b
}

func (b *ClientBuilder) WithPort(port int) *ClientBuilder {
	b.Port = port
	return b
}

func (b *ClientBuilder) WithUser(user string) *ClientBuilder {
	b.User = user
	return b
}

func (b *ClientBuilder) WithPassword(password string) *ClientBuilder {
	b.Password = password
	return b
}

func (b *ClientBuilder) WithTaskAccessKey(accessKey string) *ClientBuilder {
	b.TaskAccessKey = accessKey
	return b
}

func (b *ClientBuilder) Build() *Client {

	r := resty.New().
		SetBaseURL(fmt.Sprintf("https://%s:%d", b.Domain, b.Port)).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Auth-Username": b.User,
			"Auth-Password": b.Password,
		})

	if len(b.TaskAccessKey) > 0 {
		r.SetHeader("Task-Access-Key", b.TaskAccessKey)
	}

	b.RestyClient = r

	return b.Client
}
