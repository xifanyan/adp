package adp

import (
	"encoding/json"
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

func (cfg *ListEntitiesConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
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

func WithListEntitiesUserHasAccess(userHasAccess string) func(*ListEntitiesConfiguration) {
	return func(c *ListEntitiesConfiguration) {
		c.AdpListEntitiesUserHasAccess = userHasAccess
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
