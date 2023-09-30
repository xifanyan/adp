package adp

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

// ComputeCountsConfiguration ...
type ComputeCountsConfiguration struct {
	AdpProgressTaskTimeout                int                 `json:"adp_progressTaskTimeout,omitempty"`
	AdpComputeCountsAdvancedField         any                 `json:"adp_computeCounts_advancedField,omitempty"`
	AdpComputeCountsEngineUserPassword    string              `json:"adp_computeCounts_engineUserPassword,omitempty"`
	AdpTaskActive                         bool                `json:"adp_taskActive,omitempty"`
	AdpComputeCountsEngineGlobalSearch    string              `json:"adp_computeCounts_engineGlobalSearch,omitempty"`
	AdpExecutionPersistent                bool                `json:"adp_executionPersistent"`
	AdpComputeCountsEngineType            string              `json:"adp_computeCounts_engineType,omitempty"`
	AdpComputeCountsEngineUserName        string              `json:"adp_computeCounts_engineUserName,omitempty"`
	AdpAbortWfOnFailure                   bool                `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                     bool                `json:"adp_cleanUpHistory,omitempty"`
	AdpComputeCountsQueryParts            []string            `json:"adp_computeCounts_queryParts,omitempty"`
	AdpLoggingEnabled                     bool                `json:"adp_loggingEnabled"`
	AdpComputeCountsTypes                 string              `json:"adp_computeCounts_types,omitempty"`
	AdpComputeCountsResultName            string              `json:"adp_computeCounts_resultName,omitempty"`
	AdpComputeCountsEngineName            string              `json:"adp_computeCounts_engineName"`
	AdpTaskTimeout                        int                 `json:"adp_taskTimeout,omitempty"`
	AdpComputeCountsApplicationIdentifier string              `json:"adp_computeCounts_applicationIdentifier"`
	AdpComputeCountsEngineQuery           string              `json:"adp_computeCounts_engineQuery,omitempty"`
	AdpComputeCountsApplicationType       string              `json:"adp_computeCounts_applicationType,omitempty"`
	AdpComputeCountsEngineTaxonomies      []EngineTaxonomyArg `json:"adp_computeCounts_engineTaxonomies,omitempty"`
	AdpComputeCountsMainQueryType         any                 `json:"adp_computeCounts_mainQueryType,omitempty"`
}

func (cfg *ComputeCountsConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

// NewComputeCountsTaskRequest creates a new Request object for the ComputeCountsTask.
func NewComputeCountsTaskRequest(opts ...func(*ComputeCountsConfiguration)) *Request {

	cfg := &ComputeCountsConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	log.Debug().Msgf("cfg: %+v", cfg)

	return &Request{
		TaskType:          "Compute Counts",
		TaskDescription:   "Compute Counts.",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Compute counts task",
	}
}

func WithComputeCountsLoggingEnabled(enabled bool) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithComputeCountsExecutionPersistent(enabled bool) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpExecutionPersistent = enabled
	}
}

func WithComputeCountsEngineTaxonomies(engineTaxonomies string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		taxonomies := parseEngineTaxonomiesArgs(engineTaxonomies)
		if len(taxonomies) == 0 {
			return
		}
		c.AdpComputeCountsEngineTaxonomies = taxonomies
	}
}

func WithComputeCountsAdvancedField(advancedField string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpComputeCountsAdvancedField = advancedField
	}
}

func WithComputeCountsApplicationIdentifier(applicationIdentifier string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		if len(applicationIdentifier) > 0 {
			c.AdpComputeCountsApplicationIdentifier = applicationIdentifier
		}
	}
}

func WithComputeCountsEngineName(engineName string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		if len(engineName) > 0 {
			c.AdpComputeCountsEngineName = engineName
		}
	}
}

func WithComputeCountsEngineQuery(query string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpComputeCountsEngineQuery = query
	}
}

func WithComputeCountsQueryParts(queryParts string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpComputeCountsQueryParts = parseQueryParts(queryParts)
	}
}

func WithComputeCountsEngineUserName(user string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpComputeCountsEngineUserName = user
	}
}

func WithComputeCountsEngineUserPassword(password string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpComputeCountsEngineUserPassword = password
	}
}

type ComputeCountsTask struct {
	*Task
}

func NewComputeCountsTask(client *Client, req *Request, executionMode ExecutionMode) *ComputeCountsTask {
	return &ComputeCountsTask{
		Task: NewTask().WithClient(client).WithRequest(req).WithExecutionMode(executionMode),
	}
}

// ComputeCountsExecutionMetaData
// Example:
//
// "executionMetaData":
//
//	{"adpComputeCountResult":
//	  "{
//	     \"ax_coded_review_status=Responsive\":{\"standard\":{\"count\":0,\"size\":0}},
//	     \"ax_coded_hot=Hot\":{\"standard\":{\"count\":0,\"size\":0}}
//	  }"
//	}
type ComputeCountsExecutionMetaData struct {
	AdpComputeCountsResult json.RawMessage `json:"adpComputeCountResult"`
}

type ComputeCountsStandardOutput struct {
	Standard struct {
		Count int `json:"count"`
		Size  int `json:"size"`
	} `json:"standard"`
}

type ComputeCountsResult map[string]ComputeCountsStandardOutput

func (t *ComputeCountsTask) GetResultAsString() (string, error) {
	taskResp, err := t.Run()
	if err != nil {
		return "", err
	}

	meta := &ComputeCountsExecutionMetaData{}
	if err := json.Unmarshal(taskResp.ExecutionMetaData, meta); err != nil {
		return "", fmt.Errorf("parse ComputeCountsExecutionMetaData %w", err)
	}

	output := string(meta.AdpComputeCountsResult)
	unquoteJSONOutput(&output)

	return output, nil
}

func (t *ComputeCountsTask) GetResultAsData() (ComputeCountsResult, error) {
	var res ComputeCountsResult

	output, err := t.GetResultAsString()

	if err != nil {
		return res, err
	}

	if err := json.Unmarshal([]byte(output), &res); err != nil {
		return res, fmt.Errorf("parse ComputeCountsResult %w", err)
	}

	return res, nil
}
