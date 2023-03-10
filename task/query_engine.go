package task

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/xifanyan/adp/client"
)

// QueryEngineConfiguration ...
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

// NewQueryEngineTaskRequest ...
func NewQueryEngineTaskRequest(opts ...func(*QueryEngineConfiguration)) *Request {

	cfg := &QueryEngineConfiguration{}
	for _, opt := range opts {
		opt(cfg)
	}

	log.Debug().Msgf("cfg: %+v", cfg)

	return &Request{
		TaskType:          "Query Engine",
		TaskDescription:   "Queries an engine",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Query Engine",
	}
}

// WithQueryEngineLoggingEnabled ...
func WithQueryEngineLoggingEnabled(b bool) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpLoggingEnabled = b
	}
}

// WithQueryEngineExecutionPersistent ...
func WithQueryEngineExecutionPersistent(b bool) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpExecutionPersistent = b
	}
}

// WithQueryEngineEngineTaxonomies ...
// @TaskModelParameter(name="adp_queryEngine_engineTaxonomies", mandatory=false)
// @TableDescriptor(columnNames="Taxonomy|Negation|Query", columnTypes="String|boolean|String", separator="|")
func WithQueryEngineEngineTaxonomies(s string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		engineTaxonomies := parseEngineTaxonomiesArgs(s)
		if len(engineTaxonomies) == 0 {
			return
		}
		c.AdpQueryEngineEngineTaxonomies = engineTaxonomies
	}
}

// WithQueryEngineAdvancedRestrictions ...
// @TaskModelParameter(name="adp_queryEngine_AdvancedRestrictions", mandatory=false)
// @TableDescriptor(columnNames="Advanced Search Field|Negation|Fuzzy Search|Advanced Expression", columnTypes="String|boolean|boolean|String", separator="|")
func WithQueryEngineAdvancedRestrictions(s string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		advancedRestrictions := parseAdvancedRestrictionsArg(s)
		if len(advancedRestrictions) == 0 {
			return
		}
		c.AdpQueryEngineAdvancedRestrictions = advancedRestrictions
	}
}

// WithQueryEngineApplicationIdentifier ...
// ApplicationIdentifier and EngineName are mutually exclusive.
// If one is assigned, the other should be empty and can not be omitted.
func WithQueryEngineApplicationIdentifier(s string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		if len(s) > 0 {
			c.AdpQueryEngineApplicationIdentifier = s
		}
	}
}

// WithQueryEngineEngineName ...
func WithQueryEngineEngineName(s string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		if len(s) > 0 {
			c.AdpQueryEngineEngineName = s
		}
	}
}

// WithQueryEngineEngineQuery ...
func WithQueryEngineEngineQuery(s string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpQueryEngineEngineQuery = s
	}
}

// WithQueryEngineEngineUserName ...
func WithQueryEngineEngineUserName(s string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpQueryEngineEngineUserName = s
	}
}

// WithQueryEngineEngineUserPassword ...
func WithQueryEngineEngineUserPassword(s string) func(*QueryEngineConfiguration) {
	return func(c *QueryEngineConfiguration) {
		c.AdpQueryEngineEngineUserPassword = s
	}
}

// QueryEngineResult ...
type QueryEngineResult struct {
	Count string `json:"adp_query_engine_documents_count"`
	Size  string `json:"adp_query_engine_aggregated_value"`
}

type QueryEngineTask struct {
	*Task
}

func NewQueryEngineTask(client *client.Client, async bool, opts ...func(*QueryEngineConfiguration)) *QueryEngineTask {
	return &QueryEngineTask{
		&Task{
			client:       client,
			req:          NewQueryEngineTaskRequest(opts...),
			asynchronous: async,
		},
	}
}

func (t *QueryEngineTask) OutputToString() (string, error) {
	taskResp, err := t.Execute()
	if err != nil {
		return "", err
	}

	output := string(taskResp.ExecutionMetaData)
	return output, nil
}

func (t *QueryEngineTask) StructOutput() (QueryEngineResult, error) {
	var res QueryEngineResult

	output, err := t.OutputToString()

	if err != nil {
		return res, err
	}

	if err := json.Unmarshal([]byte(output), &res); err != nil {
		return res, fmt.Errorf("parse QueryEngineResult %w", err)
	}

	return res, nil
}
