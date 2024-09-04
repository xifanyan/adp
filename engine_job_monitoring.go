package adp

import "encoding/json"

type EngineJobMonitoringConfiguration struct {
	AdpProgressTaskTimeout                    int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpEngineJobMonitoringJobType             string `json:"adp_engineJobMonitoring_jobType,omitempty"`
	AdpEngineJobMonitoringJSONFile            string `json:"adp_engineJobMonitoring_jsonFile,omitempty"`
	AdpTaskActive                             bool   `json:"adp_taskActive,omitempty"`
	AdpEngineJobMonitoringOutputFilename      string `json:"adp_engineJobMonitoring_outputFilename,omitempty"`
	AdpExecutionPersistent                    bool   `json:"adp_executionPersistent,omitempty"`
	AdpEngineJobMonitoringEngineUserName      string `json:"adp_engineJobMonitoring_engineUserName,omitempty"`
	AdpAbortWfOnFailure                       bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpEngineJobMonitoringStartingJob         string `json:"adp_engineJobMonitoring_startingJob,omitempty"`
	AdpEngineJobMonitoringEngineIds           string `json:"adp_engineJobMonitoring_engineIds,omitempty"`
	AdpCleanUpHistory                         bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpEngineJobMonitoringWaitForJobs         string `json:"adp_engineJobMonitoring_waitForJobs,omitempty"`
	AdpLoggingEnabled                         bool   `json:"adp_loggingEnabled,omitempty"`
	AdpEngineJobMonitoringEngineUserPassword  string `json:"adp_engineJobMonitoring_engineUserPassword,omitempty"`
	AdpEngineJobMonitoringOutputJSON          string `json:"adp_engineJobMonitoring_outputJson,omitempty"`
	AdpTaskTimeout                            int    `json:"adp_taskTimeout,omitempty"`
	AdpEngineJobMonitoringJobStates           string `json:"adp_engineJobMonitoring_jobStates,omitempty"`
	AdpEngineJobMonitoringStartedAfter        string `json:"adp_engineJobMonitoring_startedAfter,omitempty"`
	AdpEngineJobMonitoringListOfJobProperties string `json:"adp_engineJobMonitoring_listOfJobProperties,omitempty"`
	AdpEngineJobMonitoringWorkspaceID         string `json:"adp_engineJobMonitoring_workspaceId,omitempty"`
	AdpEngineJobMonitoringNumberOfJobs        string `json:"adp_engineJobMonitoring_numberOfJobs,omitempty"`
	AdpEngineJobMonitoringAppIds              string `json:"adp_engineJobMonitoring_appIds,omitempty"`
}

func (cfg *EngineJobMonitoringConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithEngineJobMonitoringLoggingEnabled(enabled bool) func(*EngineJobMonitoringConfiguration) {
	return func(c *EngineJobMonitoringConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithEngineJobMonitoringEngineIds(engineIds string) func(*EngineJobMonitoringConfiguration) {
	return func(c *EngineJobMonitoringConfiguration) {
		c.AdpEngineJobMonitoringEngineIds = engineIds
	}
}

func WithEngineJobMonitoringAppIds(appIds string) func(*EngineJobMonitoringConfiguration) {
	return func(c *EngineJobMonitoringConfiguration) {
		c.AdpEngineJobMonitoringAppIds = appIds
	}
}

func WithEngineJobMonitoringJobStates(jobStates string) func(*EngineJobMonitoringConfiguration) {
	return func(c *EngineJobMonitoringConfiguration) {
		c.AdpEngineJobMonitoringJobStates = jobStates
	}
}

func WithEngineJobMonitoringJobType(jobType string) func(*EngineJobMonitoringConfiguration) {
	return func(c *EngineJobMonitoringConfiguration) {
		c.AdpEngineJobMonitoringJobType = jobType
	}
}

func WithEngineJobMonitoringStartedAfter(startedAfter string) func(*EngineJobMonitoringConfiguration) {
	return func(c *EngineJobMonitoringConfiguration) {
		c.AdpEngineJobMonitoringStartedAfter = startedAfter
	}
}

type EngineJobMonitoringExecutionMetaData struct {
	AdpEngineJobMonitoringJSONOutput     json.RawMessage `json:"adp_engineJobMonitoring_json_output"`
	AdpEngineJobMonitoringOutputFileName json.RawMessage `json:"adp_engineJobMonitoring_output_file_name"`
}

type PhaseProgressEstimation struct {
	CompletionPercentage  float32 `json:"completionPercentage,omitempty"`
	ExpectedNumberOfSteps int     `json:"expectedNumberOfSteps,omitempty"`
	NumberOfFinishedSteps int     `json:"numberOfFinishedSteps,omitempty"`
	EstimatedDuration     int     `json:"estimatedDuration,omitempty"`
}

type JobPhase struct {
	PhaseDisplayName           string                  `json:"phaseDisplayName,omitempty"`
	PhaseDescription           string                  `json:"phaseDescription,omitempty"`
	PhaseStartDate             int                     `json:"phaseStartDate,omitempty"`
	PhaseFinishDate            int                     `json:"phaseFinishDate,omitempty"`
	PhaseNumberOfErrors        int                     `json:"phaseNumberOfSteps,omitempty"`
	PhaseNumberOfFinishedSteps int                     `json:"phaseNumberOfFinishedSteps,omitempty"`
	PhaseNumberOfSkippedSteps  int                     `json:"phaseNumberOfSkippedSteps,omitempty"`
	PhaseProgressEstimation    PhaseProgressEstimation `json:"phaseProgressEstimation,omitempty"`
}

type EngineJobs []EngineJob

type JobProperties struct {
	Completion    float32 `json:"completion,omitempty"`
	DisplayName   string  `json:"displayName,omitempty"`
	Errors        int     `json:"errors,omitempty"`
	FinishedOn    string  `json:"finishedOn,omitempty"`
	JobType       string  `json:"jobType,omitempty"`
	Owner         string  `json:"owner,omitempty"`
	PausedOn      string  `json:"pausedOn,omitempty"`
	Phase         string  `json:"phase,omitempty"`
	StartedOn     string  `json:"startedOn,omitempty"`
	Status        string  `json:"status,omitempty"`
	StepsExpected int     `json:"stepsExpected,omitempty"`
	StepsFinished int     `json:"stepsFinished,omitempty"`
}

type EngineJob struct {
	JobID                 string        `json:"jobId,omitempty"`
	ApplicationIdentifier string        `json:"applicationIdentifier,omitempty"`
	EngineIdentifier      string        `json:"engineIdentifier,omitempty"`
	HostName              string        `json:"hostName,omitempty"`
	JobProperties         JobProperties `json:"jobProperties,omitempty"`
	JobPhases             []JobPhase    `json:"jobPhases,omitempty"`
}

type EngineJobMonitoringResult struct {
	Date       string      `json:"date"`
	EngineJobs []EngineJob `json:"engineJobs"`
}
