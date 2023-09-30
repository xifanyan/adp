package adp

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
)

const (
	DefaultListEntitiesWhitelist = "id,displayName,processStatus,hostId,hostName"
)

type ListEntitiesConfiguration struct {
	AdpProgressTaskTimeout                                 int      `json:"adp_progressTaskTimeout,omitempty"`
	AdpListEntitiesFile                                    string   `json:"adp_listEntities_file,omitempty"`
	AdpListEntitiesNumberOfEntities                        string   `json:"adp_listEntities_numberOfEntities,omitempty"`
	AdpListEntitiesAxcRequestTimeoutSeconds                int      `json:"adp_listEntities_axcRequestTimeoutSeconds,omitempty"`
	AdpTaskActive                                          bool     `json:"adp_taskActive,omitempty"`
	AdpListEntitiesUserHasAccess                           string   `json:"adp_listEntities_userHasAccess,omitempty"`
	AdpListEntitiesWhiteList                               string   `json:"adp_listEntities_whiteList,omitempty"`
	AdpExecutionPersistent                                 bool     `json:"adp_executionPersistent"`
	AdpAbortWfOnFailure                                    bool     `json:"adp_abortWfOnFailure,omitempty"`
	AdpListEntitiesRelatedEntity                           string   `json:"adp_listEntities_relatedEntity,omitempty"`
	AdpListEntitiesWorkspace                               string   `json:"adp_listEntities_workspace,omitempty"`
	AdpLoggingEnabled                                      bool     `json:"adp_loggingEnabled"`
	AdpListEntitiesStatus                                  string   `json:"adp_listEntities_status,omitempty"`
	AdpListEntitiesAxcServiceCoreAddress                   string   `json:"adp_listEntities_axcServiceCoreAddress,omitempty"`
	AdpListEntitiesRelatedEntityType                       string   `json:"adp_listEntities_relatedEntityType,omitempty"`
	AdpListEntitiesType                                    string   `json:"adp_listEntities_type,omitempty"`
	AdpListEntitiesHTTPSKeystoreFile                       string   `json:"adp_listEntities_httpsKeystoreFile,omitempty"`
	AdpListEntitiesHTTPSPassword                           string   `json:"adp_listEntities_httpsPassword,omitempty"`
	AdpListEntitiesAxcConnectTimeoutSeconds                int      `json:"adp_listEntities_axcConnectTimeoutSeconds,omitempty"`
	AdpListEntitiesAxcServicePassword                      string   `json:"adp_listEntities_axcServicePassword,omitempty"`
	AdpListEntitiesStartingEntity                          string   `json:"adp_listEntities_startingEntity,omitempty"`
	AdpListEntitiesOutputJSON                              string   `json:"adp_listEntities_outputJson,omitempty"`
	AdpCleanUpHistory                                      bool     `json:"adp_cleanUpHistory,omitempty"`
	AdpListEntitiesDescriptionSettingFilterValueDateFormat string   `json:"adp_listEntities_descriptionSettingFilterValueDateFormat,omitempty"`
	AdpListEntitiesDescriptionFilters                      []string `json:"adp_listEntities_descriptionFilters,omitempty"`
	AdpListEntitiesAxcServiceUser                          string   `json:"adp_listEntities_axcServiceUser,omitempty"`
	AdpListEntitiesAxcFields                               string   `json:"adp_listEntities_axcFields,omitempty"`
	AdpTaskTimeout                                         int      `json:"adp_taskTimeout,omitempty"`
	AdpListEntitiesHTTPSTrustCertificate                   string   `json:"adp_listEntities_httpsTrustCertificate,omitempty"`
	AdpListEntitiesHost                                    string   `json:"adp_listEntities_host,omitempty"`
	AdpListEntitiesOutputFilename                          string   `json:"adp_listEntities_outputFilename,omitempty"`
	AdpListEntitiesID                                      string   `json:"adp_listEntities_id,omitempty"`
	AdpListEntitiesHTTPSAllowUntrustedHosts                string   `json:"adp_listEntities_httpsAllowUntrustedHosts,omitempty"`
}

