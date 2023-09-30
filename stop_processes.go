package adp

import (
	"github.com/rs/zerolog/log"
)

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

func (cfg *StopProcessesConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func NewStopProcessesTaskRequest(opts ...func(*StopProcessesConfiguration)) *Request {

	cfg := &StopProcessesConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

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

func WithStopProcessesLoggingEnabled(enabled bool) func(*StopProcessesConfiguration) {
	return func(c *StopProcessesConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithStopProcessesExecutionPersistent(enabled bool) func(*StopProcessesConfiguration) {
	return func(c *StopProcessesConfiguration) {
		c.AdpExecutionPersistent = enabled
	}
}

// WithStopProcessesProcessIdentifiers ...
// @TaskModelParameter(name="adp_stopProcess_processIdentifiers", mandatory=true)
// @TableDescriptor(columnNames="Process identifier|Stop recursive", columnTypes="String|Boolean", separator="|")
func WithStopProcessesProcessIdentifiers(processIdentifiers string) func(*StopProcessesConfiguration) {
	return func(c *StopProcessesConfiguration) {
		c.AdpStopProcessProcessIdentifiers = parseStopProcessesProcessIdentifiersArg(processIdentifiers)
	}
}

type StopProcessesTask struct {
	*Task
}

func NewStopProcessesTask(client *Client, req *Request, executionMode ExecutionMode) *StopProcessesTask {
	return &StopProcessesTask{
		Task: NewTask().WithClient(client).WithRequest(req).WithExecutionMode(executionMode),
	}
}

func (t *StopProcessesTask) GetResultAsString() (string, error) {
	taskResp, err := t.Run()
	if err != nil {
		return "", err
	}

	output := string(taskResp.ExecutionMetaData)

	return output, nil
}

type StopProcessResult struct{}

func (t *StopProcessesTask) GetResult() (StopProcessResult, error) {
	var res StopProcessResult

	_, err := t.GetResultAsString()

	return res, err
}
