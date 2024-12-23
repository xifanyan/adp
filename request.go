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

// newTaskRequest initializes a new Request with a specified task type and configuration.
//
// Parameters:
//   taskType: A string representing the type of task for the request.
//   defaults: A configuration object of a generic type T that implements the configurator interface,
//             providing default values for the configuration.
//   opts: A variadic list of functions that modify the configuration object of type T.
//
// Returns:
//   *Request: A pointer to the newly created Request object configured with the specified task type
//             and updated configuration.

func newTaskRequest[T configurator](taskType string, defaults T, opts ...func(T)) *Request {
	cfg := defaults

	for _, opt := range opts {
		opt(cfg)
	}

	return NewRequest().
		WithTaskType(taskType).
		WithTaskConfiguration(cfg)
}

// ListEntities initializes a new Request with a specified task type and configuration
// for listing entities.
//
// Parameters:
//
//	opts: A variadic list of functions that modify the configuration object of
//	      type *ListEntitiesConfiguration.
//
// Returns:
//
//	*Request: A pointer to the newly created Request object configured with the
//	          specified task type and updated configuration.
func (req *Request) ListEntities(opts ...func(*ListEntitiesConfiguration)) *Request {
	defaults := &ListEntitiesConfiguration{
		AdpLoggingEnabled:        false,
		AdpExecutionPersistent:   false,
		AdpListEntitiesWhiteList: DefaultListEntitiesWhitelist,
	}

	return newTaskRequest("List Entities", defaults, opts...)
}

// TaxonomyStatistic initializes a new Request with a specified task type and configuration
// for retrieving a taxonomy statistic.
//
// Parameters:
//
//	opts: A variadic list of functions that modify the configuration object of
//	      type *TaxonomyStatisticConfiguration.
//
// Returns:
//
//	*Request: A pointer to the newly created Request object configured with the
//	          specified task type and updated configuration.
func (req *Request) TaxonomyStatistic(opts ...func(*TaxonomyStatisticConfiguration)) *Request {
	defaults := &TaxonomyStatisticConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Taxonomy Statistic", defaults, opts...)
}

// ComputeCounts initializes a new Request with a specified task type and configuration
// for computing counts.
//
// Parameters:
//
//	opts: A variadic list of functions that modify the configuration object of
//	      type *ComputeCountsConfiguration.
//
// Returns:
//
//	*Request: A pointer to the newly created Request object configured with the
//	          specified task type and updated configuration.
func (req *Request) ComputeCounts(opts ...func(*ComputeCountsConfiguration)) *Request {
	defaults := &ComputeCountsConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Compute Counts", defaults, opts...)
}

// PingProject initializes a new Request with a specified task type and configuration
// for pinging a project.
//
// Parameters:
//
//	opts: A variadic list of functions that modify the configuration object of
//	      type *PingProjectConfiguration.
//
// Returns:
//
//	*Request: A pointer to the newly created Request object configured with the
//	          specified task type and updated configuration.
func (req *Request) PingProject(opts ...func(*PingProjectConfiguration)) *Request {
	defaults := &PingProjectConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Ping Project", defaults, opts...)
}

// QueryEngine initializes a new Request with a specified task type and configuration
// for querying a search engine.
//
// Parameters:
//
//	opts: A variadic list of functions that modify the configuration object of
//	      type *QueryEngineConfiguration.
//
// Returns:
//
//	*Request: A pointer to the newly created Request object configured with the
//	          specified task type and updated configuration.
func (req *Request) QueryEngine(opts ...func(*QueryEngineConfiguration)) *Request {
	defaults := &QueryEngineConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Query Engine", defaults, opts...)
}

// StartApplication initializes a new Request with a specified task type and configuration
// for starting an application.
//
// Parameters:
//   opts: A variadic list of functions that modify the configuration object of
//         type *StartApplicationConfiguration.
//
// Returns:
//   *Request: A pointer to the newly created Request object configured with the
//             specified task type and updated configuration.

