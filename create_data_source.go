package adp

type CreateDataSourceConfiguration struct {
	AdpProgressTaskTimeout                                   int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpCreateDataSourceAbortOnExistingDataSource             bool   `json:"adp_createDataSource_abortOnExistingDataSource,omitempty"`
	AdpCreateDataSourceApplicationIdentifier                 any    `json:"adp_createDataSource_applicationIdentifier,omitempty"`
	AdpTaskActive                                            bool   `json:"adp_taskActive,omitempty"`
	AdpCreateDataSourceChoosenHostNameParameter              string `json:"adp_createDataSource_choosenHostNameParameter,omitempty"`
	AdpExecutionPersistent                                   bool   `json:"adp_executionPersistent,omitempty"`
	AdpCreateDataSourceChoosenHostMemoryRatio                string `json:"adp_createDataSource_choosenHostMemoryRatio,omitempty"`
	AdpAbortWfOnFailure                                      bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpCreateDataSourceChosenHostCPULoad                     string `json:"adp_createDataSource_chosenHostCpuLoad,omitempty"`
	AdpLoggingEnabled                                        bool   `json:"adp_loggingEnabled,omitempty"`
	AdpCreateDataSourceDataSourceSystemTemplateDisplayName   string `json:"adp_createDataSource_dataSourceSystemTemplateDisplayName,omitempty"`
	AdpCreateDataSourceUsedTemplate                          string `json:"adp_createDataSource_usedTemplate,omitempty"`
	AdpCreateDataSourceHostCPULoadThreshold                  string `json:"adp_createDataSource_hostCpuLoadThreshold,omitempty"`
	AdpCreateDataSourceCreatedDataSourceNameParameter        string `json:"adp_createDataSource_createdDataSourceNameParameter,omitempty"`
	AdpCreateDataSourceRetryMaxNumberRunningCrawlers         string `json:"adp_createDataSource_retryMaxNumberRunningCrawlers,omitempty"`
	AdpCreateDataSourceChoosenHostMemory                     string `json:"adp_createDataSource_choosenHostMemory,omitempty"`
	AdpCreateDataSourceWorkspaceIdentifier                   any    `json:"adp_createDataSource_workspaceIdentifier,omitempty"`
	AdpCreateDataSourceHostIdentifier                        any    `json:"adp_createDataSource_hostIdentifier,omitempty"`
	AdpCreateDataSourceHostMemoryLimit                       string `json:"adp_createDataSource_hostMemoryLimit,omitempty"`
	AdpCreateDataSourceMaxNumberRunningCrawlers              string `json:"adp_createDataSource_maxNumberRunningCrawlers,omitempty"`
	AdpCleanUpHistory                                        bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpCreateDataSourceEngineIdentifier                      any    `json:"adp_createDataSource_engineIdentifier,omitempty"`
	AdpCreateDataSourceEngineBoxDocThreshold                 string `json:"adp_createDataSource_engineBoxDocThreshold,omitempty"`
	AdpCreateDataSourceHostMemoryLimitRatio                  string `json:"adp_createDataSource_hostMemoryLimitRatio,omitempty"`
	AdpCreateDataSourceChoosenEngineNameParameter            string `json:"adp_createDataSource_choosenEngineNameParameter,omitempty"`
	AdpCreateDataSourceHostRolesBlackList                    any    `json:"adp_createDataSource_hostRolesBlackList,omitempty"`
	AdpCreateDataSourceDataSourceIdentifier                  string `json:"adp_createDataSource_dataSourceIdentifier,omitempty"`
	AdpTaskTimeout                                           int    `json:"adp_taskTimeout,omitempty"`
	AdpCreateDataSourceCreatedDataSourceDisplaynameParameter string `json:"adp_createDataSource_createdDataSourceDisplaynameParameter,omitempty"`
	AdpCreateDataSourceDataSourceTemplate                    string `json:"adp_createDataSource_dataSourceTemplate,omitempty"`
	AdpCreateDataSourceDataSourceName                        string `json:"adp_createDataSource_dataSourceName,omitempty"`
}

func (cfg *CreateDataSourceConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithCreateDatasourceLoggingEnabled(enabled bool) func(*CreateDataSourceConfiguration) {
	return func(c *CreateDataSourceConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithCreateDatasourceDatasourceIdentifier(id string) func(*CreateDataSourceConfiguration) {
	return func(c *CreateDataSourceConfiguration) {
		c.AdpCreateDataSourceDataSourceIdentifier = id
	}
}

func WithCreateDatasourceDatasourceName(name string) func(*CreateDataSourceConfiguration) {
	return func(c *CreateDataSourceConfiguration) {
		c.AdpCreateDataSourceDataSourceName = name
	}
}

func WithCreateDatasourceDatasourceTemplate(name string) func(*CreateDataSourceConfiguration) {
	return func(c *CreateDataSourceConfiguration) {
		c.AdpCreateDataSourceDataSourceTemplate = name
	}
}

func WithCreateDatasourceEngineBoxDocThreshold(threshold string) func(*CreateDataSourceConfiguration) {
	return func(c *CreateDataSourceConfiguration) {
		c.AdpCreateDataSourceEngineBoxDocThreshold = threshold
	}
}

func WithCreateDatasourceEngineIdentifier(id string) func(*CreateDataSourceConfiguration) {
	return func(c *CreateDataSourceConfiguration) {
		c.AdpCreateDataSourceEngineIdentifier = id
	}
}

func WithCreateDatasourceHostIdentifier(id string) func(*CreateDataSourceConfiguration) {
	return func(c *CreateDataSourceConfiguration) {
		c.AdpCreateDataSourceHostIdentifier = id
	}
}
