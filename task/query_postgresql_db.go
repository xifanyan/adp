package task

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/xifanyan/adp/client"
)

// Paths for secured PostgreSQL connection.
const (
	RootCertPath   = "S:/Projects/ProcessControl/security/postgres/root.billing.crt"
	ClientKeyPath  = "S:/Projects/ProcessControl/security/postgres/client.billing.key"
	ClientCertPath = "S:/Projects/ProcessControl/security/postgres/client.billing.crt"
)

// QueryPostgresqlDBConfiguration ...
type QueryPostgresqlDBConfiguration struct {
	AdpProgressTaskTimeout                 int    `json:"adp_progressTaskTimeout,omitempty"`
	AdpQpgdbDbTrustCerts                   string `json:"adp_qpgdb_dbTrustCerts,omitempty"`
	AdpQpgdbOutputJSON                     string `json:"adp_qpgdb_outputJson,omitempty"`
	AdpQpgdbDbUser                         string `json:"adp_qpgdb_dbUser,omitempty"`
	AdpQpgdbJSONFile                       any    `json:"adp_qpgdb_jsonFile,omitempty"`
	AdpTaskActive                          bool   `json:"adp_taskActive,omitempty"`
	AdpExecutionPersistent                 bool   `json:"adp_executionPersistent"`
	AdpQpgdbDbPassword                     string `json:"adp_qpgdb_dbPassword,omitempty"`
	AdpAbortWfOnFailure                    bool   `json:"adp_abortWfOnFailure,omitempty"`
	AdpQpgdbSQLQuery                       string `json:"adp_qpgdb_sqlQuery,omitempty"`
	AdpQpgdbDbConnectionURL                string `json:"adp_qpgdb_dbConnectionUrl,omitempty"`
	AdpCleanUpHistory                      bool   `json:"adp_cleanUpHistory,omitempty"`
	AdpQpgdbDbConnectionPoolClientKeyPath  string `json:"adp_qpgdb_dbConnectionPoolClientKeyPath,omitempty"`
	AdpLoggingEnabled                      bool   `json:"adp_loggingEnabled"`
	AdpQpgdbOutputFilename                 string `json:"adp_qpgdb_outputFilename,omitempty"`
	AdpQpgdbDbConnectionPoolRootCertPath   string `json:"adp_qpgdb_dbConnectionPoolRootCertPath,omitempty"`
	AdpTaskTimeout                         int    `json:"adp_taskTimeout,omitempty"`
	AdpQpgdbDbConnectionPoolClientCertPath string `json:"adp_qpgdb_dbConnectionPoolClientCertPath,omitempty"`
	AdpQpgdbDbConnectionPoolName           string `json:"adp_qpgdb_dbConnectionPoolName,omitempty"`
	AdpQpgdbJSONResultSizeLimitMB          string `json:"adp_qpgdb_jsonResultSizeLimitMB,omitempty"`
}

// NewQueryPostgresqlDBTaskRequest ...
func NewQueryPostgresqlDBTaskRequest(opts ...func(*QueryPostgresqlDBConfiguration)) *Request {

	cfg := &QueryPostgresqlDBConfiguration{}
	for _, opt := range opts {
		opt(cfg)
	}

	log.Debug().Msgf("cfg: %+v", cfg)

	return &Request{
		TaskType:          "Query Postgresql DB",
		TaskDescription:   "Queries postgresql databases.",
		TaskConfiguration: cfg,
		TaskDisplayName:   "Query Postgresql DB",
	}
}

func WithQueryPostgresqlDBLoggingEnabled(b bool) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpLoggingEnabled = b
	}
}

func WithQueryPostgresqlDBExecutionPersistent(b bool) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpExecutionPersistent = b
	}
}

func WithQueryPostgresqlDBConnectionPoolClientKeyPath(s string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbConnectionPoolClientKeyPath = s
	}
}

func WithQueryPostgresqlDBConnectionPoolRootCertPath(s string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbConnectionPoolRootCertPath = s
	}
}

func WithQueryPostgresqlDBConnectionPoolClientCertPath(s string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbConnectionPoolClientCertPath = s
	}
}

func WithQueryPostgresqlDBDbUser(s string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbUser = s
	}
}

func WithQueryPostgresqlDBDbPassword(s string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbPassword = s
	}
}

func WithQueryPostgresqlDBSQLQuery(s string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbSQLQuery = s
	}
}

func WithQueryPostgresqlDBDbConnectionURL(s string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbConnectionURL = s
	}
}

func WithQueryPostgresqlDBJSONResultSizeLimitMB(s string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbJSONResultSizeLimitMB = s
	}
}

// QueryPostgressqlDBExecutionMetaData ...
type QueryPostgressqlDBExecutionMetaData struct {
	AdpQueryDBOutputJSON json.RawMessage `json:"adp_querydb_outputjson"`
}

type QueryPostgresssqlDBTask struct {
	*Task
}

func NewQueryPostgresqlDBTask(client *client.Client, async bool, opts ...func(*QueryPostgresqlDBConfiguration)) *QueryPostgresssqlDBTask {
	return &QueryPostgresssqlDBTask{
		&Task{
			client:       client,
			req:          NewQueryPostgresqlDBTaskRequest(opts...),
			asynchronous: async,
		},
	}
}

func (t *QueryPostgresssqlDBTask) OutputToString() (string, error) {
	taskResp, err := t.Execute()
	if err != nil {
		return "", err
	}

	meta := &QueryPostgressqlDBExecutionMetaData{}
	if err := json.Unmarshal(taskResp.ExecutionMetaData, meta); err != nil {
		return "", fmt.Errorf("parse QueryPostgressqlDBExecutionMetaData %w", err)
	}

	output := string(meta.AdpQueryDBOutputJSON)
	unquoteJSONOutput(&output)

	return output, nil
}
