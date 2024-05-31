package adp

type StartDataSourceConfiguration struct {
	AdpProgressTaskTimeout                                   int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpStartDataSourceEngineUser                             any    `json:"adp_startDataSource_engineUser,omitempty"`
	AdpStartDataSourceIgnoreSuspended                        string `json:"adp_startDataSource_ignoreSuspended,omitempty"`
	AdpStartDataSourceIndexedDsDoccount                      string `json:"adp_startDataSource_indexedDsDoccount,omitempty"`
	AdpStartDataSourceExpectedDsDoccount                     string `json:"adp_startDataSource_expectedDsDoccount,omitempty"`
	AdpStartDataSourceNumberDocuments                        string `json:"adp_startDataSource_numberDocuments,omitempty"`
	AdpTaskActive                                            bool   `json:"adp_taskActive,omitempty"`
	AdpStartDataSourceNumberDocumentsToCrawl                 string `json:"adp_startDataSource_numberDocumentsToCrawl,omitempty"`
	AdpStartDataSourceSelectedHost                           string `json:"adp_startDataSource_selectedHost,omitempty"`
	AdpStartDataSourceOldDsDoccount                          string `json:"adp_startDataSource_oldDsDoccount,omitempty"`
	AdpExecutionPersistent                                   bool   `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                                      bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpStartDataSourceHostIdentifiers                        any    `json:"adp_startDataSource_hostIdentifiers,omitempty"`
	AdpCleanUpHistory                                        bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpStartDataSourceDataSourceName                         string `json:"adp_startDataSource_dataSourceName,omitempty"`
	AdpLoggingEnabled                                        bool   `json:"adp_loggingEnabled,omitempty"`
	AdpStartDataSourceSynchronous                            bool   `json:"adp_startDataSource_synchronous,omitempty"`
	AdpStartDataSourceLoadBalancingEnabled                   string `json:"adp_startDataSource_loadBalancingEnabled,omitempty"`
	AdpStartDataSourceEnginePassword                         any    `json:"adp_startDataSource_enginePassword,omitempty"`
	AdpTaskTimeout                                           int    `json:"adp_taskTimeout,omitempty"`
	AdpStartDataSourceMaxNumberRunningCrawlers               string `json:"adp_startDataSource_maxNumberRunningCrawlers,omitempty"`
	AdpStartDataSourceMaxUnknownStatusPollingDurationSeconds int    `json:"adp_startDataSource_maxUnknownStatusPollingDurationSeconds,omitempty"`
	AdpStartDataSourceHostRolesBlackList                     any    `json:"adp_startDataSource_hostRolesBlackList,omitempty"`
}

func (cfg *StartDataSourceConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithStartDataSourceLoggingEnabled(enabled bool) func(*StartDataSourceConfiguration) {
	return func(c *StartDataSourceConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithStartDataSourceDataSourceName(name string) func(*StartDataSourceConfiguration) {
	return func(c *StartDataSourceConfiguration) {
		c.AdpStartDataSourceDataSourceName = name
	}
}

func WithStartDataSourceSynchronous(enabled bool) func(*StartDataSourceConfiguration) {
	return func(c *StartDataSourceConfiguration) {
		c.AdpStartDataSourceSynchronous = enabled
	}
}