func (req *Request) StartApplication(opts ...func(*StartApplicationConfiguration)) *Request {
	defaults := &StartApplicationConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Start Application", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// StopProcesses initializes a new Request with a specified task type and configuration
// for stopping a list of processes.
//
// Parameters:
//   opts: A variadic list of functions that modify the configuration object of
//         type *StopProcessesConfiguration.
//
// Returns:
//   *Request: A pointer to the newly created Request object configured with the
//             specified task type and updated configuration.
/******  744ffa90-8eb4-420e-b5bf-7ed721b2df4c  *******/
func (req *Request) StopProcesses(opts ...func(*StopProcessesConfiguration)) *Request {
	defaults := &StopProcessesConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Stop Processes", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// QueryPostgresqlDB initializes a new Request with a specified task type and configuration
// for querying a Postgresql database.
//
// Parameters:
//   opts: A variadic list of functions that modify the configuration object of
//         type *QueryPostgresqlDBConfiguration.
//
// Returns:
//   *Request: A pointer to the newly created Request object configured with the
//             specified task type and updated configuration.
/******  d161183c-39d1-495e-a28f-75738e28e0ed  *******/
func (req *Request) QueryPostgresqlDB(opts ...func(*QueryPostgresqlDBConfiguration)) *Request {
	defaults := &QueryPostgresqlDBConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Query Postgresql DB", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// ManageTaxonomy initializes a new Request with a specified task type and configuration
// for managing taxonomy.
//
// Parameters:
//   opts: A variadic list of functions that modify the configuration object of
//         type *ManageTaxonomyConfiguration.
//
// Returns:
//   *Request: A pointer to the newly created Request object configured with the
//             specified task type and updated configuration.
/******  ed9b820c-9fef-40ee-8581-fa7d21abee79  *******/
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

// ConfigureDataSource configures a data source with the provided configurations.
//
// Parameters:
//
//	opts: Optional configurations for configuring the data source.
//
// Returns:
//
//	*Request: A request object configured to configure a data source.
func (req *Request) ConfigureDataSource(opts ...func(*ConfigureDataSourceConfiguration)) *Request {
	defaults := &ConfigureDataSourceConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Configure Data Source", defaults, opts...)
}

// StartDataSource initializes a request to start a data source with the specified configurations.
//
// Parameters:
//   opts: A variadic list of functions that modify the StartDataSourceConfiguration.
//         These can be used to customize the behavior of the data source start process.
//
// Returns:
//   *Request: A pointer to the Request object configured to start a data source.

func (req *Request) StartDataSource(opts ...func(*StartDataSourceConfiguration)) *Request {
	defaults := &StartDataSourceConfiguration{
		AdpLoggingEnabled:             false,
		AdpExecutionPersistent:        false,
		AdpStartDataSourceSynchronous: false,
	}

	return newTaskRequest("Start Data Source", defaults, opts...)
}

// ExportDocuments initializes a request to export documents with the specified configurations.
//
// Parameters:
//
//	opts: A variadic list of functions that modify the ExportDocumentsConfiguration.
//	      These can be used to customize the behavior of the export documents process.
//
// Returns:
//
//	*Request: A pointer to the Request object configured to export documents.
func (req *Request) ExportDocuments(opts ...func(*ExportDocumentsConfiguration)) *Request {
	defaults := &ExportDocumentsConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Export Documents", defaults, opts...)
}

// CLI initializes a request to run a command line interface with the specified configurations.
//
// Parameters:
//
//	opts: A variadic list of functions that modify the CLIConfiguration.
//	      These can be used to customize the behavior of the CLI process.
//
// Returns:
//
//	*Request: A pointer to the Request object configured to run a CLI.
func (req *Request) CLI(opts ...func(*CLIConfiguration)) *Request {
	defaults := &CLIConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("CLI", defaults, opts...)
}

// CreateCustodian initializes a request to create a custodian with the specified configurations.
//
// Parameters:
//   opts: A variadic list of functions that modify the CreateCustodianConfiguration.
//         These can be used to customize the behavior of the create custodian process.
//
// Returns:
//   *Request: A pointer to the Request object configured to create a custodian.

func (req *Request) CreateCustodian(opts ...func(*CreateCustodianConfiguration)) *Request {
	defaults := &CreateCustodianConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}

	return newTaskRequest("Create Custodian", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// CSVMerge initializes a request to merge CSV files with the specified configurations.
//
// Parameters:
//   opts: A variadic list of functions that modify the CSVMergeConfiguration.
//         These can be used to customize the behavior of the CSV merge process.
//
// Returns:
//   *Request: A pointer to the Request object configured to merge CSV files.
/******  80aed027-5f11-4d39-b7be-271902962b1e  *******/
func (req *Request) CSVMerge(opts ...func(*CSVMergeConfiguration)) *Request {
	defaults := &CSVMergeConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("CSV Merge", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// CreateDataSource initializes a request to create a data source with the specified configurations.
//
// Parameters:
//
//	opts: A variadic list of functions that modify the CreateDataSourceConfiguration.
//         These can be used to customize the behavior of the create data source process.
//
// Returns:
//   *Request: A pointer to the Request object configured to create a data source.
/******  1b121b0a-03bc-440c-84b8-08b20b9388c3  *******/
func (req *Request) CreateDataSource(opts ...func(*CreateDataSourceConfiguration)) *Request {
	defaults := &CreateDataSourceConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Create Data Source", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// ManageUsersAndGroups initializes a request to manage users and groups with the specified configurations.
//
// Parameters:
//   opts: A variadic list of functions that modify the ManageUsersAndGruopsConfiguration.
//         These can be used to customize the behavior of the manage users and groups process.
//
// Returns:
//   *Request: A pointer to the Request object configured to manage users and groups.
/******  9cf745ab-f315-4fb9-a717-2689cc2f9a26  *******/
func (req *Request) ManageUsersAndGroups(opts ...func(*ManageUsersAndGruopsConfiguration)) *Request {
	defaults := &ManageUsersAndGruopsConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Manage Users and Groups", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// ReadConfiguration initializes a request to read configuration data with the specified options.
//
// Parameters:
//   opts: A variadic list of functions that modify the ReadConfigurationConfiguration.
//         These can be used to customize the behavior of the read configuration process.
//
// Returns:
//   *Request: A pointer to the Request object configured to read configuration data.

/******  0c508ea1-5010-4faa-ad17-d6790abf1c7b  *******/
func (req *Request) ReadConfiguration(opts ...func(*ReadConfigurationConfiguration)) *Request {
	defaults := &ReadConfigurationConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Read Configuration", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// ConfigureEntity initializes a request to configure an entity with the specified configurations.
//
// Parameters:
//   opts: A variadic list of functions that modify the ConfigureEntityConfiguration.
//         These can be used to customize the behavior of the configure entity process.
//
// Returns:
//   *Request: A pointer to the Request object configured to configure an entity.
/******  bfa99828-530b-4687-8f3a-074feb88401a  *******/
func (req *Request) ConfigureEntity(opts ...func(*ConfigureEntityConfiguration)) *Request {
	defaults := &ConfigureEntityConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Configure Entity", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// AddSubengine initializes a request to add a subengine with the specified configurations.
//
// Parameters:
//   opts: A variadic list of functions that modify the AddSubengineConfiguration.
//         These can be used to customize the behavior of the add subengine process.
//
// Returns:
//   *Request: A pointer to the Request object configured to add a subengine.

/******  61580160-fc5c-4ca5-9201-434107d3c5ea  *******/
func (req *Request) AddSubengine(opts ...func(*AddSubengineConfiguration)) *Request {
	defaults := &AddSubengineConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Add Subengine", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// EngineJobMonitoring initializes a request to monitor/list engine jobs with the specified configurations.
//
// Parameters:
//   opts: A variadic list of functions that modify the EngineJobMonitoringConfiguration.
//         These can be used to customize the behavior of the engine job monitoring process.
//
// Returns:
//   *Request: A pointer to the Request object configured to monitor/list engine jobs.
/******  caa5addb-eed4-4714-a47e-ba85d0240809  *******/
func (req *Request) EngineJobMonitoring(opts ...func(*EngineJobMonitoringConfiguration)) *Request {
	defaults := &EngineJobMonitoringConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Engine Job Monitoring", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// CsvDetection initializes a request to detect a CSV file's delimiter and quote
// character with the specified configurations.
//
// Parameters:
//   opts: A variadic list of functions that modify the CsvDetectionConfiguration.
//         These can be used to customize the behavior of the CSV detection process.
//
// Returns:
//   *Request: A pointer to the Request object configured to detect a CSV file's
//             delimiter and quote character.
/******  b0489d7a-f571-4565-ba06-fd27d059c4d6  *******/
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

/*************  ✨ Codeium Command ⭐  *************/
// AddFields initializes a request to add fields to an engine with the specified
// configurations.
//
// Parameters:
//   opts: A variadic list of functions that modify the AddFieldsConfiguration.
//         These can be used to customize the behavior of the add fields process.
//
// Returns:
//   *Request: A pointer to the Request object configured to add fields to an engine.
/******  4e4220ea-a165-4e5a-b492-72202dd8c279  *******/
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

/*************  ✨ Codeium Command ⭐  *************/
// ConfigureServiceProperties configures service properties using the given
// configuration options.
//
// Parameters:
//   opts: A variadic list of functions that modify the
//         ConfigureServicePropertiesConfiguration. These can be used to
//         customize the behavior of the configure service properties process.
//
// Returns:
//   *Request: A pointer to the Request object configured to configure service
//             properties.
/******  507f1e9b-9579-4ce0-946d-48ed687bbd4b  *******/
func (req *Request) ConfigureServiceProperties(opts ...func(*ConfigureServicePropertiesConfiguration)) *Request {
	defaults := &ConfigureServicePropertiesConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Configure Service Properties", defaults, opts...)
}

// GlobalSearches manages global searches on an engine with the specified
// configurations.
//
// Parameters:
//
//	opts: A variadic list of functions that modify the GlobalSearchesConfiguration.
//	      These can be used to customize the behavior of the global searches
//	      process.
//
// Returns:
//
//	*Request: A pointer to the Request object configured to manage global
//	          searches on an engine.
func (req *Request) GlobalSearches(opts ...func(*GlobalSearchesConfiguration)) *Request {
	defaults := &GlobalSearchesConfiguration{
		AdpLoggingEnabled:                              false,
		AdpExecutionPersistent:                         false,
		AdpGlobalSearchesUpdateCollisionResolutionMode: "error",
	}
	return newTaskRequest("Global Searches", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// ManageTaggers configures taggers for an application using the given
// configurations.
//
// Parameters:
//   opts: A variadic list of functions that modify the ManageTaggersConfiguration.
//         These can be used to customize the behavior of the manage taggers
//         process.
//
// Returns:
//   *Request: A pointer to the Request object configured to configure taggers
//             for an application.
/******  e716a517-2989-47d1-b7e7-09907ba168e9  *******/
func (req *Request) ManageTaggers(opts ...func(*ManageTaggersConfiguration)) *Request {
	defaults := &ManageTaggersConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Manage Taggers", defaults, opts...)
}

/*************  ✨ Codeium Command ⭐  *************/
// CreateApplication initializes a request to create an application with the specified
// configurations.
//
// Parameters:
//   opts: A variadic list of functions that modify the CreateApplicationConfiguration.
//         These can be used to customize the behavior of the create application process.
//
// Returns:
//   *Request: A pointer to the Request object configured to create an application.
/******  8bdfbd08-ef33-4063-bfac-c181b96ce643  *******/
func (req *Request) CreateApplication(opts ...func(*CreateApplicationConfiguration)) *Request {
	defaults := &CreateApplicationConfiguration{
		AdpLoggingEnabled:      false,
		AdpExecutionPersistent: false,
	}
	return newTaskRequest("Create Application", defaults, opts...)
}
