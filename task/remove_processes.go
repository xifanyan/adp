package task

import (
	"github.com/rs/zerolog/log"
	"opentext.com/axcelerate/adp/client"
)

// RemoveProcessesConfiguration ...
type RemoveProcessesConfiguration struct {
	AdpProgressTaskTimeout                   int                                   `json:"adp_progressTaskTimeout,omitempty"`
	AdpLoggingEnabled                        bool                                  `json:"adp_loggingEnabled"`
	AdpRemoveProcessRemoveAssociatedStorages string                                `json:"adp_removeProcess_removeAssociatedStorages"`
	AdpRemoveProcessProcessIdentifiers       []RemoveProcessesProcessIdentifierArg `json:"adp_removeProcess_processIdentifiers"`
	AdpTaskActive                            bool                                  `json:"adp_taskActive"`
	AdpTaskTimeout                           int                                   `json:"adp_taskTimeout,omitempty"`
	AdpExecutionPersistent                   bool                                  `json:"adp_executionPersistent"`
	AdpRemoveProcessStoragesRemoved          string                                `json:"adp_removeProcess_storagesRemoved,omitempty"`
	AdpAbortWfOnFailure                      bool                                  `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                        bool                                  `json:"adp_cleanUpHistory,omitempty"`
}

// NewRemoveProcessesTaskRequest ...
func NewRemoveProcessesTaskRequest(opts ...func(*RemoveProcessesConfiguration)) *Request {
	cfg := &RemoveProcessesConfiguration{}
	for _, opt := range opts {
		opt(cfg)
	}

	log.Debug().Msgf("cfg: %+v", cfg)

	for _, opt := range opts {
		opt(cfg)
	}

	return &Request{
		TaskType:          "Remove Processes",
		TaskDescription:   "Removes processes",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Remove Processes",
	}
}

// WithRemoveProcessesLoggingEnabled ...
func WithRemoveProcessesLoggingEnabled(b bool) func(*RemoveProcessesConfiguration) {
	return func(c *RemoveProcessesConfiguration) {
		c.AdpLoggingEnabled = b
	}
}

// WithRemoveProcessesExecutionPersistent ...
func WithRemoveProcessesExecutionPersistent(b bool) func(*RemoveProcessesConfiguration) {
	return func(c *RemoveProcessesConfiguration) {
		c.AdpExecutionPersistent = b
	}
}

// WithRemoveProcessesTaskActive, need to be true to enable the operation
func WithRemoveProcessesTaskActive(b bool) func(*RemoveProcessesConfiguration) {
	return func(c *RemoveProcessesConfiguration) {
		c.AdpTaskActive = b
	}
}

// WithRemoveProcessesRemoveAssociatedStorages, default "false"
func WithRemoveProcessesRemoveAssociatedStorages(s string) func(*RemoveProcessesConfiguration) {
	return func(c *RemoveProcessesConfiguration) {
		c.AdpRemoveProcessRemoveAssociatedStorages = s
	}
}

// WithRemoveProcessesProcessIdentifiers ...
// @TaskModelParameter(name="adp_removeProcess_processIdentifiers", mandatory=true)
// @TableDescriptor(columnNames="Process identifier|Remove recursive|Preserve files|Source files location|Target location", columnTypes="String|Boolean|Boolean|String|String", separator="|")
func WithRemoveProcessesProcessIdentifiers(s string) func(*RemoveProcessesConfiguration) {
	return func(c *RemoveProcessesConfiguration) {
		c.AdpRemoveProcessProcessIdentifiers = parseRemoveProcessesProcessIdentifierArgs(s)
	}
}

type RemoveProcessesTask struct {
	*Task
}

func NewRemoveProcessesTask(client *client.Client, async bool, opts ...func(*RemoveProcessesConfiguration)) *RemoveProcessesTask {
	return &RemoveProcessesTask{
		&Task{
			client:       client,
			req:          NewRemoveProcessesTaskRequest(opts...),
			asynchronous: async,
		},
	}
}

// StringOutput ...
func (t *RemoveProcessesTask) StringOutput() (string, error) {
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

type RemoveProcessResult struct{}

// StructOutput ...
func (t *RemoveProcessesTask) StructOutput() (RemoveProcessResult, error) {
	var res RemoveProcessResult

	_, err := t.StringOutput()

	if err != nil {
		return res, err
	}

	return res, nil
}
