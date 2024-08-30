package adp

type AddSubengineConfiguration struct {
	AdpProgressTaskTimeout                     int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpAddSubengineJvmOptions                  any    `json:"adp_addSubengine_jvmOptions,omitempty"`
	AdpTaskActive                              bool   `json:"adp_taskActive,omitempty"`
	AdpExecutionPersistent                     bool   `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                        bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                          bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpAddSubengineHostIdentifier              any    `json:"adp_addSubengine_hostIdentifier,omitempty"`
	AdpAddSubengineChosenEngineHostMemory      string `json:"adp_addSubengine_chosenEngineHostMemory,omitempty"`
	AdpAddSubengineHostMemoryLimitRatio        string `json:"adp_addSubengine_hostMemoryLimitRatio,omitempty"`
	AdpAddSubengineHostMemoryLimit             string `json:"adp_addSubengine_hostMemoryLimit,omitempty"`
	AdpLoggingEnabled                          bool   `json:"adp_loggingEnabled,omitempty"`
	AdpAddSubengineHeapSetting                 any    `json:"adp_addSubengine_heapSetting,omitempty"`
	AdpAddSubengineEngineName                  string `json:"adp_addSubengine_engineName,omitempty"`
	AdpAddSubengineAbortOnExistingEngine       bool   `json:"adp_addSubengine_abortOnExistingEngine,omitempty"`
	AdpTaskTimeout                             int    `json:"adp_taskTimeout,omitempty"`
	AdpAddSubengineSingleEngineTemplate        any    `json:"adp_addSubengine_singleEngineTemplate,omitempty"`
	AdpAddSubengineChosenEngineHostMemoryRatio string `json:"adp_addSubengine_chosenEngineHostMemoryRatio,omitempty"`
	AdpAddSubengineCreatedEngineIdentifier     string `json:"adp_addSubengine_createdEngineIdentifier,omitempty"`
	AdpAddSubengineMergingMetaEngineTemplate   string `json:"adp_addSubengine_mergingMetaEngineTemplate,omitempty"`
	AdpAddSubengineApplicationIdentifier       string `json:"adp_addSubengine_applicationIdentifier,omitempty"`
}

func (cfg *AddSubengineConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithAddSubengineLoggingEnabled(enabled bool) func(*AddSubengineConfiguration) {
	return func(c *AddSubengineConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithAddSubengineApplicationIdentifier(id string) func(*AddSubengineConfiguration) {
	return func(c *AddSubengineConfiguration) {
		c.AdpAddSubengineApplicationIdentifier = id
	}
}

func WithAddSubengineSingleEngineTemplate(template string) func(*AddSubengineConfiguration) {
	return func(c *AddSubengineConfiguration) {
		c.AdpAddSubengineSingleEngineTemplate = template
	}
}

func WithAddSubengineCreatedEngineIdentifier(id string) func(*AddSubengineConfiguration) {
	return func(c *AddSubengineConfiguration) {
		c.AdpAddSubengineCreatedEngineIdentifier = id
	}
}

func WithAddSubengineHostIdentifier(id string) func(*AddSubengineConfiguration) {
	return func(c *AddSubengineConfiguration) {
		c.AdpAddSubengineHostIdentifier = id
	}
}

func WithAddSubengineEngineName(name string) func(*AddSubengineConfiguration) {
	return func(c *AddSubengineConfiguration) {
		c.AdpAddSubengineEngineName = name
	}
}

func WithAddSubengineJvmOptions(options any) func(*AddSubengineConfiguration) {
	return func(c *AddSubengineConfiguration) {
		c.AdpAddSubengineJvmOptions = options
	}
}

func WithAddSubengineHeapSetting(heapSetting any) func(*AddSubengineConfiguration) {
	return func(c *AddSubengineConfiguration) {
		c.AdpAddSubengineHeapSetting = heapSetting
	}
}
