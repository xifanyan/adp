package adp

import (
	"encoding/json"
)

// PingProjectConfiguration defines the configuration options for the Ping Project task in ADP.
type PingProjectConfiguration struct {
	AdpProgressTaskTimeout       int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpLoggingEnabled            bool   `json:"adp_loggingEnabled,omitempty"`
	AdpPingProjectPassword       string `json:"adp_pingProject_Password,omitempty"`
	AdpTaskActive                bool   `json:"adp_taskActive,omitempty"`
	AdpTaskTimeout               int    `json:"adp_taskTimeout,omitempty"`
	AdpExecutionPersistent       bool   `json:"adp_executionPersistent,omitempty"`
	AdpPingProjectUser           string `json:"adp_pingProject_User,omitempty"`
	AdpAbortWfOnFailure          bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpPingProjectIdentifierType string `json:"adp_pingProject_IdentifierType,omitempty"`
	AdpPingProjectJSONOutput     string `json:"adp_pingProject_JsonOutput,omitempty"`
	AdpCleanUpHistory            bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpPingProjectIdentifiers    string `json:"adp_pingProject_Identifiers,omitempty"`
}

func (cfg *PingProjectConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithPingProjectLoggingEnabled(enabled bool) func(*PingProjectConfiguration) {
	return func(c *PingProjectConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithPingProjectExecutionPersistent(enabled bool) func(*PingProjectConfiguration) {
	return func(c *PingProjectConfiguration) {
		c.AdpExecutionPersistent = enabled
	}
}

func WithPingProjectIdentifiers(identifiers string) func(*PingProjectConfiguration) {
	return func(c *PingProjectConfiguration) {
		// internally it will be split by mIdentifiers.split("\\s*\\,\\s*")), so we do not have to handle it.
		c.AdpPingProjectIdentifiers = identifiers
	}
}

type PingProjectExecutionMetaData struct {
	PingProjectResult json.RawMessage `json:"ping_project_result"`
}
