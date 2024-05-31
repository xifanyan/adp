package adp

type CSVMergeConfiguration struct {
	AdpCsvMergeNoUniqueMatch                  string            `json:"adp_csvMerge_noUniqueMatch,omitempty"`
	AdpCsvMergeNoFlushAfterMerge              string            `json:"adp_csvMerge_noFlushAfterMerge,omitempty"`
	AdpCsvMergeForceChange                    string            `json:"adp_csvMerge_forceChange,omitempty"`
	AdpCsvMergeEngineName                     any               `json:"adp_csvMerge_engineName,omitempty"`
	AdpCsvMergeApplicationIdentifier          string            `json:"adp_csvMerge_applicationIdentifier,omitempty"`
	AdpCsvMergeMergeType                      string            `json:"adp_csvMerge_mergeType,omitempty"`
	AdpCsvMergeDisplayNameMappingErrorFile    any               `json:"adp_csvMerge_displayNameMappingErrorFile,omitempty"`
	AdpCsvMergeLockDocumentChanges            string            `json:"adp_csvMerge_lockDocumentChanges,omitempty"`
	AdpCsvMergeCsvFieldNativeLocation         any               `json:"adp_csvMerge_csvFieldNativeLocation,omitempty"`
	AdpExecutionPersistent                    bool              `json:"adp_executionPersistent,omitempty"`
	AdpCsvMergeDoNotChangeProtectedDocuments  string            `json:"adp_csvMerge_doNotChangeProtectedDocuments,omitempty"`
	AdpCsvMergeImageBasePath                  any               `json:"adp_csvMerge_imageBasePath,omitempty"`
	AdpLoggingEnabled                         bool              `json:"adp_loggingEnabled,omitempty"`
	AdpCsvMergeNullValue                      string            `json:"adp_csvMerge_nullValue,omitempty"`
	AdpCsvMergeDeduplicateWhenAppending       string            `json:"adp_csvMerge_deduplicateWhenAppending,omitempty"`
	AdpCsvMergeLineSeparatorForFulltext       string            `json:"adp_csvMerge_lineSeparatorForFulltext,omitempty"`
	AdpCsvMergeAllowUpdateMultipleDocs        string            `json:"adp_csvMerge_allowUpdateMultipleDocs,omitempty"`
	AdpCsvMergeCsvIDPostfix                   any               `json:"adp_csvMerge_csvIdPostfix,omitempty"`
	AdpCsvMergeOrigFile                       any               `json:"adp_csvMerge_origFile,omitempty"`
	AdpCsvMergeMaxCategoryLength              string            `json:"adp_csvMerge_maxCategoryLength,omitempty"`
	AdpCsvMergeFieldSeperator                 string            `json:"adp_csvMerge_fieldSeperator,omitempty"`
	AdpCsvMergeApplicationType                string            `json:"adp_csvMerge_applicationType,omitempty"`
	AdpCsvMergeEnginePassword                 any               `json:"adp_csvMerge_enginePassword,omitempty"`
	AdpTaskTimeout                            int               `json:"adp_taskTimeout,omitempty"`
	AdpCsvMergeCharset                        string            `json:"adp_csvMerge_charset,omitempty"`
	AdpProgressTaskTimeout                    int               `json:"adp_progressTaskTimeout,omitempty"`
	AdpCsvMergeTextIndicator                  string            `json:"adp_csvMerge_textIndicator,omitempty"`
	AdpCsvMergeNoFlushBeforeMerge             string            `json:"adp_csvMerge_noFlushBeforeMerge,omitempty"`
	AdpTaskActive                             bool              `json:"adp_taskActive,omitempty"`
	AdpCsvMergeEngineIDFieldKey               any               `json:"adp_csvMerge_engineIdFieldKey,omitempty"`
	AdpAbortWfOnFailure                       bool              `json:"adp_abortWfOnFailure,omitempty"`
	AdpCsvMergeCustomLineSeparator            string            `json:"adp_csvMerge_customLineSeparator,omitempty"`
	AdpCsvMergeMultiValueDelimiter            any               `json:"adp_csvMerge_multiValueDelimiter,omitempty"`
	AdpCsvMergeCsvFile                        any               `json:"adp_csvMerge_csvFile,omitempty"`
	AdpCsvMergeCsvIDFieldKey                  any               `json:"adp_csvMerge_csvIdFieldKey,omitempty"`
	AdpCsvMergeTextFileRefIndicator           string            `json:"adp_csvMerge_textFileRefIndicator,omitempty"`
	AdpCsvMergeEngineUser                     any               `json:"adp_csvMerge_engineUser,omitempty"`
	AdpCsvMergeTextFileCharset                string            `json:"adp_csvMerge_textFileCharset,omitempty"`
	AdpCsvMergeDoNotAddTimestampToOutputFiles string            `json:"adp_csvMerge_doNotAddTimestampToOutputFiles,omitempty"`
	AdpCsvMergeNativeBasePath                 any               `json:"adp_csvMerge_nativeBasePath,omitempty"`
	AdpCsvMergeCsvMode                        string            `json:"adp_csvMerge_csvMode,omitempty"`
	AdpCsvMergeFieldMappings                  []FieldMappingArg `json:"adp_csvMerge_fieldMappings,omitempty"`
	AdpCleanUpHistory                         bool              `json:"adp_cleanUpHistory,omitempty"`
	AdpCsvMergeCsvMergeConfiguration          string            `json:"adp_csvMerge_csvMergeConfiguration,omitempty"`
	AdpCsvMergeErrorLogFile                   any               `json:"adp_csvMerge_errorLogFile,omitempty"`
	AdpCsvMergeDryRun                         string            `json:"adp_csvMerge_dryRun,omitempty"`
	AdpCsvMergeCsvFieldImageLocation          any               `json:"adp_csvMerge_csvFieldImageLocation,omitempty"`
	AdpCsvMergeMatchesLogFile                 any               `json:"adp_csvMerge_matchesLogFile,omitempty"`
	AdpCsvMergeCsvIDPrefix                    any               `json:"adp_csvMerge_csvIdPrefix,omitempty"`
	AdpCsvMergeExpandToFamily                 string            `json:"adp_csvMerge_expandToFamily,omitempty"`
}

