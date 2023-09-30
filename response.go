package adp

import "encoding/json"

type Response struct {
	ExecutionID         string          `json:"executionId"`
	TaskType            string          `json:"taskType"`
	LoggingEnabled      string          `json:"loggingEnabled"`
	ProgressMax         int             `json:"progressMax"`
	ExecutionStatus     string          `json:"executionStatus"`
	ExecutionRootDir    string          `json:"executionRootDir"`
	ContextID           string          `json:"contextId"`
	ExecutionPersistent string          `json:"executionPersistent"`
	ProgressCurrent     int             `json:"progressCurrent"`
	ProgressPercentage  float64         `json:"progressPercentage"`
	TaskDisplayName     string          `json:"taskDisplayName"`
	ExecutionMetaData   json.RawMessage `json:"executionMetaData"`
}
