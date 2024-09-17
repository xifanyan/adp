package adp

type FieldCloneTemplate struct {
	Type  string `json:"Type,omitempty"`
	Field string `json:"Field,omitempty"`
}

type NewFieldConfiguration struct {
	Enable       string `json:"Enable,omitempty"`
	NewFieldName string `json:"New field name,omitempty"`
	NewFieldID   string `json:"New field Id,omitempty"`
	Type         string `json:"Type,omitempty"`
	FieldToClone string `json:"Field to clone,omitempty"`
}

type AddFieldsConfiguration struct {
	AdpProgressTaskTimeout      int                     `json:"adp_progressTaskTimeout,omitempty"`
	AdpAddFieldsEngineType      string                  `json:"adp_addFields_engineType,omitempty"`
	AdpAddFieldsDataModelName   string                  `json:"adp_addFields_dataModelName,omitempty"`
	AdpTaskActive               bool                    `json:"adp_taskActive,omitempty"`
	AdpAddFieldsApplicationName any                     `json:"adp_addFields_applicationName,omitempty"`
	AdpAddFieldsEngineName      any                     `json:"adp_addFields_engineName,omitempty"`
	AdpAddFieldsCloneTemplates  []FieldCloneTemplate    `json:"adp_addFields_cloneTemplates,omitempty"`
	AdpExecutionPersistent      bool                    `json:"adp_executionPersistent,omitempty"`
	AdpAddFieldsEnginePassword  any                     `json:"adp_addFields_enginePassword,omitempty"`
	AdpAbortWfOnFailure         bool                    `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory           bool                    `json:"adp_cleanUpHistory,omitempty"`
	AdpLoggingEnabled           bool                    `json:"adp_loggingEnabled,omitempty"`
	AdpAddFieldsEngineUser      any                     `json:"adp_addFields_engineUser,omitempty"`
	AdpTaskTimeout              int                     `json:"adp_taskTimeout,omitempty"`
	AdpAddFieldsFields2Clone    []NewFieldConfiguration `json:"adp_addFields_fields2Clone,omitempty"`
}

func (cfg *AddFieldsConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithAdpLoggingEnabled(enabled bool) func(*AddFieldsConfiguration) {
	return func(c *AddFieldsConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithAdpAddFieldsApplicationName(name string) func(*AddFieldsConfiguration) {
	return func(c *AddFieldsConfiguration) {
		c.AdpAddFieldsApplicationName = name
	}
}

func WithAdpAddFieldsEngineName(name string) func(*AddFieldsConfiguration) {
	return func(c *AddFieldsConfiguration) {
		c.AdpAddFieldsEngineName = name
	}
}

func WithAdpAddFieldsEngineUser(user string) func(*AddFieldsConfiguration) {
	return func(c *AddFieldsConfiguration) {
		c.AdpAddFieldsEngineUser = user
	}
}

func WithAdpAddFieldsEnginePassword(password string) func(*AddFieldsConfiguration) {
	return func(c *AddFieldsConfiguration) {
		c.AdpAddFieldsEnginePassword = password
	}
}

func WithAdpAddFieldsFields2Clone(fields []NewFieldConfiguration) func(*AddFieldsConfiguration) {
	return func(c *AddFieldsConfiguration) {
		c.AdpAddFieldsFields2Clone = fields
	}
}
