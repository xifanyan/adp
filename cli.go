package adp

type CLIConfiguration struct {
	AdpBatchScriptRedirectLogging        string            `json:"adp_batchScriptRedirectLogging,omitempty"`
	AdpProgressTaskTimeout               int               `json:"adp_progressTaskTimeout,omitempty"`
	AdpEnvVariables                      any               `json:"adp_envVariables,omitempty"`
	AdpTaskActive                        bool              `json:"adp_taskActive,omitempty"`
	AdpPathEntries                       any               `json:"adp_pathEntries,omitempty"`
	AdpBatchScriptPositiveExecutionCodes string            `json:"adp_batchScriptPositiveExecutionCodes,omitempty"`
	AdpExecutionPersistent               bool              `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                  bool              `json:"adp_abortWfOnFailure,omitempty"`
	AdpBatchScriptResultCode             string            `json:"adp_batchScriptResultCode,omitempty"`
	AdpCleanUpHistory                    bool              `json:"adp_cleanUpHistory,omitempty"`
	AdpBatchScriptResultLogPath          string            `json:"adp_batchScriptResultLogPath,omitempty"`
	AdpLoggingEnabled                    bool              `json:"adp_loggingEnabled,omitempty"`
	AdpBatchScriptFilterPasswords        bool              `json:"adp_batchScriptFilterPasswords,omitempty"`
	AdpBatchScriptLoggingDirectory       any               `json:"adp_batchScriptLoggingDirectory,omitempty"`
	AdpBatchScriptErrorLogPath           string            `json:"adp_batchScriptErrorLogPath,omitempty"`
	AdpTaskTimeout                       int               `json:"adp_taskTimeout,omitempty"`
	AdpBatchScriptPath                   any               `json:"adp_batchScriptPath,omitempty"`
	AdpWorkingDirectory                  any               `json:"adp_workingDirectory,omitempty"`
	AdpBatchScriptParameters             []CliParameterArg `json:"adp_batchScriptParameters,omitempty"`
	AdpBatchScriptJSONLogOutput          string            `json:"adp_batchScriptJsonLogOutput,omitempty"`
}

func (cfg *CLIConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithCLILoggingEnabled(enabled bool) func(*CLIConfiguration) {
	return func(c *CLIConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithCLIExecutionPersistent(enabled bool) func(*CLIConfiguration) {
	return func(c *CLIConfiguration) {
		c.AdpExecutionPersistent = enabled
	}
}

func WithCLIBatchScriptPath(path string) func(*CLIConfiguration) {
	return func(c *CLIConfiguration) {
		c.AdpBatchScriptPath = path
	}
}

func WithCLIWorkingDirectory(path string) func(*CLIConfiguration) {
	return func(c *CLIConfiguration) {
		c.AdpWorkingDirectory = path
	}
}

func WithCLIJSONLogOutput(output string) func(*CLIConfiguration) {
	return func(c *CLIConfiguration) {
		c.AdpBatchScriptJSONLogOutput = output
	}
}

func WithCLIBatchScriptParameters(params []CliParameterArg) func(*CLIConfiguration) {
	return func(c *CLIConfiguration) {
		c.AdpBatchScriptParameters = params
	}
}
