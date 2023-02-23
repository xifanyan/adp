package task

import (
	"opentext.com/axcelerate/adp/client"
)

// NewStopProcessesTaskRequest ...
func NewStopProcessesTaskRequest(opts ...func(*StopProcessesConfiguration)) *Request {

	cfg := &StopProcessesConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &Request{
		TaskType:          "Stop Processes",
		TaskDescription:   "Stops Processes",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Stop Processes",
	}
}

// StopProcessesConfiguration ...
type StopProcessesConfiguration struct {
	AdpProgressTaskTimeout           int                    `json:"adp_progressTaskTimeout,omitempty"`
	AdpLoggingEnabled                bool                   `json:"adp_loggingEnabled,omitempty"`
	AdpStopProcessProcessIdentifiers []ProcessIdentifierArg `json:"adp_stopProcess_processIdentifiers"`
	AdpTaskActive                    bool                   `json:"adp_taskActive,omitempty"`
	AdpTaskTimeout                   int                    `json:"adp_taskTimeout,omitempty"`
	AdpExecutionPersistent           bool                   `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure              bool                   `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                bool                   `json:"adp_cleanUpHistory,omitempty"`
}

type StopProcessesTask struct {
	Task
}

func NewStopProcessesTask(client *client.Client, opts ...func(*StopProcessesConfiguration)) *StopProcessesTask {
	return &StopProcessesTask{
		Task{
			client:       client,
			req:          NewStopProcessesTaskRequest(opts...),
			asynchronous: false,
		},
	}
}

// WithListEntitiesLoggingEnabled ...
func WithStopProcessesLoggingEnabled(b bool) func(*StopProcessesConfiguration) {
	return func(c *StopProcessesConfiguration) {
		c.AdpLoggingEnabled = b
	}
}

// WithListEntitiesExecutionPersistent ...
func WithStopProcessesExecutionPersistent(b bool) func(*StopProcessesConfiguration) {
	return func(c *StopProcessesConfiguration) {
		c.AdpExecutionPersistent = b
	}
}

// WithStopProcessProcessProcessIdentifiers ...
// @TaskModelParameter(name="adp_stopProcess_processIdentifiers", mandatory=true)
// @TableDescriptor(columnNames="Process identifier|Stop recursive", columnTypes="String|Boolean", separator="|")
func WithStopProcessProcessProcessIdentifiers(s string) func(*StopProcessesConfiguration) {
	return func(c *StopProcessesConfiguration) {
		c.AdpStopProcessProcessIdentifiers = parseProcessIdentifiersArg(s)
	}
}

func (t *StopProcessesTask) StringOutput() (string, error) {
	taskResp, err := t.Execute()
	if err != nil {
		return "", err
	}

	output := string(taskResp.ExecutionMetaData)

	return output, nil
}