// NewListEntitiesTaskRequest creates a new Request object for task listing entities.
// It accepts optional configuration functions and returns the created Request object.
func NewListEntitiesRequest(opts ...func(*ListEntitiesConfiguration)) *Request {
	cfg := &ListEntitiesConfiguration{
		AdpListEntitiesWhiteList: DefaultListEntitiesWhitelist,
	}

	// Apply the optional configuration functions to update the configuration object
	for _, opt := range opts {
		opt(cfg)
	}

	return &Request{
		TaskType:          "List Entities",
		TaskDescription:   "List Entities (description)",
		TaskConfiguration: cfg,
		TaskDisplayName:   "List Entities (displayName)",
	}
}

// WithListEntitiesLoggingEnabled returns a function which can be used to enable or disable logging in the ListEntitiesConfiguration.
func WithListEntitiesLoggingEnabled(enabled bool) func(*ListEntitiesConfiguration) {
	// Return a function that takes a pointer to a ListEntitiesConfiguration and updates the
	// AdpLoggingEnabled field with the provided enabled value.
	return func(c *ListEntitiesConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

// WithListEntitiesExecutionPersistent returns a function to enable or disable execution persistent in the ListEntitiesConfiguration.
func WithListEntitiesExecutionPersistent(enabled bool) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpExecutionPersistent = enabled
	}
}

// WithListEntitiesWhiteList returns a function that sets the white list filed of ListEntitiesConfiguration.
func WithListEntitiesWhiteList(whiteList string) func(*ListEntitiesConfiguration) {
	// Return a closure function that sets the white list value in the configuration.
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesWhiteList = whiteList
	}
}

// WithListEntitiesRelatedEntity returns a function that sets related entity field of ListEntitiesConfiguration.
func WithListEntitiesRelatedEntity(relatedEntity string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesRelatedEntity = relatedEntity
	}
}

// WithListEntitiesType returns a function that sets the type field of ListEntitiesConfiguration.
func WithListEntitiesType(typ string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesType = typ
	}
}

// WithListEntitiesID returns a function that sets ID in ListEntitiesConfiguration.
func WithListEntitiesID(id string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesID = id
	}
}

// WithListEntitiesStatus returns a function that sets the status field of ListEntitiesConfiguration.
func WithListEntitiesStatus(status string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesStatus = status
	}
}

// WithListEntitiesWorkspace returns a function that sets the workspace field of ListEntitiesConfiguration.
func WithListEntitiesWorkspace(workspace string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesWorkspace = workspace
	}
}

type ListEntitiesTask struct {
	*Task
}

func NewListEntitiesTask(client *Client, req *Request, executionMode ExecutionMode) *ListEntitiesTask {
	return &ListEntitiesTask{
		Task: NewTask().WithClient(client).WithRequest(req).WithExecutionMode(executionMode),
	}
}

type ListEntitiesExecutionMetaData struct {
	AdpEntitiesOutputFileName string          `json:"adp_entities_output_file_name"`
	AdpEntitiesJSONOutput     json.RawMessage `json:"adp_entities_json_output"`
}

type ListEntitiesResult []Entity

type Entity struct {
	ID            string `json:"id"`
	DisplayName   string `json:"displayName,omitempty"`
	ProcessStatus string `json:"processStatus,omitempty"`
	HostName      string `json:"hostName,omitempty"`
	HostID        string `json:"hostId,omitempty"`
}

// GetResponseAsString returns the string representation of the response received from running the ListEntitiesTask.
//
// It returns a string and an error. The string represents the output of the task, while the error indicates any issues encountered during execution.
func (t *ListEntitiesTask) GetResultAsString() (string, error) {
	taskResp, err := t.Run()
	if err != nil {
		return "", err
	}

	meta := &ListEntitiesExecutionMetaData{}
	if err := json.Unmarshal(taskResp.ExecutionMetaData, meta); err != nil {
		return "", fmt.Errorf("parse ListEntitiesExecutionMetaData %w", err)
	}

	output := string(meta.AdpEntitiesJSONOutput)
	unquoteJSONOutput(&output)

	return output, nil
}

// GetResponseAsData returns the ListEntitiesResult and an error.
//
// It makes a rest call to ADP service and get the response as string, then unmarshal it to ListEntitiesResult.
func (t *ListEntitiesTask) GetResultAsData() (ListEntitiesResult, error) {
	var res ListEntitiesResult

	output, err := t.GetResultAsString()
	log.Debug().Msgf("output: %s", output)

	if err != nil {
		return res, err
	}

	if err := json.Unmarshal([]byte(output), &res); err != nil {
		return res, fmt.Errorf("parse ListEntitiesJSONOutput %w", err)
	}

	return res, nil
}
