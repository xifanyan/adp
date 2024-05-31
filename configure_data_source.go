package adp

type ConfigureDataSourceConfiguration struct {
	AdpProgressTaskTimeout                                    int                  `json:"adp_progressTaskTimeout,omitempty"`
	AdpConfiguredDataSourceConfiguredDataSourceNameParameter  string               `json:"adp_configuredDataSource_configuredDataSourceNameParameter,omitempty"`
	AdpTaskActive                                             bool                 `json:"adp_taskActive,omitempty"`
	AdpConfigureDataSourceConfigurationDifferences            string               `json:"adp_configureDataSource_configurationDifferences,omitempty"`
	AdpExecutionPersistent                                    bool                 `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                                       bool                 `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                                         bool                 `json:"adp_cleanUpHistory,omitempty"`
	AdpLoggingEnabled                                         bool                 `json:"adp_loggingEnabled,omitempty"`
	AdpConfigureDataSourceMetaDataMappingToConfigTables       []ConfigTableMapsArg `json:"adp_configureDataSource_metaDataMappingToConfigTables,omitempty"`
	AdpTaskTimeout                                            int                  `json:"adp_taskTimeout,omitempty"`
	AdpConfigureDataSourceJavaHeapSize                        string               `json:"adp_configureDataSource_javaHeapSize,omitempty"`
	AdpConfigureDataSourceDataSourceNames                     string               `json:"adp_configureDataSource_dataSourceNames"`
	AdpConfigureDataSourceMetaDataMappingToSimpleConfigParams []any                `json:"adp_configureDataSource_metaDataMappingToSimpleConfigParams,omitempty"`
}

func (cfg *ConfigureDataSourceConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithConfigureDataSourceLoggingEnabled(enabled bool) func(*ConfigureDataSourceConfiguration) {
	return func(c *ConfigureDataSourceConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithConfigureDataSourceNames(names string) func(*ConfigureDataSourceConfiguration) {
	return func(c *ConfigureDataSourceConfiguration) {
		c.AdpConfigureDataSourceDataSourceNames = names
	}
}

func WithConfigureDataSourceMetaDataMappingToConfigTables(configs []ConfigTableMapsArg) func(*ConfigureDataSourceConfiguration) {
	return func(c *ConfigureDataSourceConfiguration) {
		c.AdpConfigureDataSourceMetaDataMappingToConfigTables = configs
	}
}
