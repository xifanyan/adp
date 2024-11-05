package adp

type ManageTaggersConfiguration struct {
	AdpMantagsJSONInstall           string `json:"adp_mantags_jsonInstall,omitempty"`
	AdpProgressTaskTimeout          int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpMantagsEngineUser            any    `json:"adp_mantags_engineUser,omitempty"`
	AdpMantagsJSONFile              any    `json:"adp_mantags_jsonFile,omitempty"`
	AdpMantagsJSONUninstall         string `json:"adp_mantags_jsonUninstall,omitempty"`
	AdpMantagsOutputJSON            string `json:"adp_mantags_outputJson,omitempty"`
	AdpTaskActive                   bool   `json:"adp_taskActive,omitempty"`
	AdpMantagsApplicationIdentifier string `json:"adp_mantags_applicationIdentifier,omitempty"`
	AdpMantagsApplicationType       string `json:"adp_mantags_applicationType,omitempty"`
	AdpExecutionPersistent          bool   `json:"adp_executionPersistent,omitempty"`
	AdpMantagsEngineIdentifier      string `json:"adp_mantags_engineIdentifier,omitempty"`
	AdpAbortWfOnFailure             bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory               bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpLoggingEnabled               bool   `json:"adp_loggingEnabled,omitempty"`
	AdpMantagsEnginePassword        any    `json:"adp_mantags_enginePassword,omitempty"`
	AdpMantagsEngineType            string `json:"adp_mantags_engineType,omitempty"`
	AdpTaskTimeout                  int    `json:"adp_taskTimeout,omitempty"`
	AdpMantagsOutputFilename        string `json:"adp_mantags_outputFilename,omitempty"`
	AdpMantagsAllowOverwrite        string `json:"adp_mantags_allowOverwrite,omitempty"`
	AdpMantagsWait4Completion       string `json:"adp_mantags_wait4Completion,omitempty"`
}

type TaggerInfo struct {
	ID             string `json:"id"`
	Description    string `json:"description,omitempty"`
	TermTaxonomy   string `json:"termTaxonomy,omitempty"`
	TypeTaxonomy   string `json:"typeTaxonomy,omitempty"`
	GlobalSearchID string `json:"globalSearchId,omitempty"`
}

func (cfg *ManageTaggersConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithManageTaggersLoggingEnabled(enabled bool) func(*ManageTaggersConfiguration) {
	return func(c *ManageTaggersConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithAdpManTagsJSONInstall(path string) func(*ManageTaggersConfiguration) {
	return func(c *ManageTaggersConfiguration) {
		c.AdpMantagsJSONInstall = path
	}
}

func WithAdpManTagsApplicationIdentifier(id string) func(*ManageTaggersConfiguration) {
	return func(c *ManageTaggersConfiguration) {
		c.AdpMantagsApplicationIdentifier = id
	}
}

func WithAdpManTagsApplicationType(appType string) func(*ManageTaggersConfiguration) {
	return func(c *ManageTaggersConfiguration) {
		c.AdpMantagsApplicationType = appType
	}
}

func WithAdpManTagsEngineIdentifier(id string) func(*ManageTaggersConfiguration) {
	return func(c *ManageTaggersConfiguration) {
		c.AdpMantagsEngineIdentifier = id
	}
}

func WithAdpManTagsEngineType(engineType string) func(*ManageTaggersConfiguration) {
	return func(c *ManageTaggersConfiguration) {
		c.AdpMantagsEngineType = engineType
	}
}

func WithAdpManTagsWait4Completion(wait string) func(*ManageTaggersConfiguration) {
	return func(c *ManageTaggersConfiguration) {
		c.AdpMantagsWait4Completion = wait
	}
}
