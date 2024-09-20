package adp

import (
	"encoding/json"
)

type configurator interface {
	enforcePersistentExecution()
}

type Request struct {
	TaskType          string       `json:"taskType"`
	TaskDescription   string       `json:"taskDescription"`
	TaskDisplayName   string       `json:"taskDisplayName"`
	TaskConfiguration configurator `json:"taskConfiguration"`
}

func NewRequest() *Request {
	return &Request{}
}

func (req *Request) WithTaskType(typ string) *Request {
	req.TaskType = typ
	return req
}

func (req *Request) WithTaskDescription(description string) *Request {
	req.TaskDescription = description
	return req
}

func (req *Request) WithTaskDisplayName(displayName string) *Request {
	req.TaskDisplayName = displayName
	return req
}

func (req *Request) WithTaskConfiguration(cfg configurator) *Request {
	req.TaskConfiguration = cfg
	return req
}

// JSON generates a JSON representation of the Request.
//
// Returns a string and an error.
func (req *Request) JSON() string {
	jsonBytes, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func newTaskRequest[T configurator](taskType string, defaults T, opts ...func(T)) *Request {
	cfg := defaults

	for _, opt := range opts {
		opt(cfg)
	}

	return NewRequest().
		WithTaskType(taskType).
		WithTaskConfiguration(cfg)
}

// ListEntities applies optional configuration functions to update the configuration object and returns a new request.
//
// opts: optional configuration functions
// *Request: the updated request object
func (req *Request) ListEntities(opts ...func(*ListEntitiesConfiguration)) *Request {
	defaults := &ListEntitiesConfiguration{
		AdpLoggingEnabled:        false,
		AdpExecutionPersistent:   false,
		AdpListEntitiesWhiteList: DefaultListEntitiesWhitelist,
	}

	return newTaskRequest("List Entities", defaults, opts...)
}

func (req *Request) TaxonomyStatistic(opts ...func(*TaxonomyStatisticConfiguration)) *Request {
	defaults := &TaxonomyStatisticConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Taxonomy Statistic", defaults, opts...)
}

// ComputeCounts updates the configuration object with optional settings and returns a new Request.
//
// opts: variadic list of functions to update the configuration object.
// *Request: returns a pointer to the updated Request object.
func (req *Request) ComputeCounts(opts ...func(*ComputeCountsConfiguration)) *Request {
	defaults := &ComputeCountsConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Compute Counts", defaults, opts...)
}

// PingProject pings a project with the provided configuration options.
//
// opts: variadic function options to customize the ping project request.
// *Request: returns a pointer to the modified Request.
func (req *Request) PingProject(opts ...func(*PingProjectConfiguration)) *Request {
	defaults := &PingProjectConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Ping Project", defaults, opts...)
}

// QueryEngine generates a new Request with the Query Engine task type and configuration.
//
// opts: variadic functions to configure the Query Engine.
// *Request: returns a pointer to the modified Request.
func (req *Request) QueryEngine(opts ...func(*QueryEngineConfiguration)) *Request {
	defaults := &QueryEngineConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Query Engine", defaults, opts...)
}

// StartApplication starts the application with the given configuration options.
//
// opts: variadic list of functions to configure the start application process.
// *Request: returns a pointer to a Request.
func (req *Request) StartApplication(opts ...func(*StartApplicationConfiguration)) *Request {
	defaults := &StartApplicationConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Start Application", defaults, opts...)
}

// StopProcesses stops processes based on the provided configuration.
//
// opts: variadic functions to configure the stopping of processes.
// *Request: returns a pointer to a Request.
func (req *Request) StopProcesses(opts ...func(*StopProcessesConfiguration)) *Request {
	defaults := &StopProcessesConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Stop Processes", defaults, opts...)
}

// QueryPostgresqlDB generates a request to query a Postgresql database with the provided configurations.
//
// opts: variadic functions to configure QueryPostgresqlDB.
// *Request: returns a pointer to the generated request.
func (req *Request) QueryPostgresqlDB(opts ...func(*QueryPostgresqlDBConfiguration)) *Request {
	defaults := &QueryPostgresqlDBConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Query Postgresql DB", defaults, opts...)
}

// ManageTaxonomy manages the taxonomy based on the provided configurations.
//
// opts: variadic functions to configure ManageTaxonomy.
// *Request: returns a Request object.
func (req *Request) ManageTaxonomy(opts ...func(*ManageTaxonomyConfiguration)) *Request {
	defaults := &ManageTaxonomyConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Manage Taxonomy", defaults, opts...)
}

// ReadServiceAlerts reads service alerts from the data source.
//
// Parameters:
//
//	opts: Optional configurations for reading service alerts.
//
// Returns:
//
//	*Request: A request object configured to read service alerts.
func (req *Request) ReadServiceAlerts(opts ...func(*ReadServiceAlertsConfiguration)) *Request {
	defaults := &ReadServiceAlertsConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Read Service Alerts", defaults, opts...)
}

func (req *Request) ConfigureDataSource(opts ...func(*ConfigureDataSourceConfiguration)) *Request {
	defaults := &ConfigureDataSourceConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Configure Data Source", defaults, opts...)
}

func (req *Request) StartDataSource(opts ...func(*StartDataSourceConfiguration)) *Request {
	defaults := &StartDataSourceConfiguration{
		AdpLoggingEnabled:             false,
		AdpExecutionPersistent:        false,
		AdpStartDataSourceSynchronous: false,
	}

	return newTaskRequest("Start Data Source", defaults, opts...)
}

func (req *Request) ExportDocuments(opts ...func(*ExportDocumentsConfiguration)) *Request {
	defaults := &ExportDocumentsConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Export Documents", defaults, opts...)
}

func (req *Request) CLI(opts ...func(*CLIConfiguration)) *Request {
	defaults := &CLIConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("CLI", defaults, opts...)
}

func (req *Request) CreateCustodian(opts ...func(*CreateCustodianConfiguration)) *Request {
	defaults := &CreateCustodianConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Create Custodian", defaults, opts...)
}

func (req *Request) CSVMerge(opts ...func(*CSVMergeConfiguration)) *Request {
	defaults := &CSVMergeConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("CSV Merge", defaults, opts...)
}

func (req *Request) CreateDataSource(opts ...func(*CreateDataSourceConfiguration)) *Request {
	defaults := &CreateDataSourceConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Create Data Source", defaults, opts...)
}

func (req *Request) ManageUsersAndGroups(opts ...func(*ManageUsersAndGruopsConfiguration)) *Request {
	defaults := &ManageUsersAndGruopsConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Manage Users and Groups", defaults, opts...)
}

func (req *Request) ReadConfiguration(opts ...func(*ReadConfigurationConfiguration)) *Request {
	defaults := &ReadConfigurationConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Read Configuration", defaults, opts...)
}

func (req *Request) ConfigureEntity(opts ...func(*ConfigureEntityConfiguration)) *Request {
	defaults := &ConfigureEntityConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Configure Entity", defaults, opts...)
}

func (req *Request) AddSubengine(opts ...func(*AddSubengineConfiguration)) *Request {
	defaults := &AddSubengineConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Add Subengine", defaults, opts...)
}

func (req *Request) EngineJobMonitoring(opts ...func(*EngineJobMonitoringConfiguration)) *Request {
	defaults := &EngineJobMonitoringConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Engine Job Monitoring", defaults, opts...)
}

func (req *Request) CsvDetection(opts ...func(*CsvDetectionConfiguration)) *Request {
	defaults := &CsvDetectionConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
		AdpCsvDetectionValueDelimiter: []UnicodeCharacter{
			{UnicodeCharacter: "u+0022"},
			{UnicodeCharacter: "u+00FE"},
			{UnicodeCharacter: "u+005E"},
		},
		AdpCsvDetectionFieldSeparator: []UnicodeCharacter{
			{UnicodeCharacter: "u+002C"},
			{UnicodeCharacter: "u+0014"},
			{UnicodeCharacter: "u+007C"},
			{UnicodeCharacter: "u+003B"},
			{UnicodeCharacter: "u+0009"},
		},
	}

	return newTaskRequest("CSV Detection", defaults, opts...)
}

func (req *Request) AddFields(opts ...func(*AddFieldsConfiguration)) *Request {
	defaults := &AddFieldsConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
		AdpAddFieldsCloneTemplates: []FieldCloneTemplate{
			{
				Type:  "Smart Filter",
				Field: "meta_bcc",
			},
			{
				Type:  "Text field",
				Field: "attachment",
			},
			{
				Type:  "Unique Dates",
				Field: "meta_docdate",
			},
			{
				Type:  "Grouped Dates",
				Field: "meta_bcc",
			},
		},
	}
	return newTaskRequest("Add Fields", defaults, opts...)
}

func (req *Request) ConfigureServiceProperties(opts ...func(*ConfigureServicePropertiesConfiguration)) *Request {
	defaults := &ConfigureServicePropertiesConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Configure Service Properties", defaults, opts...)
}
