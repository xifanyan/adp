package task

import (
	"encoding/json"
	"fmt"

	"opentext.com/axcelerate/adp/client"
)

// TaxonomyStatisticConfiguration ...
type TaxonomyStatisticConfiguration struct {
	AdpProgressTaskTimeout                                int                   `json:"adp_progressTaskTimeout,omitempty"`
	AdpTaxonomyStatisticOutputJSONAbsFilePath             string                `json:"adp_taxonomyStatistic_outputJsonAbsFilePath,omitempty"`
	AdpTaxonomyStatisticApplicationIdentifier             string                `json:"adp_taxonomyStatistic_applicationIdentifier"`
	AdpTaskActive                                         bool                  `json:"adp_taskActive,omitempty"`
	AdpTaxonomyStatisticAdpTaxonomyStatisticMainQueryType any                   `json:"adp_taxonomyStatistic_adp_taxonomyStatistic_mainQueryType,omitempty"`
	AdpExecutionPersistent                                bool                  `json:"adp_executionPersistent,omitempty"`
	AdpTaxonomyStatisticEngineUserName                    string                `json:"adp_taxonomyStatistic_engineUserName,omitempty"`
	AdpAbortWfOnFailure                                   bool                  `json:"adp_abortWfOnFailure,omitempty"`
	AdpTaxonomyStatisticApplicationType                   string                `json:"adp_taxonomyStatistic_applicationType,omitempty"`
	AdpTaxonomyStatisticComputeCounts                     string                `json:"adp_taxonomyStatistic_computeCounts,omitempty"`
	AdpLoggingEnabled                                     bool                  `json:"adp_loggingEnabled,omitempty"`
	AdpTaxonomyStatisticOutputJSONFilePath                string                `json:"adp_taxonomyStatistic_outputJsonFilePath,omitempty"`
	AdpTaxonomyStatisticEngineTaxonomies                  []EngineTaxonomyArg   `json:"adp_taxonomyStatistic_engineTaxonomies"`
	AdpTaxonomyStatisticEngineUserPassword                string                `json:"adp_taxonomyStatistic_engineUserPassword,omitempty"`
	AdpTaxonomyStatisticOutputXMLAbsFilePath              string                `json:"adp_taxonomyStatistic_outputXmlAbsFilePath,omitempty"`
	AdpTaxonomyStatisticEngineQuery                       string                `json:"adp_taxonomyStatistic_engineQuery,omitempty"`
	AdpTaxonomyStatisticListCategoryProperties            string                `json:"adp_taxonomyStatistic_listCategoryProperties,omitempty"`
	AdpTaxonomyStatisticOutputTaxonomies                  []OutputTaxonomiesArg `json:"adp_taxonomyStatistic_outputTaxonomies,omitempty"`
	AdpTaxonomyStatisticOutputJSON                        string                `json:"adp_taxonomyStatistic_outputJson,omitempty"`
	AdpTaxonomyStatisticEngineType                        string                `json:"adp_taxonomyStatistic_engineType,omitempty"`
	AdpCleanUpHistory                                     bool                  `json:"adp_cleanUpHistory,omitempty"`
	AdpTaxonomyStatisticOutputXMLFilePath                 string                `json:"adp_taxonomyStatistic_outputXmlFilePath,omitempty"`
	AdpTaxonomyStatisticOutputFields                      []any                 `json:"adp_taxonomyStatistic_outputFields,omitempty"`
	AdpTaxonomyStatisticEngineGlobalSearch                string                `json:"adp_taxonomyStatistic_engineGlobalSearch,omitempty"`
	AdpTaxonomyStatisticListDocuments                     string                `json:"adp_taxonomyStatistic_listDocuments,omitempty"`
	AdpTaskTimeout                                        int                   `json:"adp_taskTimeout,omitempty"`
	AdpTaxonomyStatisticEngineName                        string                `json:"adp_taxonomyStatistic_engineName"`
}

// EnableAdpLogging ...
func (c *TaxonomyStatisticConfiguration) EnableAdpLogging() {
	c.AdpLoggingEnabled = true
}

// EnableAdpExecutionPersistent ...
func (c *TaxonomyStatisticConfiguration) EnableAdpExecutionPersistent() {
	c.AdpExecutionPersistent = true
}

// NewTaxonomyStatisticTaskRequest ...
func NewTaxonomyStatisticTaskRequest(opts ...func(*TaxonomyStatisticConfiguration)) *Request {
	cfg := &TaxonomyStatisticConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &Request{
		TaskType:          "Taxonomy Statistic",
		TaskDescription:   "Retrieves category counts for a taxonomy",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Taxonomy statistic",
	}
}

