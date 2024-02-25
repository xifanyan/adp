package adp

import (
	"github.com/rs/zerolog/log"
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

func (cfg *StartApplicationConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func NewStartApplicationRequest(opts ...func(*StartApplicationConfiguration)) *Request {

	cfg := &StartApplicationConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

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

func WithStartApplicationLoggingEnabled(enabled bool) func(*StartApplicationConfiguration) {
	return func(c *StartApplicationConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithStartApplicationExecutionPersistent(enabled bool) func(*StartApplicationConfiguration) {
	return func(c *StartApplicationConfiguration) {
		c.AdpExecutionPersistent = enabled
	}
}

func WithStartApplicationApplicationIdentifier(applicationIdentifier string) func(*StartApplicationConfiguration) {
	return func(c *StartApplicationConfiguration) {
		c.AdpStartApplicationApplicationIdentifier = applicationIdentifier
	}
}

func WithStartApplicationApplicationURL(applicationURL string) func(*StartApplicationConfiguration) {
	return func(c *StartApplicationConfiguration) {
		c.AdpStartApplicationApplicationURL = applicationURL
	}
}

type StartApplicationTask struct {
	*Task
}

func NewStartApplicationTask(client *Client, req *Request, executionMode ExecutionMode) *StartApplicationTask {
	return &StartApplicationTask{
		Task: NewTask().WithClient(client).WithRequest(req).WithExecutionMode(executionMode),
	}
}

func (t *StartApplicationTask) GetResultAsString() (string, error) {
	taskResp, err := t.Run()
	if err != nil {
		return "", err
	}

	output := string(taskResp.ExecutionMetaData)

	return output, nil
}

type StartApplicationResult struct{}

func (t *StartApplicationTask) GetResult() (StartApplicationResult, error) {
	var res StartApplicationResult

	_, err := t.GetResultAsString()

	return res, err
}