func (cfg *CSVMergeConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithCSVMergeLoggingEnabled(enabled bool) func(*CSVMergeConfiguration) {
	return func(c *CSVMergeConfiguration) {
		c.AdpLoggingEnabled = true
	}
}

func WithCSVMergeApplicationIdentifier(id string) func(*CSVMergeConfiguration) {
	return func(c *CSVMergeConfiguration) {
		c.AdpCsvMergeApplicationIdentifier = id
	}
}

func WithCSVMergeEngineName(name string) func(*CSVMergeConfiguration) {
	return func(c *CSVMergeConfiguration) {
		c.AdpCsvMergeEngineName = name
	}
}

func WithCSVMergeCsvIDFieldKey(key string) func(*CSVMergeConfiguration) {
	return func(c *CSVMergeConfiguration) {
		c.AdpCsvMergeCsvIDFieldKey = key
	}
}

func WithCSVMergeMergeType(t string) func(*CSVMergeConfiguration) {
	return func(c *CSVMergeConfiguration) {
		c.AdpCsvMergeMergeType = t
	}
}

func WithCSVMergeEngineIDFieldKey(key string) func(*CSVMergeConfiguration) {
	return func(c *CSVMergeConfiguration) {
		c.AdpCsvMergeEngineIDFieldKey = key
	}
}

func WithCSVMergeCSVFile(file string) func(*CSVMergeConfiguration) {
	return func(c *CSVMergeConfiguration) {
		c.AdpCsvMergeCsvFile = file
	}
}

func WithCSVMergeFieldMappings(mappings []FieldMappingArg) func(*CSVMergeConfiguration) {
	return func(c *CSVMergeConfiguration) {
		c.AdpCsvMergeFieldMappings = mappings
	}
}

func WithCSVMergeAllowUpdateMultipleDocs(s string) func(*CSVMergeConfiguration) {
	return func(c *CSVMergeConfiguration) {
		c.AdpCsvMergeAllowUpdateMultipleDocs = s
	}
}

func WithCSVMergeCSVMode(s string) func(*CSVMergeConfiguration) {
	return func(c *CSVMergeConfiguration) {
		c.AdpCsvMergeCsvMode = s
	}
}
