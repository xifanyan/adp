package adp

type UnicodeCharacter struct {
	UnicodeCharacter string `json:"Unicode character"`
}

type CsvDetectionConfiguration struct {
	AdpProgressTaskTimeout                int                `json:"adp_progressTaskTimeout,omitempty"`
	AdpCsvDetectionDetectedFieldSeparator string             `json:"adp_csvDetection_detectedFieldSeparator,omitempty"`
	AdpCsvDetectionReadBufferSize         int                `json:"adp_csvDetection_readBufferSize,omitempty"`
	AdpTaskActive                         bool               `json:"adp_taskActive,omitempty"`
	AdpCsvDetectionSetEncoding            bool               `json:"adp_csvDetection_setEncoding,omitempty"`
	AdpCsvDetectionValueDelimiter         []UnicodeCharacter `json:"adp_csvDetection_valueDelimiter,omitempty"`
	AdpCsvDetectionFieldSeparator         []UnicodeCharacter `json:"adp_csvDetection_fieldSeparator,omitempty"`
	AdpExecutionPersistent                bool               `json:"adp_executionPersistent,omitempty"`
	AdpCsvDetectionCsvFile                string             `json:"adp_csvDetection_csvFile,omitempty"`
	AdpCsvDetectionEncoding               string             `json:"adp_csvDetection_encoding,omitempty"`
	AdpAbortWfOnFailure                   bool               `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                     bool               `json:"adp_cleanUpHistory,omitempty"`
	AdpCsvDetectionMaxCharactersLoad      int                `json:"adp_csvDetection_maxCharactersLoad,omitempty"`
	AdpLoggingEnabled                     bool               `json:"adp_loggingEnabled,omitempty"`
	AdpCsvDetectionAmountOfLinesToLoad    string             `json:"adp_csvDetection_amountOfLinesToLoad,omitempty"`
	AdpCsvDetectionDetectedValueDelimiter string             `json:"adp_csvDetection_detectedValueDelimiter,omitempty"`
	AdpTaskTimeout                        int                `json:"adp_taskTimeout,omitempty"`
	AdpCsvDetectionDetectedLineSeparator  string             `json:"adp_csvDetection_detectedLineSeparator,omitempty"`
	AdpCsvDetectionDetectedEncoding       string             `json:"adp_csvDetection_detectedEncoding,omitempty"`
	AdpCsvDetectionMaxLineLength          int                `json:"adp_csvDetection_maxLineLength,omitempty"`
}

func (cfg *CsvDetectionConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithAdpCSVDetectionLoggingEnabled(enabled bool) func(*CsvDetectionConfiguration) {
	return func(c *CsvDetectionConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithAdpCSVDetectionCsvFile(file string) func(*CsvDetectionConfiguration) {
	return func(c *CsvDetectionConfiguration) {
		c.AdpCsvDetectionCsvFile = file
	}
}

func WithAdpCSVDetectionAmountOfLinesToLoad(amount string) func(*CsvDetectionConfiguration) {
	return func(c *CsvDetectionConfiguration) {
		c.AdpCsvDetectionAmountOfLinesToLoad = amount
	}
}

type CsvDetectionResult struct {
	DetectedLineSeparator  string `json:"detected_line_separator"`
	DetectedFieldSeparator string `json:"detected_field_separator"`
	DetectedEncoding       string `json:"detected_encoding"`
	DetectedValueDelimiter string `json:"detected_value_delimiter"`
}
