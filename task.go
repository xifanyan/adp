package adp

import (
	"encoding/json"
	"fmt"

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

type ExecutionMode uint

const (
	SYNCHRONOUS ExecutionMode = iota
	ASYNCHRONOUS
)

type configurator interface {
	enforcePersistentExecution()
}

type Tasker interface {
	GetResultAsString() (string, error)
}

type Task struct {
	client        *Client
	req           *Request
	executionMode ExecutionMode
}

// NewTask creates a new Task object.
//
// It returns a pointer to a Task struct.
func NewTask() *Task {
	return &Task{
		executionMode: SYNCHRONOUS,
	}
}

// WithClient sets the client for the Task.
//
// client : The client to be set.
// Returns: The Task with the updated client.
func (t *Task) WithClient(client *Client) *Task {
	t.client = client
	return t
}

// WithRequest sets the request for the Task.
//
// req: The request to be set.
// Returns: The Task with the updated request.
func (t *Task) WithRequest(req *Request) *Task {
	t.req = req
	return t
}

// WithExecutionMode sets the execution mode of the Task.
//
// executionMode : The executionMode to be set.
// Returns: The Task with the updated executionMode.
func (t *Task) WithExecutionMode(executionMode ExecutionMode) *Task {
	t.executionMode = executionMode

	// Persistent Execution required for ASYNCHRONOUS mode
	if executionMode == ASYNCHRONOUS {
		t.req.TaskConfiguration.enforcePersistentExecution()
	}

	return t
}

// Run executes the ADP task and returns the response.
//
// Returns:
// - *Response: the response of the task execution.
// - error: an error if the task execution fails.
func (t *Task) Run() (*Response, error) {
	mode := EXECUTETASK
	if t.executionMode == ASYNCHRONOUS {
		mode = EXECUTETASKASYNC
	}
	endpoint := fmt.Sprintf("/%s", mode)

	resp, err := t.client.RestyClient.R().
		SetBody(t.req).
		Put(endpoint)
	if err != nil {
		return nil, err
	}

	log.Debug().Msgf("resp: %+v", resp)

	taskResp := &Response{}
	err = json.Unmarshal(resp.Body(), taskResp)
	if err != nil {
		return nil, err
	}

	if t.executionMode == SYNCHRONOUS && taskResp.ExecutionStatus != "success" {
		return nil, fmt.Errorf("ADP Task %s failure: %s", taskResp.TaskType, taskResp.ExecutionID)
	}

	return taskResp, nil
}
