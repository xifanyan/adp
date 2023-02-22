package task

import (
	"encoding/json"
	"fmt"

	"opentext.com/axcelerate/adp/client"
)

// ComputeCountsConfiguration ...
type ComputeCountsConfiguration struct {
	AdpProgressTaskTimeout                int                 `json:"adp_progressTaskTimeout,omitempty"`
	AdpComputeCountsAdvancedField         any                 `json:"adp_computeCounts_advancedField,omitempty"`
	AdpComputeCountsEngineUserPassword    string              `json:"adp_computeCounts_engineUserPassword,omitempty"`
	AdpTaskActive                         bool                `json:"adp_taskActive,omitempty"`
	AdpComputeCountsEngineGlobalSearch    string              `json:"adp_computeCounts_engineGlobalSearch,omitempty"`
	AdpExecutionPersistent                bool                `json:"adp_executionPersistent,omitempty"`
	AdpComputeCountsEngineType            string              `json:"adp_computeCounts_engineType,omitempty"`
	AdpComputeCountsEngineUserName        string              `json:"adp_computeCounts_engineUserName,omitempty"`
	AdpAbortWfOnFailure                   bool                `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                     bool                `json:"adp_cleanUpHistory,omitempty"`
	AdpComputeCountsQueryParts            []string            `json:"adp_computeCounts_queryParts,omitempty"`
	AdpLoggingEnabled                     bool                `json:"adp_loggingEnabled,omitempty"`
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

// NewComputeCountsTaskRequest ...
func NewComputeCountsTaskRequest(opts ...func(*ComputeCountsConfiguration)) *Request {

	cfg := &ComputeCountsConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &Request{
		TaskType:          "Compute Counts",
		TaskDescription:   "Compute Counts.",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Compute counts task",
	}
}

// WithComputeCountsLoggingEnabled ...
func WithComputeCountsLoggingEnabled(b bool) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpLoggingEnabled = b
	}
}

// WithComputeCountsExecutionPersistent ...
func WithCommputeCountsExecutionPersistent(b bool) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpExecutionPersistent = b
	}
}

// WithComputeCountsEngineTaxonomies ...
func WithComputeCountsEngineTaxonomies(s string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		engineTaxonomies := parseEngineTaxonomiesArgs(s)
		if len(engineTaxonomies) == 0 {
			return
		}
		c.AdpComputeCountsEngineTaxonomies = engineTaxonomies
	}
}

// WithComputeCountsAdvancedField ...
func WithComputeCountsAdvancedField(s string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpComputeCountsAdvancedField = s
	}
}

// WithComputeCountsApplicationIdentifier ...
func WithComputeCountsApplicationIdentifier(s string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		if len(s) > 0 {
			c.AdpComputeCountsApplicationIdentifier = s
		}
	}
}

// WithComputeCountsEngineName ...
func WithComputeCountsEngineName(s string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		if len(s) > 0 {
			c.AdpComputeCountsEngineName = s
		}
	}
}

// WithComputeCountsEngineQuery ...
func WithComputeCountsEngineQuery(s string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpComputeCountsEngineQuery = s
	}
}

// WithComputeCountsQueryParts ...
func WithComputeCountsQueryParts(s string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpComputeCountsQueryParts = parseQueryParts(s)
	}
}

// WithComputeCountsEngineUserName ...
func WithComputeCountsEngineUserName(s string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpComputeCountsEngineUserName = s
	}
}

// WithComputeCountsEngineUserPassword ...
func WithComputeCountsEngineUserPassword(s string) func(*ComputeCountsConfiguration) {
	return func(c *ComputeCountsConfiguration) {
		c.AdpComputeCountsEngineUserPassword = s
	}
}

// ComputeCountsExecutionMetaData ...
// Example:
//
//		"executionMetaData":
//	   {"adpComputeCountResult":
//	     "{
//	        \"ax_coded_review_status=Responsive\":{\"standard\":{\"count\":0,\"size\":0}},
//	        \"ax_coded_hot=Hot\":{\"standard\":{\"count\":0,\"size\":0}}
//	     }"
//	   }
type ComputeCountsExecutionMetaData struct {
	AdpComputeCountsResult json.RawMessage `json:"adpComputeCountResult"`
}

type ComputeCountsStandardOutput struct {
	Standard struct {
		Count int `json:"count"`
		Size  int `json:"size"`
	} `json:"standard"`
}

// ComputeCountsResult ...
type ComputeCountsResult map[string]ComputeCountsStandardOutput

type ComputeCountsTask struct {
	Task
}

func NewComputeCountsTask(client *client.Client, opts ...func(*ComputeCountsConfiguration)) ComputeCountsTask {
	return ComputeCountsTask{
		Task{
			client:       client,
			req:          NewComputeCountsTaskRequest(opts...),
			asynchronous: false,
		},
	}
}

func (t ComputeCountsTask) StringOutput() (string, error) {
	taskResp, err := t.Execute()
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

func (t ComputeCountsTask) StructOutput() (ComputeCountsResult, error) {
	var res ComputeCountsResult

	output, err := t.StringOutput()

	if err != nil {
		return res, err
	}

	if err := json.Unmarshal([]byte(output), &res); err != nil {
		return res, fmt.Errorf("parse ComputeCountsResult %w", err)
	}

	return res, nil
}
