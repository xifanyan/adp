package adp

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
func WithStopProcessesProcessIdentifiers(processIdentifiers []StopProcessesProcessIdentifierArg) func(*StopProcessesConfiguration) {
	return func(c *StopProcessesConfiguration) {
		c.AdpStopProcessProcessIdentifiers = processIdentifiers
	}
}
