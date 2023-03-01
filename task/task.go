// Package task ...
package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
	"opentext.com/axcelerate/adp/client"
)

const (
	// EXECUTETASK uri for executeAdpTask
	EXECUTETASK = "adp/rest/api/task/executeAdpTask"
	// EXECUTETASKASYNC uri for executeAdpTaskAsync
	EXECUTETASKASYNC = "adp/rest/api/task/executeAdpTaskAsync"
	// STATUSANDPROGRESS uri for execute
	STATUSANDPROGRESS = "adp/rest/api/task/statusAndProgress"
)

type Tasker interface {
	StringOutput() (string, error)
	SetAsync()
}

type Task struct {
	client       *client.Client
	req          *Request
	asynchronous bool
}

func (t *Task) NewHttpRequest() (*http.Request, error) {

	payload, err := json.Marshal(t.req)
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("payload: %s", Prettify(string(payload)))

	req, err := http.NewRequest(http.MethodPut, t.endpoint(), bytes.NewBuffer(payload))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Auth-Username", t.client.User)
	req.Header.Set("Auth-Password", t.client.Password)

	// QueryPostgresqlDB requires Task-Access-Key
	if len(t.client.TaskAccessKey) > 0 {
		req.Header.Set("Task-Access-Key", t.client.TaskAccessKey)
	}

	return req, nil
}

func (t *Task) Execute() (*Response, error) {
	var req *http.Request
	var resp *http.Response

	var data []byte
	var err error

	if req, err = t.NewHttpRequest(); err != nil {
		return nil, err
	}

	if resp, err = t.client.HTTPClient.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if data, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	log.Debug().Msgf("raw: %s", Prettify(string(data)))

	taskResp := &Response{}
	if err = json.Unmarshal(data, taskResp); err != nil {
		return nil, err
	}

	if !t.asynchronous && taskResp.ExecutionStatus != "success" {
		return nil, fmt.Errorf("ADP Task %s failure: %s", taskResp.TaskType, taskResp.ExecutionID)
	}

	return taskResp, err
}

func (t *Task) endpoint() string {
	if t.asynchronous {
		return fmt.Sprintf("https://%s/%s", t.client.Host, EXECUTETASKASYNC)
	}

	return fmt.Sprintf("https://%s/%s", t.client.Host, EXECUTETASK)
}

func (t *Task) SetAsync() {
	t.asynchronous = true
}

// Request ...
type Request struct {
	TaskType          string `json:"taskType"`
	TaskDescription   string `json:"taskDescription"`
	TaskDisplayName   string `json:"taskDisplayName"`
	TaskConfiguration any    `json:"taskConfiguration"`
}

// Type ...
func (req *Request) Type(s string) *Request {
	req.TaskType = s
	return req
}

// DisplayName ...
func (req *Request) DisplayName(s string) *Request {
	req.TaskDisplayName = s
	return req
}

// Description ...
func (req *Request) Description(s string) *Request {
	req.TaskDescription = s
	return req
}

// JSON ...
func (req *Request) JSON() string {
	b, _ := json.MarshalIndent(req, "", "    ")
	return string(b)
}

// Response ...
type Response struct {
	ExecutionID         string          `json:"executionId"`
	TaskType            string          `json:"taskType"`
	LoggingEnabled      string          `json:"loggingEnabled"`
	ProgressMax         int             `json:"progressMax"`
	ExecutionStatus     string          `json:"executionStatus"`
	ExecutionRootDir    string          `json:"executionRootDir"`
	ContextID           string          `json:"contextId"`
	ExecutionPersistent string          `json:"executionPersistent"`
	ProgressCurrent     int             `json:"progressCurrent"`
	ProgressPercentage  float64         `json:"progressPercentage"`
	TaskDisplayName     string          `json:"taskDisplayName"`
	ExecutionMetaData   json.RawMessage `json:"executionMetaData"`
}
