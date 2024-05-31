package adp

import (
	"encoding/json"
	"time"
)

type ReadServiceAlertsConfiguration struct {
	AdpProgressTaskTimeout               int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpTaskActive                        bool   `json:"adp_taskActive,omitempty"`
	AdpExecutionPersistent               bool   `json:"adp_executionPersistent,omitempty"`
	AdpReadServiceAlertsDate             any    `json:"adp_readServiceAlerts_date,omitempty"`
	AdpAbortWfOnFailure                  bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpReadServiceAlertsBlacklist        string `json:"adp_readServiceAlerts_blacklist,omitempty"`
	AdpCleanUpHistory                    bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpLoggingEnabled                    bool   `json:"adp_loggingEnabled,omitempty"`
	AdpReadServiceAlertsLastDate         string `json:"adp_readServiceAlerts_lastDate,omitempty"`
	AdpTaskTimeout                       int    `json:"adp_taskTimeout,omitempty"`
	AdpReadServiceAlertsListOfProperties string `json:"adp_readServiceAlerts_listOfProperties,omitempty"`
	AdpReadServiceAlertsMaximum          string `json:"adp_readServiceAlerts_maximum,omitempty"`
	AdpReadServiceAlertsOutputJSON       string `json:"adp_readServiceAlerts_outputJson,omitempty"`
}

func (c *ReadServiceAlertsConfiguration) enforcePersistentExecution() {
	c.AdpExecutionPersistent = true
}

func WithReadServiceAlertsLoggingEnabled(enabled bool) func(*ReadServiceAlertsConfiguration) {
	return func(c *ReadServiceAlertsConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func (c *ReadServiceAlertsConfiguration) WithReadServiceAlertsLastDate(date string) func(*ReadServiceAlertsConfiguration) {
	return func(c *ReadServiceAlertsConfiguration) {
		c.AdpReadServiceAlertsLastDate = date
	}
}

/*
WithReadServiceAlertsBlacklist sets the blacklist for reading service alerts.

Parameters:
blacklist string:

	The blacklist contains a JSON string with a list of regEx strings.
	Alerts with matching text will be be moved from the result to automatically exclude known alerts.

Returns:

	func(*ReadServiceAlertsConfiguration)
*/
func (c *ReadServiceAlertsConfiguration) WithReadServiceAlertsBlacklist(blacklist string) func(*ReadServiceAlertsConfiguration) {
	return func(c *ReadServiceAlertsConfiguration) {
		c.AdpReadServiceAlertsBlacklist = blacklist
	}
}

type ReadServiceAlertsExecutionMetaData struct {
	AdpReadServiceAlertsJSONOutput json.RawMessage `json:"adp_readServiceAlerts_json_output"`
}

type ReadServiceAlertsResult []ServiceAlert

type ServiceAlert struct {
	Message                   string    `json:"message"`
	ID                        string    `json:"id"`
	Identification            string    `json:"identification"`
	AlternativeIdentification string    `json:"alternativeIdentification"`
	HostName                  string    `json:"hostName"`
	Applicationsa             []string  `json:"applications"`
	Severity                  string    `json:"severity"`
	ReportOn                  time.Time `json:"reportOn"`
}
