package adp

type CreateApplicationConfiguration struct {
	AdpCreateApplicationChosenEngineHostMemory             string `json:"adp_createApplication_chosenEngineHostMemory,omitempty"`
	AdpProgressTaskTimeout                                 int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpCreateApplicationAbortOnExistingApplication         bool   `json:"adp_createApplication_abortOnExistingApplication,omitempty"`
	AdpCreateApplicationApplicationEngineHost              string `json:"adp_createApplication_applicationEngineHost,omitempty"`
	AdpTaskActive                                          bool   `json:"adp_taskActive,omitempty"`
	AdpCreateApplicationApplicationWorkspaceID             string `json:"adp_createApplication_applicationWorkspaceId,omitempty"`
	AdpCreateApplicationApplicationHostMemoryLimitRatio    string `json:"adp_createApplication_applicationHostMemoryLimitRatio,omitempty"`
	AdpCreateApplicationApplicationDescription             string `json:"adp_createApplication_applicationDescription,omitempty"`
	AdpCreateApplicationApplicationHost                    string `json:"adp_createApplication_applicationHost,omitempty"`
	AdpExecutionPersistent                                 bool   `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                                    bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpCreateApplicationChosenApplicationHostMemoryRatio   string `json:"adp_createApplication_chosenApplicationHostMemoryRatio,omitempty"`
	AdpCreateApplicationCreateWorkspace                    string `json:"adp_createApplication_createWorkspace,omitempty"`
	AdpCreateApplicationMetaEngineIdentifier               string `json:"adp_createApplication_metaEngineIdentifier,omitempty"`
	AdpLoggingEnabled                                      bool   `json:"adp_loggingEnabled,omitempty"`
	AdpCreateApplicationEngineHostMemoryLimitRatio         string `json:"adp_createApplication_engineHostMemoryLimitRatio,omitempty"`
	AdpCreateApplicationEngineIdentifier2                  string `json:"adp_createApplication_engineIdentifier_2,omitempty"`
	AdpCreateApplicationApplicationWorkspace               string `json:"adp_createApplication_applicationWorkspace,omitempty"`
	AdpCreateApplicationEngineIdentifier1                  string `json:"adp_createApplication_engineIdentifier_1,omitempty"`
	AdpCreateApplicationApplicationHostDetection           string `json:"adp_createApplication_applicationHostDetection,omitempty"`
	AdpCreateApplicationChosenApplicationHostMemory        string `json:"adp_createApplication_chosenApplicationHostMemory,omitempty"`
	AdpCreateApplicationApplicationContext                 any    `json:"adp_createApplication_applicationContext,omitempty"`
	AdpCreateApplicationApplicationTemplate                string `json:"adp_createApplication_applicationTemplate,omitempty"`
	AdpCreateApplicationApplicationID                      string `json:"adp_createApplication_applicationId,omitempty"`
	AdpCreateApplicationNamespacePrefix                    string `json:"adp_createApplication_namespacePrefix,omitempty"`
	AdpCreateApplicationApplicationDisplayname             string `json:"adp_createApplication_applicationDisplayname,omitempty"`
	AdpCreateApplicationEngineHostMemoryLimit              string `json:"adp_createApplication_engineHostMemoryLimit,omitempty"`
	AdpCreateApplicationChosenEngineHostMemoryRatio        string `json:"adp_createApplication_chosenEngineHostMemoryRatio,omitempty"`
	AdpCreateApplicationApplicationHostMemoryLimit         string `json:"adp_createApplication_applicationHostMemoryLimit,omitempty"`
	AdpCreateApplicationApplicationIdentifier              string `json:"adp_createApplication_applicationIdentifier,omitempty"`
	AdpCreateApplicationApplicationName                    string `json:"adp_createApplication_applicationName,omitempty"`
	AdpCreateApplicationChosenWorkspaceNameParameter       string `json:"adp_createApplication_chosenWorkspaceNameParameter,omitempty"`
	AdpCreateApplicationApplicationTypeEnum                string `json:"adp_createApplication_applicationTypeEnum,omitempty"`
	AdpCleanUpHistory                                      bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpTaskTimeout                                         int    `json:"adp_taskTimeout,omitempty"`
	AdpCreateApplicationChosenApplicationHostNameParameter string `json:"adp_createApplication_chosenApplicationHostNameParameter,omitempty"`
	AdpCreateApplicationApplicationType                    string `json:"adp_createApplication_applicationType,omitempty"`
	AdpCreateApplicationChosenEngineHostNameParameter      string `json:"adp_createApplication_chosenEngineHostNameParameter,omitempty"`
}

func (cfg *CreateApplicationConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithCreateApplicationLoggingEnabled(enabled bool) func(*CreateApplicationConfiguration) {
	return func(c *CreateApplicationConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithCreateApplicationApplicationIdentifier(id string) func(*CreateApplicationConfiguration) {
	return func(c *CreateApplicationConfiguration) {
		c.AdpCreateApplicationApplicationIdentifier = id
	}
}

func WithCreateApplicationApplicationName(name string) func(*CreateApplicationConfiguration) {
	return func(c *CreateApplicationConfiguration) {
		c.AdpCreateApplicationApplicationName = name
	}
}

func WithCreateApplicationApplicationHost(host string) func(*CreateApplicationConfiguration) {
	return func(c *CreateApplicationConfiguration) {
		c.AdpCreateApplicationApplicationHost = host
	}
}

func WithCreateApplicationApplicationType(applicationType string) func(*CreateApplicationConfiguration) {
	return func(c *CreateApplicationConfiguration) {
		c.AdpCreateApplicationApplicationType = applicationType
	}
}

func WithCreateApplicationAppliationWorkspace(workspace string) func(*CreateApplicationConfiguration) {
	return func(c *CreateApplicationConfiguration) {
		c.AdpCreateApplicationApplicationWorkspace = workspace
	}
}

// "true" or "false"
func WithCreateApplicationCreateWorkspace(b string) func(*CreateApplicationConfiguration) {
	return func(c *CreateApplicationConfiguration) {
		c.AdpCreateApplicationCreateWorkspace = b
	}
}

func WithCreateApplicationApplicationTemplate(template string) func(*CreateApplicationConfiguration) {
	return func(c *CreateApplicationConfiguration) {
		c.AdpCreateApplicationApplicationTemplate = template
	}
}

type CreateApplicationResult struct {
	ApplicationDisplayName  string `json:"adp_create_application_application_displayname"`
	ApplicationIdentifier   string `json:"adp_create_application_application_identifier"`
	ApplicationHostMemRatio string `json:"adp_create_application_application_host_memory_ratio"`
	Workspace               string `json:"adp_create_application_workspace"`
	EngineHost              string `json:"adp_create_application_engine_host"`
	ApplicationHostMemory   string `json:"adp_create_application_application_host_memory"`
	ApplicationHost         string `json:"adp_create_application_application_host"`
	EngineHostMemory        string `json:"adp_create_application_engine_host_memory"`
	MetaEngineIdentifier    string `json:"adp_create_application_metaengine_identifier"`
	EngineIdentifier1       string `json:"adp_create_application_engine_identifier_1"`
	EngineIdentifier2       string `json:"adp_create_application_engine_identifier_2"`
	EngineHostMemoryRatio   string `json:"adp_create_application_engine_host_memory_ratio"`
}
