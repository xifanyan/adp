package adp

type ExportDocumentsConfiguration struct {
	AdpExportDocumentsFieldSeparator                  string `json:"adp_exportDocuments_field_separator,omitempty"`
	AdpProgressTaskTimeout                            int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpExportDocumentsWaitForExport                   bool   `json:"adp_exportDocuments_waitForExport,omitempty"`
	AdpExportDocumentsImageField                      any    `json:"adp_exportDocuments_image_field,omitempty"`
	AdpExportDocumentsSearchResultSize                string `json:"adp_exportDocuments_searchResultSize,omitempty"`
	AdpTaskActive                                     bool   `json:"adp_taskActive,omitempty"`
	AdpExportDocumentsFileEnding                      string `json:"adp_exportDocuments_File_Ending,omitempty"`
	AdpExportDocumentsApplicationType                 string `json:"adp_exportDocuments_applicationType,omitempty"`
	AdpExecutionPersistent                            bool   `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                               bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpExportDocumentsQuery                           string `json:"adp_exportDocuments_query,omitempty"`
	AdpLoggingEnabled                                 bool   `json:"adp_loggingEnabled,omitempty"`
	AdpExportDocumentsExportName                      any    `json:"adp_exportDocuments_exportName,omitempty"`
	AdpExportDocumentsTextIndicator                   string `json:"adp_exportDocuments_text_indicator,omitempty"`
	AdpExportDocumentsNativesField                    any    `json:"adp_exportDocuments_natives_field,omitempty"`
	AdpExportDocumentsMultivalueSeparator             string `json:"adp_exportDocuments_multivalue_separator,omitempty"`
	AdpExportDocumentsLineBreak                       string `json:"adp_exportDocuments_line_break,omitempty"`
	AdpExportDocumentsApplicationIdentifier           string `json:"adp_exportDocuments_applicationIdentifier,omitempty"`
	AdpExportDocumentsEngineIdentifier                any    `json:"adp_exportDocuments_engineIdentifier,omitempty"`
	AdpExportDocumentsExportFileName                  string `json:"adp_exportDocuments_exportFileName,omitempty"`
	AdpExportDocumentsEngineUser                      any    `json:"adp_exportDocuments_engineUser,omitempty"`
	AdpExportDocumentsImageVolume                     string `json:"adp_exportDocuments_image_volume,omitempty"`
	AdpExportDocumentsExportFields                    string `json:"adp_exportDocuments_exportFields,omitempty"`
	AdpExportDocumentsFullExportPath                  string `json:"adp_exportDocuments_fullExportPath,omitempty"`
	AdpExportDocumentsTextField                       any    `json:"adp_exportDocuments_text_field,omitempty"`
	AdpCleanUpHistory                                 bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpExportDocumentsExportDirectory                 any    `json:"adp_exportDocuments_exportDirectory,omitempty"`
	AdpTaskTimeout                                    int    `json:"adp_taskTimeout,omitempty"`
	AdpExportDocumentsEnginePassword                  any    `json:"adp_exportDocuments_enginePassword,omitempty"`
	AdpExportDocumentsAdpExportDocumentsMainQueryType any    `json:"adp_exportDocuments_adp_exportDocuments_mainQueryType,omitempty"`
}

func (cfg *ExportDocumentsConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithExportDocumentsLoggingEnabled(enabled bool) func(*ExportDocumentsConfiguration) {
	return func(c *ExportDocumentsConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithExportDocumentsExportName(name string) func(*ExportDocumentsConfiguration) {
	return func(c *ExportDocumentsConfiguration) {
		c.AdpExportDocumentsExportName = name
	}
}

func WithExportDocumentsExportFileName(fileName string) func(*ExportDocumentsConfiguration) {
	return func(c *ExportDocumentsConfiguration) {
		c.AdpExportDocumentsExportFileName = fileName
	}
}

func WithExportDocumentsApplicationIdentifier(id string) func(*ExportDocumentsConfiguration) {
	return func(c *ExportDocumentsConfiguration) {
		c.AdpExportDocumentsApplicationIdentifier = id
	}
}

func WithExportDocumentsEngineIdentifier(id string) func(*ExportDocumentsConfiguration) {
	return func(c *ExportDocumentsConfiguration) {
		c.AdpExportDocumentsEngineIdentifier = id
	}
}

func WithExportDocumentsFieldSeparator(sep string) func(*ExportDocumentsConfiguration) {
	return func(c *ExportDocumentsConfiguration) {
		c.AdpExportDocumentsFieldSeparator = sep
	}
}

// json "{ "rm_numeric_identifier": "ID" }"
func WithExportDocumentsExportFields(fields string) func(*ExportDocumentsConfiguration) {
	return func(c *ExportDocumentsConfiguration) {
		c.AdpExportDocumentsExportFields = fields
	}
}

func WithExportDocumentsExportDirectory(dir string) func(*ExportDocumentsConfiguration) {
	return func(c *ExportDocumentsConfiguration) {
		c.AdpExportDocumentsExportDirectory = dir
	}
}

func WithExportDocumentsFullExportPath(path string) func(*ExportDocumentsConfiguration) {
	return func(c *ExportDocumentsConfiguration) {
		c.AdpExportDocumentsFullExportPath = path
	}
}

func WithExportDocumentsQuery(query string) func(*ExportDocumentsConfiguration) {
	return func(c *ExportDocumentsConfiguration) {
		c.AdpExportDocumentsQuery = query
	}
}

func WithExportDocumentsNativesField(field string) func(*ExportDocumentsConfiguration) {
	return func(c *ExportDocumentsConfiguration) {
		c.AdpExportDocumentsNativesField = field
	}
}
