package adp

type ConfigureEntityConfiguration struct {
	AdpProgressTaskTimeout                                                 int                                     `json:"adp_progressTaskTimeout,omitempty"`
	AdpConfigureEntityEntityIdentifier                                     string                                  `json:"adp_configureEntity_entityIdentifier,omitempty"`
	AdpTaskActive                                                          bool                                    `json:"adp_taskActive,omitempty"`
	AdpConfigureEntityMetaDataMappingToDynamicComponentsSimpleConfigParams []DynamicComponentsSimpleConfigParamArg `json:"adp_configureEntity_metaDataMappingToDynamicComponentsSimpleConfigParams,omitempty"`
	AdpConfigureEntityMetaDataMappingToDynamicComponentsConfigStringLists  []DynamicComponentsConfigStringListArg  `json:"adp_configureEntity_metaDataMappingToDynamicComponentsConfigStringLists,omitempty"`
	AdpConfigureEntityMetaDataMappingToSimpleConfigParams                  []SimpleConfigParamArg                  `json:"adp_configureEntity_metaDataMappingToSimpleConfigParams,omitempty"`
	AdpExecutionPersistent                                                 bool                                    `json:"adp_executionPersistent,omitempty"`
	AdpConfigureEntityMetaDataMappingToConfigStringLists                   []ConfigStringListArg                   `json:"adp_configureEntity_metaDataMappingToConfigStringLists,omitempty"`
	AdpConfigureEntityConfiguredEntityIdentifier                           string                                  `json:"adp_configureEntity_configuredEntityIdentifier,omitempty"`
	AdpAbortWfOnFailure                                                    bool                                    `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                                                      bool                                    `json:"adp_cleanUpHistory,omitempty"`
	AdpConfigureEntityConfigurationDifferences                             string                                  `json:"adp_configureEntity_configurationDifferences,omitempty"`
	AdpConfigureEntityJavaHeapSize                                         string                                  `json:"adp_configureEntity_javaHeapSize,omitempty"`
	AdpConfigureEntityMetaDataMappingToDynamicComponentsConfigTables       []DynamicComponentsConfigTableArg       `json:"adp_configureEntity_metaDataMappingToDynamicComponentsConfigTables,omitempty"`
	AdpLoggingEnabled                                                      bool                                    `json:"adp_loggingEnabled,omitempty"`
	AdpConfigureEntityMetaDataMappingToConfigTables                        []ConfigTableArg                        `json:"adp_configureEntity_metaDataMappingToConfigTables,omitempty"`
	AdpConfigureEntityMetaDataMappingToDynamicComponents                   []DynamicComponentsArg                  `json:"adp_configureEntity_metaDataMappingToDynamicComponents,omitempty"`
	AdpTaskTimeout                                                         int                                     `json:"adp_taskTimeout,omitempty"`
}

func (cfg *ConfigureEntityConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithConfigureEntityLoggingEnabled(enabled bool) func(*ConfigureEntityConfiguration) {
	return func(c *ConfigureEntityConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithAdpConfigureEntityIdentifier(entityIdentifier string) func(*ConfigureEntityConfiguration) {
	return func(c *ConfigureEntityConfiguration) {
		c.AdpConfigureEntityEntityIdentifier = entityIdentifier
	}
}

func WithAdpConfigureEntityJavaHeapSize(javaHeapSize string) func(*ConfigureEntityConfiguration) {
	return func(c *ConfigureEntityConfiguration) {
		c.AdpConfigureEntityJavaHeapSize = javaHeapSize
	}
}

func WithAdpConfigurateEntityMetaDataMappingToSimpleConfigParams(simpleConfigParams []SimpleConfigParamArg) func(*ConfigureEntityConfiguration) {
	return func(c *ConfigureEntityConfiguration) {
		c.AdpConfigureEntityMetaDataMappingToSimpleConfigParams = simpleConfigParams
	}
}

func WithAdpConfigurateEntityMetaDataMappingToConfigTables(configTables []ConfigTableArg) func(*ConfigureEntityConfiguration) {
	return func(c *ConfigureEntityConfiguration) {
		c.AdpConfigureEntityMetaDataMappingToConfigTables = configTables
	}
}

func WithAdpConfigureEntityMetaDataMappingToConfigStringLists(configStringLists []ConfigStringListArg) func(*ConfigureEntityConfiguration) {
	return func(c *ConfigureEntityConfiguration) {
		c.AdpConfigureEntityMetaDataMappingToConfigStringLists = configStringLists
	}
}

func WithAdpConfigureEntityMetaDataMappingToDynamicComponents(dynamicComponentsArgs []DynamicComponentsArg) func(*ConfigureEntityConfiguration) {
	return func(c *ConfigureEntityConfiguration) {
		c.AdpConfigureEntityMetaDataMappingToDynamicComponents = dynamicComponentsArgs
	}
}

func WithAdpConfigureEntityMetaDataMappingToDynamicComponentsSimpleConfigParams(dynamicComponentsSimpleConfigParams []DynamicComponentsSimpleConfigParamArg) func(*ConfigureEntityConfiguration) {
	return func(c *ConfigureEntityConfiguration) {
		c.AdpConfigureEntityMetaDataMappingToDynamicComponentsSimpleConfigParams = dynamicComponentsSimpleConfigParams
	}
}

func WithAdpConfigureEntityMetaDataMappingToDynamicComponentsConfigTables(dynamicComponentsConfigTables []DynamicComponentsConfigTableArg) func(*ConfigureEntityConfiguration) {
	return func(c *ConfigureEntityConfiguration) {
		c.AdpConfigureEntityMetaDataMappingToDynamicComponentsConfigTables = dynamicComponentsConfigTables
	}
}

func WithAdpConfigureEntityMetaDataMappingToDynamicComponentsConfigStringLists(dynamicComponentsConfigStringLists []DynamicComponentsConfigStringListArg) func(*ConfigureEntityConfiguration) {
	return func(c *ConfigureEntityConfiguration) {
		c.AdpConfigureEntityMetaDataMappingToDynamicComponentsConfigStringLists = dynamicComponentsConfigStringLists
	}
}
