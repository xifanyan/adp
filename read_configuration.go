package adp

import "encoding/json"

// ReadConfigurationConfiguration ...
type ReadConfigurationConfiguration struct {
	AdpReadConfigurationOutputJSON     string             `json:"adp_readConfiguration_outputJson,omitempty"`
	AdpProgressTaskTimeout             int                `json:"adp_progressTaskTimeout,omitempty"`
	AdpReadConfigurationConfigsToRead  []ConfigurationArg `json:"adp_readConfiguration_configsToRead,omitempty"`
	AdpTaskActive                      bool               `json:"adp_taskActive,omitempty"`
	AdpReadConfigurationOutputFilename string             `json:"adp_readConfiguration_outputFilename,omitempty"`
	AdpExecutionPersistent             bool               `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                bool               `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                  bool               `json:"adp_cleanUpHistory,omitempty"`
	AdpReadConfigurationEntityIDToRead string             `json:"adp_readConfiguration_entityIdToRead,omitempty"`
	AdpLoggingEnabled                  bool               `json:"adp_loggingEnabled,omitempty"`
	AdpReadConfigurationFile           string             `json:"adp_readConfiguration_file,omitempty"`
	AdpTaskTimeout                     int                `json:"adp_taskTimeout,omitempty"`
	AdpReadConfigurationFileFormat     string             `json:"adp_readConfiguration_fileFormat,omitempty"`
}

func (cfg *ReadConfigurationConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithReadConfigurationLoggingEnabled(enabled bool) func(*ReadConfigurationConfiguration) {
	return func(c *ReadConfigurationConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithReadConfigurationExecutionPersistent(enabled bool) func(*ReadConfigurationConfiguration) {
	return func(c *ReadConfigurationConfiguration) {
		c.AdpExecutionPersistent = enabled
	}
}

// WithReadConfigurationConfigsToRead ...
// @TableDescriptor(columnNames="Configuration ID|Field list|Name value list|Dynamic Component Names|Application type|Entity type", columnTypes="String|String|String|String|String|String", separator="|")
// @TaskModelParameter(name="configsToRead", mandatory=true, prefixed=true)
func WithReadConfigurationConfigsToRead(cfg []ConfigurationArg) func(*ReadConfigurationConfiguration) {
	return func(c *ReadConfigurationConfiguration) {
		c.AdpReadConfigurationConfigsToRead = cfg
	}
}

// WithReadConfigurationEntityIDToRead specifies entity id of which configuration file will be loaded.
func WithReadConfigurationEntityIDToRead(s string) func(*ReadConfigurationConfiguration) {
	return func(c *ReadConfigurationConfiguration) {
		c.AdpReadConfigurationEntityIDToRead = s
	}
}

// ReadConfigurationExecutionMetaData ...
type ReadConfigurationExecutionMetaData struct {
	AdpEntitiesOutputFileName string          `json:"adp_entities_output_file_name"`
	AdpEntitiesJSONOutput     json.RawMessage `json:"adp_entities_json_output"`
}
