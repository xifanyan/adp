package adp

type QueryEngineConfiguration struct {
	AdpProgressTaskTimeout                 int                       `json:"adp_progressTaskTimeout,omitempty"`
	AdpQueryEngineFieldName                string                    `json:"adp_queryEngine_fieldName,omitempty"`
	AdpQueryEngineEngineName               string                    `json:"adp_queryEngine_engineName"`
	AdpTaskActive                          bool                      `json:"adp_taskActive,omitempty"`
	AdpExecutionPersistent                 bool                      `json:"adp_executionPersistent"`
	AdpQueryEngineEngineUserPassword       string                    `json:"adp_queryEngine_engineUserPassword,omitempty"`
	AdpAbortWfOnFailure                    bool                      `json:"adp_abortWfOnFailure,omitempty"`
	AdpLoggingEnabled                      bool                      `json:"adp_loggingEnabled"`
	AdpQueryEngineEngineTaxonomies         []EngineTaxonomyArg       `json:"adp_queryEngine_engineTaxonomies,omitempty"`
	AdpQueryEngineEngineUserName           string                    `json:"adp_queryEngine_engineUserName,omitempty"`
	AdpQueryEngineEngineType               string                    `json:"adp_queryEngine_engineType,omitempty"`
	AdpQueryEngineSaveVariable             string                    `json:"adp_queryEngine_saveVariable,omitempty"`
	AdpQueryEngineCategoryToDelete         string                    `json:"adp_queryEngine_categoryToDelete,omitempty"`
	AdpQueryEngineActivateCategoryDeletion bool                      `json:"adp_queryEngine_activateCategoryDeletion,omitempty"`
	AdpQueryEngineApplicationIdentifier    string                    `json:"adp_queryEngine_applicationIdentifier"`
	AdpQueryEngineTaxonomyToDelete         string                    `json:"adp_queryEngine_taxonomyToDelete,omitempty"`
	AdpQueryEngineSuccessIfCountIs         string                    `json:"adp_queryEngine_successIfCountIs,omitempty"`
	AdpQueryEngineCategory                 string                    `json:"adp_queryEngine_category,omitempty"`
	AdpQueryEngineActivateTagging          bool                      `json:"adp_queryEngine_activateTagging,omitempty"`
	AdpQueryEngineGlobalSearchID           string                    `json:"adp_queryEngine_globalSearchId,omitempty"`
	AdpQueryEngineAggregatedValue          string                    `json:"adp_queryEngine_aggregatedValue,omitempty"`
	AdpQueryEngineAdvancedRestrictions     []AdvancedRestrictionsArg `json:"adp_queryEngine_AdvancedRestrictions,omitempty"`
	AdpQueryEngineTaxonomy                 string                    `json:"adp_queryEngine_taxonomy,omitempty"`
	AdpQueryEngineGlobalSearchJSON         string                    `json:"adp_queryEngine_globalSearchJson,omitempty"`
	AdpQueryEngineSaveCompareString        string                    `json:"adp_queryEngine_saveCompareString,omitempty"`
	AdpCleanUpHistory                      bool                      `json:"adp_cleanUpHistory,omitempty"`
	AdpQueryEngineNumberOfDocuments        string                    `json:"adp_queryEngine_numberOfDocuments,omitempty"`
	AdpQueryEngineEngineQuery              string                    `json:"adp_queryEngine_engineQuery,omitempty"`
	AdpQueryEngineMainQueryType            any                       `json:"adp_queryEngine_mainQueryType,omitempty"`
	AdpQueryEngineWaitForResult            bool                      `json:"adp_queryEngine_waitForResult,omitempty"`
	AdpQueryEngineCategoryDisplayName      string                    `json:"adp_queryEngine_categoryDisplayName,omitempty"`
	AdpQueryEngineWaitWhileCountIs         string                    `json:"adp_queryEngine_waitWhileCountIs,omitempty"`
	AdpTaskTimeout                         int                       `json:"adp_taskTimeout,omitempty"`
	AdpQueryEngineApplicationType          string                    `json:"adp_queryEngine_applicationType,omitempty"`
	AdpQueryEngineExitOnValueChanged       bool                      `json:"adp_queryEngine_exitOnValueChanged,omitempty"`
}

func (cfg *QueryEngineConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

// WithQueryEngineLoggingEnabled sets the persistent flag for logging in the QueryEngineConfiguration.
func WithQueryEngineLoggingEnabled(enabled bool) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

// WithQueryEngineExecutionPersistent sets the persistent flag for query engine execution in the QueryEngineConfiguration.
func WithQueryEngineExecutionPersistent(enabled bool) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpExecutionPersistent = enabled
	}
}

// @TaskModelParameter(name="adp_queryEngine_engineTaxonomies", mandatory=false)
// @TableDescriptor(columnNames="Taxonomy|Negation|Query", columnTypes="String|boolean|String", separator="|")
func WithQueryEngineEngineTaxonomies(engineTaxonomies []EngineTaxonomyArg) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpQueryEngineEngineTaxonomies = engineTaxonomies
	}
}

// @TaskModelParameter(name="adp_queryEngine_AdvancedRestrictions", mandatory=false)
// @TableDescriptor(columnNames="Advanced Search Field|Negation|Fuzzy Search|Advanced Expression", columnTypes="String|boolean|boolean|String", separator="|")
func WithQueryEngineAdvancedRestrictions(advancedRestrictions []AdvancedRestrictionsArg) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpQueryEngineAdvancedRestrictions = advancedRestrictions
	}
}

// ApplicationIdentifier and EngineName are mutually exclusive.
// If one is assigned, the other should be empty and can not be omitted.
func WithQueryEngineApplicationIdentifier(applicationIdentifier string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		if len(applicationIdentifier) > 0 {
			c.AdpQueryEngineApplicationIdentifier = applicationIdentifier
		}
	}
}

// WithQueryEngineEngineName sets the query engine engine name in the QueryEngineConfiguration.
//
// It takes a string parameter and returns a function that modifies the QueryEngineConfiguration object.
func WithQueryEngineEngineName(engineName string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		if len(engineName) > 0 {
			c.AdpQueryEngineEngineName = engineName
		}
	}
}

// WithQueryEngineEngineQuery updates the QueryEngineConfiguration with the specified QueryEngineEngineQuery.
//
// It takes a string as a parameter and returns a function that accepts a pointer to QueryEngineConfiguration.
func WithQueryEngineEngineQuery(query string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpQueryEngineEngineQuery = query
	}
}

// WithQueryEngineEngineUserName sets the QueryEngineConfiguration's AdpQueryEngineEngineUserName field to the given string.
//
// It takes a string as input and returns a function that modifies a QueryEngineConfiguration object.
func WithQueryEngineEngineUserName(user string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpQueryEngineEngineUserName = user
	}
}

// WithQueryEngineEngineUserPassword sets the QueryEngineEngineUserPassword field in the QueryEngineConfiguration struct.
//
// It takes a string parameter and returns a function that modifies the QueryEngineConfiguration object.
func WithQueryEngineEngineUserPassword(password string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpQueryEngineEngineUserPassword = password
	}
}

type QueryEngineResult struct {
	Count string `json:"adp_query_engine_documents_count"`
	Size  string `json:"adp_query_engine_aggregated_value"`
}
