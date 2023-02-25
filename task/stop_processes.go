package task

import (
	"github.com/rs/zerolog/log"
	"opentext.com/axcelerate/adp/client"
)

// StopProcessesConfiguration ...
type StopProcessesConfiguration struct {
	AdpProgressTaskTimeout           int                                 `json:"adp_progressTaskTimeout,omitempty"`
	AdpLoggingEnabled                bool                                `json:"adp_loggingEnabled"`
	AdpStopProcessProcessIdentifiers []StopProcessesProcessIdentifierArg `json:"adp_stopProcess_processIdentifiers"`
	AdpTaskActive                    bool                                `json:"adp_taskActive,omitempty"`
	AdpTaskTimeout                   int                                 `json:"adp_taskTimeout,omitempty"`
	AdpExecutionPersistent           bool                                `json:"adp_executionPersistent"`
	AdpAbortWfOnFailure              bool                                `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                bool                                `json:"adp_cleanUpHistory,omitempty"`
}

// NewStopProcessesTaskRequest ...
func NewStopProcessesTaskRequest(opts ...func(*StopProcessesConfiguration)) *Request {

	cfg := &StopProcessesConfiguration{}
	for _, opt := range opts {
		opt(cfg)
	}

	log.Debug().Msgf("cfg: %+v", cfg)

	return &Request{
		TaskType:          "Stop Processes",
		TaskDescription:   "Stops Processes",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Stop Processes",
	}
}

// WithStopProcessProcessProcessIdentifiers ...
// @TaskModelParameter(name="adp_stopProcess_processIdentifiers", mandatory=true)
// @TableDescriptor(columnNames="Process identifier|Stop recursive", columnTypes="String|Boolean", separator="|")
func WithStopProcessProcessProcessIdentifiers(s string) func(*StopProcessesConfiguration) {
	return func(c *StopProcessesConfiguration) {
		c.AdpStopProcessProcessIdentifiers = parseStopProcessesProcessIdentifiersArg(s)
	}
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

func (t *StopProcessesTask) StringOutput() (string, error) {
	taskResp, err := t.Execute()
	if err != nil {
		return "", err
	}

	output := string(taskResp.ExecutionMetaData)

	return output, nil
}

type StopProcessResult struct{}

func (t *StopProcessesTask) StructOutput() (StopProcessResult, error) {
	var res StopProcessResult

	_, err := t.StringOutput()

	if err != nil {
		return res, err
	}

	return res, nil
}
