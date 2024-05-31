package adp

import (
	"encoding/json"
)

type TaxonomyStatisticConfiguration struct {
	AdpProgressTaskTimeout                                int                   `json:"adp_progressTaskTimeout,omitempty"`
	AdpTaxonomyStatisticOutputJSONAbsFilePath             string                `json:"adp_taxonomyStatistic_outputJsonAbsFilePath,omitempty"`
	AdpTaxonomyStatisticApplicationIdentifier             string                `json:"adp_taxonomyStatistic_applicationIdentifier"`
	AdpTaskActive                                         bool                  `json:"adp_taskActive,omitempty"`
	AdpTaxonomyStatisticAdpTaxonomyStatisticMainQueryType any                   `json:"adp_taxonomyStatistic_adp_taxonomyStatistic_mainQueryType,omitempty"`
	AdpExecutionPersistent                                bool                  `json:"adp_executionPersistent"`
	AdpTaxonomyStatisticEngineUserName                    string                `json:"adp_taxonomyStatistic_engineUserName,omitempty"`
	AdpAbortWfOnFailure                                   bool                  `json:"adp_abortWfOnFailure,omitempty"`
	AdpTaxonomyStatisticApplicationType                   string                `json:"adp_taxonomyStatistic_applicationType,omitempty"`
	AdpTaxonomyStatisticComputeCounts                     string                `json:"adp_taxonomyStatistic_computeCounts,omitempty"`
	AdpLoggingEnabled                                     bool                  `json:"adp_loggingEnabled"`
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

func (cfg *TaxonomyStatisticConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithTaxonomyStatisticLoggingEnabled(enabled bool) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithTaxonomyStatisticExecutionPersistent(enabled bool) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpExecutionPersistent = enabled
	}
}

// @TaskModelParameter(name="adp_taxonomyStatistic_engineTaxonomies", mandatory=true)
// @TableDescriptor(columnNames="Taxonomy|Negation|Query", columnTypes="String|boolean|String", separator="|")
func WithTaxonomyStatisticEngineTaxonomies(engineTaxonomies []EngineTaxonomyArg) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticEngineTaxonomies = engineTaxonomies
	}
}

func WithTaxonomyStatisticApplicationIdentifier(applicationIdentifier string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		if len(applicationIdentifier) > 0 {
			c.AdpTaxonomyStatisticApplicationIdentifier = applicationIdentifier
		}
	}
}

func WithTaxonomyStatisticEngineName(engineName string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		if len(engineName) > 0 {
			c.AdpTaxonomyStatisticEngineName = engineName
		}
	}
}

// ComputeCounts is nessessary for pulling hierachical categories.
func WithTaxonomyStatisticComputeCounts(computeCounts string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticComputeCounts = computeCounts
	}
}

// ListCategoryProperties defaults to "false"
func WithTaxonomyStatisticListCategoryProperties(listCategoryProperties string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticListCategoryProperties = listCategoryProperties
	}
}

func WithTaxonomyStatisticEngineUserName(user string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticEngineUserName = user
	}
}

func WithTaxonomyStatisticEngineUserPassword(password string) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticEngineUserPassword = password
	}
}

// @TaskModelParameter(name="adp_taxonomyStatistic_outputTaxonomies", mandatory=true)
// @TableDescriptor(columnNames="Taxonomy|Mode|Maximum number of categories", columnTypes="String|String|integer", separator="|")
func WithTaxonomyStatisticOutputTaxonomies(outputTaxonomies []OutputTaxonomiesArg) func(*TaxonomyStatisticConfiguration) {
	return func(c *TaxonomyStatisticConfiguration) {
		c.AdpTaxonomyStatisticOutputTaxonomies = outputTaxonomies
	}
}

type TaxonomyStatisticExecutionMetaData struct {
	AdpTaxonomyStatisticsJSONOutput json.RawMessage `json:"adp_taxonomy_statistics_json_output"`
}

type Category struct {
	ID          string           `json:"id"`
	DisplayName string           `json:"displayName"`
	Count       int              `json:"count"`
	Properties  map[string][]any `json:"properties,omitempty"`
}

type Taxonomy struct {
	ID       string     `json:"id"`
	Category []Category `json:"category"`
}

type TaxonomyStatisticResult []Taxonomy

type TaxonomyStatisticJSONOutput struct {
	Date            string              `json:"-"`
	SearchParameter []map[string]string `json:"-"`
	Statistics      struct {
		Taxonomy []Taxonomy `json:"taxonomy"`
	} `json:"statistics"`
}
