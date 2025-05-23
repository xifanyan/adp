package adp

import "encoding/json"

type GlobalSearchesConfiguration struct {
	AdpGlobalSearchesOutputFilename                string `json:"adp_globalSearches_outputFilename,omitempty"`
	AdpProgressTaskTimeout                         int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpTaskActive                                  bool   `json:"adp_taskActive,omitempty"`
	AdpGlobalSearchesCreateUpdateGlobalSearches    string `json:"adp_globalSearches_createUpdateGlobalSearches,omitempty"`
	AdpExecutionPersistent                         bool   `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                            bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                              bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpLoggingEnabled                              bool   `json:"adp_loggingEnabled,omitempty"`
	AdpGlobalSearchesRegex                         string `json:"adp_globalSearches_regex,omitempty"`
	AdpGlobalSearchesOutputJSON                    string `json:"adp_globalSearches_outputJson,omitempty"`
	AdpTaskTimeout                                 int    `json:"adp_taskTimeout,omitempty"`
	AdpGlobalSearchesJSONFile                      string `json:"adp_globalSearches_jsonFile,omitempty"`
	AdpGlobalSearchesGlobalSearchesToDelete        string `json:"adp_globalSearches_globalSearchesToDelete,omitempty"`
	AdpGlobalSearchesUpdateCollisionResolutionMode string `json:"adp_globalSearches_updateCollisionResolutionMode,omitempty"`
}

func (c *GlobalSearchesConfiguration) enforcePersistentExecution() {
	c.AdpExecutionPersistent = true
}

func WithGlobalSearchesLoggingEnabled() func(*GlobalSearchesConfiguration) {
	return func(c *GlobalSearchesConfiguration) {
		c.AdpLoggingEnabled = true
	}
}

func WithGlobalSearchesRegex(regex string) func(*GlobalSearchesConfiguration) {
	return func(c *GlobalSearchesConfiguration) {
		c.AdpGlobalSearchesRegex = regex
	}
}

func WithGlobalSearchesCreateUpdateGlobalSearches(gs string) func(*GlobalSearchesConfiguration) {
	return func(c *GlobalSearchesConfiguration) {
		c.AdpGlobalSearchesCreateUpdateGlobalSearches = gs
	}
}

func WithGlobalSearchesGlobalSearchesToDelete(gs string) func(*GlobalSearchesConfiguration) {
	return func(c *GlobalSearchesConfiguration) {
		c.AdpGlobalSearchesGlobalSearchesToDelete = gs
	}
}

// Mode: "error", "skip", "override"
func WithGlobalSearchesUpdateCollisionResolutionMode(mode string) func(*GlobalSearchesConfiguration) {
	return func(c *GlobalSearchesConfiguration) {
		c.AdpGlobalSearchesUpdateCollisionResolutionMode = mode
	}
}

type GlobalSearchesExecutionMetaData struct {
	AdpGlobalSearchesOutputFileName string          `json:"adp_globalSearches_output_file_name,omitempty"`
	AdpGlobalSearchesJSONOutput     json.RawMessage `json:"adp_globalSearches_json_output"`
}

type ActiveQueryPart struct {
	Type  string `json:"type,omitempty"`
	Query string `json:"query,omitempty"`
	Valid bool   `json:"valid,omitempty"`
	Term  string `json:"term,omitempty"`
}

type QueryBundle struct {
	SearchTypeName      string            `json:"searchTypeName,omitempty"`
	ActiveQueryParts    []ActiveQueryPart `json:"activeQueryParts,omitempty"`
	DisabledQueryParts  []interface{}     `json:"disabledQueryParts,omitempty"`
	AdvancedSearchIndex interface{}       `json:"advancedSearchIndex,omitempty"`
}

type GlobalSearch struct {
	ID               string `json:"id"`
	DisplayName      string `json:"displayName,omitempty"`
	QueryBundle      `json:"queryBundle,omitempty"`
	Workspace        string              `json:"workspace,omitempty"`
	SearchParameters map[string][]string `json:"searchParameters,omitempty"`
	Description      string              `json:"description,omitempty"`
}

type GlobalSearchDefinition struct {
	ID          string   `json:"id"`
	DisplayName string   `json:"displayName,omitempty"`
	Queries     []string `json:"queries,omitempty"`
}
