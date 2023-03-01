package task

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
	"opentext.com/axcelerate/adp/client"
)

// StartApplicationConfiguration ...
type StartApplicationConfiguration struct {
	AdpProgressTaskTimeout                   int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpLoggingEnabled                        bool   `json:"adp_loggingEnabled"`
	AdpTaskActive                            bool   `json:"adp_taskActive,omitempty"`
	AdpTaskTimeout                           int    `json:"adp_taskTimeout,omitempty"`
	AdpStartApplicationUseHTTPS              bool   `json:"adp_startApplication_useHttps,omitempty"`
	AdpStartApplicationApplicationURL        string `json:"adp_startApplication_applicationUrl,omitempty"`
	AdpExecutionPersistent                   bool   `json:"adp_executionPersistent"`
	AdpAbortWfOnFailure                      bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                        bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpStartApplicationApplicationIdentifier string `json:"adp_startApplication_applicationIdentifier"`
}

// NewStartApplicationTaskRequest ...
func NewStartApplicationTaskRequest(opts ...func(*StartApplicationConfiguration)) *Request {

	cfg := &StartApplicationConfiguration{}
	for _, opt := range opts {
		opt(cfg)
	}

	log.Debug().Msgf("cfg: %+v", cfg)

	return &Request{
		TaskType:          "Start Application",
		TaskDescription:   "Starts an application",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Start Application",
	}
}

// WithStartApplicationLoggingEnabled...
func WithStartApplicationLoggingEnabled(b bool) func(*StartApplicationConfiguration) {
	return func(c *StartApplicationConfiguration) {
		c.AdpLoggingEnabled = b
	}
}

// WithStartApplicationExecutionPersistent ...
func WithStartApplicationExecutionPersistent(b bool) func(*StartApplicationConfiguration) {
	return func(c *StartApplicationConfiguration) {
		c.AdpExecutionPersistent = b
	}
}

// WithStartApplicationApplicationIdentifier ...
func WithStartApplicationApplicationIdentifier(s string) func(*StartApplicationConfiguration) {
	return func(c *StartApplicationConfiguration) {
		c.AdpStartApplicationApplicationIdentifier = s
	}
}

// WithStartApplicationApplicationURL ...
func WithStartApplicationApplicationURL(s string) func(*StartApplicationConfiguration) {
	return func(c *StartApplicationConfiguration) {
		c.AdpStartApplicationApplicationURL = s
	}
}

// StartApplicationExecutionMetaData ...
type StartApplicationExecutionMetaData struct{}

// Output ...
func (meta *StartApplicationExecutionMetaData) Output(raw json.RawMessage) (string, error) {
	return string(raw), nil
}

type StartApplicationTask struct {
	*Task
}

func NewStartApplicationTask(client *client.Client, async bool, opts ...func(*StartApplicationConfiguration)) *StartApplicationTask {
	return &StartApplicationTask{
		&Task{
			client:       client,
			req:          NewStartApplicationTaskRequest(opts...),
			asynchronous: async,
		},
	}
}

func (t *StartApplicationTask) StringOutput() (string, error) {
	taskResp, err := t.Execute()
	if err != nil {
		return "", err
	}

	if t.asynchronous {
		return PrettyStruct(taskResp), nil
	}

	output := string(taskResp.ExecutionMetaData)

	return output, nil
}

type StartApplicationResult struct{}

func (t *StartApplicationTask) StructOutput() (StartApplicationResult, error) {
	var res StartApplicationResult

	_, err := t.StringOutput()

	return res, err
}
