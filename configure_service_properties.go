package adp

type ConfigureServicePropertiesConfiguration struct {
	AdpProgressTaskTimeout                       int                 `json:"adp_progressTaskTimeout,omitempty"`
	AdpLoggingEnabled                            bool                `json:"adp_loggingEnabled,omitempty"`
	AdpConfigureServicePropertiesIgnoreSuspended string              `json:"adp_configureServiceProperties_ignoreSuspended,omitempty"`
	AdpTaskActive                                bool                `json:"adp_taskActive,omitempty"`
	AdpTaskTimeout                               int                 `json:"adp_taskTimeout,omitempty"`
	AdpExecutionPersistent                       bool                `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                          bool                `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                            bool                `json:"adp_cleanUpHistory,omitempty"`
	AdpConfigureServicePropertiesActionsOnEntity []ActionOnEntityArg `json:"adp_configureServiceProperties_actionsOnEntity,omitempty"`
}

func (c *ConfigureServicePropertiesConfiguration) enforcePersistentExecution() {
	c.AdpExecutionPersistent = true
}

func WithConfigureServicePropertiesLoggingEnabled(enabled bool) func(*ConfigureServicePropertiesConfiguration) {
	return func(c *ConfigureServicePropertiesConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithConfigureServicePropertiesActionsOnEntity(actionsOnEntity []ActionOnEntityArg) func(*ConfigureServicePropertiesConfiguration) {
	return func(c *ConfigureServicePropertiesConfiguration) {
		c.AdpConfigureServicePropertiesActionsOnEntity = actionsOnEntity
	}
}
