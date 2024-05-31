package adp

type ManageTaxonomyConfiguration struct {
	AdpProgressTaskTimeout           int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpManageTaxonomyEnginePassword  string `json:"adp_manageTaxonomy_enginePassword,omitempty"`
	AdpManageTaxonomyEngineUser      string `json:"adp_manageTaxonomy_engineUser,omitempty"`
	AdpManageTaxonomyApplicationName string `json:"adp_manageTaxonomy_applicationName,omitempty"`
	AdpManageTaxonomyEngineName      string `json:"adp_manageTaxonomy_engineName,omitempty"`
	AdpTaskActive                    bool   `json:"adp_taskActive,omitempty"`
	AdpExecutionPersistent           bool   `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure              bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpLoggingEnabled                bool   `json:"adp_loggingEnabled,omitempty"`
	AdpManageTaxonomyTaxonomiesJSON  any    `json:"adp_manageTaxonomy_taxonomiesJson"`
	AdpManageTaxonomyEngineType      string `json:"adp_manageTaxonomy_engineType,omitempty"`
	AdpTaskTimeout                   int    `json:"adp_taskTimeout,omitempty"`
	AdpManageTaxonomyApplicationType string `json:"adp_manageTaxonomy_applicationType,omitempty"`
	AdpManageTaxonomyOutputJSON      string `json:"adp_manageTaxonomy_outputJson,omitempty"`
}

func (cfg *ManageTaxonomyConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithManageTaxonomyLoggingEnabled(enabled bool) func(*ManageTaxonomyConfiguration) {
	return func(c *ManageTaxonomyConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}
func WithManageTaxonomyApplicationName(name string) func(*ManageTaxonomyConfiguration) {
	return func(c *ManageTaxonomyConfiguration) {
		c.AdpManageTaxonomyApplicationName = name
	}
}

func WithManageTaxonomyTaxonomiesJSON(json string) func(*ManageTaxonomyConfiguration) {
	return func(c *ManageTaxonomyConfiguration) {
		c.AdpManageTaxonomyTaxonomiesJSON = json
	}
}
