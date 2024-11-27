package adp

import "encoding/json"

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

type ManageTaxonomyTaxonomyInput struct {
	TaxonomyName string                        `json:"taxonomyName,omitempty"`
	Categories   []ManageTaxonomyCategoryInput `json:"categories"`
}

type ManageTaxonomyCategoryInput struct {
	ID         string            `json:"id"`
	Mode       string            `json:"mode,omitempty"`
	Name       string            `json:"name,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

type ManageTaxonomyExecutionMetaData struct {
	AdpManageTaxonomyJSONOutput json.RawMessage `json:"adp_manage_taxonomy_json_output"`
}

type ManageTaxonomyResult []ManageTaxonomyTaxonomyItem

type ManageTaxonomyTaxonomyItem struct {
	ErrorMessage       string                     `json:"errorMessage"`
	ExecutionPerEngine map[string]EngineExecution `json:"executionPerEngine"`
	TaxonomyID         string                     `json:"taxonomyId"`
	ExecutionStatus    string                     `json:"executionStatus"`
	Action             ManageTaxonomyAction       `json:"action"`
}

type EngineExecution struct {
	ErrorMessage    string `json:"errorMessage"`
	ExecutionStatus string `json:"executionStatus"`
}

type ManageTaxonomyAction struct {
	ActionType string `json:"actionType"`
	ManageTaxonomyCategoryInput
}
