package adp

import (
	"encoding/json"
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

func (cfg *QueryPostgresqlDBConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithQueryPostgresqlDBLoggingEnabled(enabled bool) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithQueryPostgresqlDBExecutionPersistent(enabled bool) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpExecutionPersistent = enabled
	}
}

func WithQueryPostgresqlDBConnectionPoolClientKeyPath(clientKeyPath string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbConnectionPoolClientKeyPath = clientKeyPath
	}
}

func WithQueryPostgresqlDBConnectionPoolRootCertPath(rootCertPath string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbConnectionPoolRootCertPath = rootCertPath
	}
}

func WithQueryPostgresqlDBConnectionPoolClientCertPath(clientCertPath string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbConnectionPoolClientCertPath = clientCertPath
	}
}

func WithQueryPostgresqlDBDbUser(s string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbUser = s
	}
}

func WithQueryPostgresqlDBDbPassword(user string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbPassword = user
	}
}

func WithQueryPostgresqlDBSQLQuery(query string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbSQLQuery = query
	}
}

func WithQueryPostgresqlDBDbConnectionURL(connectionURL string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbDbConnectionURL = connectionURL
	}
}

func WithQueryPostgresqlDBJSONResultSizeLimitMB(sizeLimit string) func(*QueryPostgresqlDBConfiguration) {
	return func(c *QueryPostgresqlDBConfiguration) {
		c.AdpQpgdbJSONResultSizeLimitMB = sizeLimit
	}
}

type QueryPostgressqlDBExecutionMetaData struct {
	AdpQueryDBOutputJSON json.RawMessage `json:"adp_querydb_outputjson"`
}
