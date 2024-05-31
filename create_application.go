package adp

type CreateApplicationConfiguration struct {
	AdpCreateApplicationChosenEngineHostMemory             string `json:"adp_createApplication_chosenEngineHostMemory,omitempty"`
	AdpProgressTaskTimeout                                 int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpCreateApplicationAbortOnExistingApplication         bool   `json:"adp_createApplication_abortOnExistingApplication,omitempty"`
	AdpCreateApplicationApplicationEngineHost              any    `json:"adp_createApplication_applicationEngineHost,omitempty"`
	AdpTaskActive                                          bool   `json:"adp_taskActive,omitempty"`
	AdpCreateApplicationApplicationWorkspaceID             any    `json:"adp_createApplication_applicationWorkspaceId,omitempty"`
	AdpCreateApplicationApplicationHostMemoryLimitRatio    string `json:"adp_createApplication_applicationHostMemoryLimitRatio,omitempty"`
	AdpCreateApplicationApplicationDescription             any    `json:"adp_createApplication_applicationDescription,omitempty"`
	AdpCreateApplicationApplicationHost                    any    `json:"adp_createApplication_applicationHost,omitempty"`
	AdpExecutionPersistent                                 bool   `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                                    bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpCreateApplicationChosenApplicationHostMemoryRatio   string `json:"adp_createApplication_chosenApplicationHostMemoryRatio,omitempty"`
	AdpCreateApplicationCreateWorkspace                    string `json:"adp_createApplication_createWorkspace,omitempty"`
	AdpCreateApplicationMetaEngineIdentifier               string `json:"adp_createApplication_metaEngineIdentifier,omitempty"`
	AdpLoggingEnabled                                      bool   `json:"adp_loggingEnabled,omitempty"`
	AdpCreateApplicationEngineHostMemoryLimitRatio         string `json:"adp_createApplication_engineHostMemoryLimitRatio,omitempty"`
	AdpCreateApplicationEngineIdentifier2                  string `json:"adp_createApplication_engineIdentifier_2,omitempty"`
	AdpCreateApplicationApplicationWorkspace               any    `json:"adp_createApplication_applicationWorkspace,omitempty"`
	AdpCreateApplicationEngineIdentifier1                  string `json:"adp_createApplication_engineIdentifier_1,omitempty"`
	AdpCreateApplicationApplicationHostDetection           string `json:"adp_createApplication_applicationHostDetection,omitempty"`
	AdpCreateApplicationChosenApplicationHostMemory        string `json:"adp_createApplication_chosenApplicationHostMemory,omitempty"`
	AdpCreateApplicationApplicationContext                 any    `json:"adp_createApplication_applicationContext,omitempty"`
	AdpCreateApplicationApplicationTemplate                any    `json:"adp_createApplication_applicationTemplate,omitempty"`
	AdpCreateApplicationApplicationID                      any    `json:"adp_createApplication_applicationId,omitempty"`
	AdpCreateApplicationNamespacePrefix                    any    `json:"adp_createApplication_namespacePrefix,omitempty"`
	AdpCreateApplicationApplicationDisplayname             string `json:"adp_createApplication_applicationDisplayname,omitempty"`
	AdpCreateApplicationEngineHostMemoryLimit              string `json:"adp_createApplication_engineHostMemoryLimit,omitempty"`
	AdpCreateApplicationChosenEngineHostMemoryRatio        string `json:"adp_createApplication_chosenEngineHostMemoryRatio,omitempty"`
	AdpCreateApplicationApplicationHostMemoryLimit         string `json:"adp_createApplication_applicationHostMemoryLimit,omitempty"`
	AdpCreateApplicationApplicationIdentifier              string `json:"adp_createApplication_applicationIdentifier,omitempty"`
	AdpCreateApplicationApplicationName                    any    `json:"adp_createApplication_applicationName,omitempty"`
	AdpCreateApplicationChosenWorkspaceNameParameter       string `json:"adp_createApplication_chosenWorkspaceNameParameter,omitempty"`
	AdpCreateApplicationApplicationTypeEnum                string `json:"adp_createApplication_applicationTypeEnum,omitempty"`
	AdpCleanUpHistory                                      bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpTaskTimeout                                         int    `json:"adp_taskTimeout,omitempty"`
	AdpCreateApplicationChosenApplicationHostNameParameter string `json:"adp_createApplication_chosenApplicationHostNameParameter,omitempty"`
	AdpCreateApplicationApplicationType                    any    `json:"adp_createApplication_applicationType,omitempty"`
	AdpCreateApplicationChosenEngineHostNameParameter      string `json:"adp_createApplication_chosenEngineHostNameParameter,omitempty"`
}
