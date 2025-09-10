package adp

import (
	"crypto/tls"
	"fmt"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

const (
	// EXECUTETASK uri for executeAdpTask
	EXECUTETASK = "adp/rest/api/task/executeAdpTask"
	// EXECUTETASKASYNC uri for executeAdpTaskAsync
	EXECUTETASKASYNC = "adp/rest/api/task/executeAdpTaskAsync"
	// STATUSANDPROGRESS uri for execute
	STATUSANDPROGRESS = "adp/rest/api/task/statusAndProgress"
)

type ClientBuilder struct {
	*Client
}

type Client struct {
	domain        string
	port          int
	endpoint      string
	user          string
	password      string
	taskAccessKey string
	restyClient   *resty.Client
	mu            sync.Mutex
}

func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		&Client{
			domain:        "localhost",
			port:          443,
			user:          "adpuser",
			password:      "",
			taskAccessKey: "",
		},
	}
}

func (b *ClientBuilder) WithDomain(domain string) *ClientBuilder {
	b.domain = domain
	return b
}

func (b *ClientBuilder) WithPort(port int) *ClientBuilder {
	b.port = port
	return b
}

func (b *ClientBuilder) WithUser(user string) *ClientBuilder {
	b.user = user
	return b
}

func (b *ClientBuilder) WithPassword(password string) *ClientBuilder {
	b.password = password
	return b
}

func (b *ClientBuilder) WithTaskAccessKey(accessKey string) *ClientBuilder {
	b.taskAccessKey = accessKey
	return b
}

func (b *ClientBuilder) WithEndpoint(endpoint string) *ClientBuilder {
	b.endpoint = endpoint
	return b
}

func (b *ClientBuilder) Build() *Client {
	if b.endpoint == "" {
		b.endpoint = fmt.Sprintf("https://%s:%d", b.domain, b.port)
	}

	r := resty.New().
		SetBaseURL(b.endpoint).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Auth-Username": b.user,
			"Auth-Password": b.password,
		})

	if len(b.taskAccessKey) > 0 {
		r.SetHeader("Task-Access-Key", b.taskAccessKey)
	}

	b.restyClient = r

	return b.Client
}

/*
Send sends a request and returns a response along with any errors.

	Parameters:
	- req *Request: the request to send
	Returns:
	- *Response: the response received
	- error: an error, if any
*/
func (c *Client) Send(req *Request) (*Response, error) {
	var err error

	log.Trace().Msgf("REQUEST: %+v", req)

	resp, err := c.restyClient.R().SetBody(req).SetResult(&Response{}).Put(EXECUTETASK)
	if err != nil {
		return nil, err
	}

	taskResp := resp.Result().(*Response)
	log.Trace().Msgf("RESPONSE: %+v", taskResp)

	if taskResp.ExecutionStatus != "success" {
		return nil, fmt.Errorf("ADP Task %s failure: %s", taskResp.TaskType, taskResp.ExecutionID)
	}
	log.Trace().Msgf("ExecutionMetaData: %s", string(taskResp.ExecutionMetaData))

	return taskResp, err
}

/*
SendAsync executes the ADP task asynchronously and returns the task execution id.
The persistent execution mode is enabled if not explicitly disabled in the request.

	Parameters:
	- req *Request: the request to send
	Returns:
	- string: the execution ID of the ADP task
	- error: an error, if any
*/
func (c *Client) SendAsync(req *Request) (string, error) {
	// ensure persistent execution mode is enabled
	req.TaskConfiguration.enforcePersistentExecution()

	log.Trace().Msgf("[async] REQUEST: %+v", req)

	resp, err := c.restyClient.R().SetBody(req).SetResult(&Response{}).Put(EXECUTETASKASYNC)
	if err != nil {
		return "", err
	}

	taskResp := resp.Result().(*Response)
	log.Trace().Msgf("[async] RESPONSE: %+v", taskResp)

	return taskResp.ExecutionID, err
}

/*
StatusAndProgress retrieves the status and progress of a task execution.

	Parameters:
	- executionID: exuection ID of the task
	Returns:
	- *Response: the response received
	- error: an error, if any
*/
func (c *Client) StatusAndProgress(executionID string) (*Response, error) {
	endpoint := fmt.Sprintf("%s/%s", STATUSANDPROGRESS, executionID)
	log.Debug().Msgf("[status and progress] REQUEST: %s", endpoint)

	resp, err := c.restyClient.R().SetResult(&Response{}).Get(endpoint)
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("[status and progress] RESPONSE: %s", resp.Body())

	taskResp := resp.Result().(*Response)
	if taskResp.ExecutionStatus == "success" || taskResp.ExecutionStatus == "running" {
		return taskResp, err
	}

	return nil, fmt.Errorf("status and Progress: Task %s failure: %s", taskResp.TaskType, taskResp.ExecutionID)
}

/*
ResetCredentials updates the username and password for the client and refreshes the resty client headers.

	Parameters:
	- user: new username
	- password: new password
*/
func (c *Client) ResetCredentials(user, password string) {
	if c.user == user && c.password == password {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	log.Debug().Msgf("reset credentials: %s", user)

	c.user = user
	c.password = password

	// Update the resty client headers with new credentials
	c.restyClient.SetHeaders(map[string]string{
		"Content-Type":  "application/json",
		"Auth-Username": c.user,
		"Auth-Password": c.password,
	})

	// If task access key was previously set, preserve it
	if len(c.taskAccessKey) > 0 {
		c.restyClient.SetHeader("Task-Access-Key", c.taskAccessKey)
	}
}
