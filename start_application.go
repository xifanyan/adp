package adp

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