// WithTaxonomyStatisticEngineTaxonomies ...
// @TaskModelParameter(name="adp_taxonomyStatistic_engineTaxonomies", mandatory=true)
// @TableDescriptor(columnNames="Taxonomy|Negation|Query", columnTypes="String|boolean|String", separator="|")
func WithTaxonomyStatisticEngineTaxonomies(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		engineTaxonomies := parseEngineTaxonomiesArgs(s)
		if len(engineTaxonomies) == 0 {
			return
		}
		c.AdpTaxonomyStatisticEngineTaxonomies = engineTaxonomies
	}
}

// WithTaxonomyStatisticApplicationIdentifier ...
// ApplicationIdentifier and EngineName are mutually exclusive.
// If one is assigned, the other should be empty and can not be omitted.
func WithTaxonomyStatisticApplicationIdentifier(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		if len(s) > 0 {
			c.AdpTaxonomyStatisticApplicationIdentifier = s
		}
	}
}

// WithTaxonomyStatisticEngineName ...
func WithTaxonomyStatisticEngineName(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		if len(s) > 0 {
			c.AdpTaxonomyStatisticEngineName = s
		}
	}
}

// WithTaxonomyStatisticComputeCounts ...
func WithTaxonomyStatisticComputeCounts(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticComputeCounts = s
	}
}

// WithTaxonomyStatisticListCategoryProperties ...
func WithTaxonomyStatisticListCategoryProperties(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticListCategoryProperties = s
	}
}

// WithTaxonomyStatisticEngineUserName ...
func WithTaxonomyStatisticEngineUserName(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticEngineUserName = s
	}
}

// WithTaxonomyStatisticEngineUserPassword ...
func WithTaxonomyStatisticEngineUserPassword(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticEngineUserPassword = s
	}
}

// WithTaxonomyStatisticOutputTaxonomies ...
// @TaskModelParameter(name="adp_taxonomyStatistic_outputTaxonomies", mandatory=true)
// @TableDescriptor(columnNames="Taxonomy|Mode|Maximum number of categories", columnTypes="String|String|integer", separator="|")
func WithTaxonomyStatisticOutputTaxonomies(s string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		outputTaxonomies := parseOutputTaxonomiesArg(s)
		if len(outputTaxonomies) == 0 {
			return
		}

		c.AdpTaxonomyStatisticOutputTaxonomies = outputTaxonomies
	}
}

// TaxonomyStatisticExecutionMetaData ...
type TaxonomyStatisticExecutionMetaData struct {
	AdpTaxonomyStatisticsJSONOutput json.RawMessage `json:"adp_taxonomy_statistics_json_output"`
}

// TaxonomyStatisticResult ...
type TaxonomyStatisticResult []struct {
	ID       string `json:"id"`
	Category []struct {
		ID          string           `json:"id"`
		DisplayName string           `json:"displayName"`
		Count       int              `json:"count,omitempty"`
		Properties  map[string][]any `json:"properties,omitempty"`
	} `json:"category"`
}

// TaxonomyStatisticJSONOutput ...
type TaxonomyStatisticJSONOutput struct {
	Date            string              `json:"-"`
	SearchParameter []map[string]string `json:"-"`
	Statistics      struct {
		Taxonomy TaxonomyStatisticResult `json:"taxonomy"`
	} `json:"statistics"`
}

type TaxonomyStatisticTask struct {
	Task
}

func NewTaxonomyStatisticTask(client *client.Client, opts ...func(*TaxonomyStatisticConfiguration)) TaxonomyStatisticTask {
	return TaxonomyStatisticTask{
		Task{
			client:       client,
			req:          NewTaxonomyStatisticTaskRequest(opts...),
			asynchronous: false,
		},
	}
}

func (t TaxonomyStatisticTask) StringOutput() (string, error) {
	taskResp, err := t.Execute()
	if err != nil {
		return "", err
	}

	meta := &TaxonomyStatisticExecutionMetaData{}
	if err := json.Unmarshal(taskResp.ExecutionMetaData, meta); err != nil {
		return "", fmt.Errorf("parse TaxonomyStatisticExecutionMetaData %w", err)
	}

	output := string(meta.AdpTaxonomyStatisticsJSONOutput)
	unquoteJSONOutput(&output)

	return output, nil
}

func (t TaxonomyStatisticTask) StructOutput() (TaxonomyStatisticResult, error) {
	var res TaxonomyStatisticResult

	output, err := t.StringOutput()

	if err != nil {
		return res, err
	}

	if err := json.Unmarshal([]byte(output), &res); err != nil {
		return res, fmt.Errorf("parse TaxonomyStatisticJSONOutput %w", err)
	}

	return res, nil
}
