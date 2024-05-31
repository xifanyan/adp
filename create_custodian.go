package adp

type Custodian struct {
	Id          string `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description,omitempty"`
	Address     string `json:"Address,omitempty"`
	Department  string `json:"Department,omitempty"`
	Email       string `json:"E-mail,omitempty"`
}

type CreateCustodianConfiguration struct {
	AdpProgressTaskTimeout                     int         `json:"adp_progressTaskTimeout,omitempty"`
	AdpTaskActive                              bool        `json:"adp_taskActive,omitempty"`
	AdpExecutionPersistent                     bool        `json:"adp_executionPersistent,omitempty"`
	AdpCreateCustodianEngineUser               string      `json:"adp_createCustodian_engineUser,omitempty"`
	AdpAbortWfOnFailure                        bool        `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                          bool        `json:"adp_cleanUpHistory,omitempty"`
	AdpCreateCustodianApplicationType          string      `json:"adp_createCustodian_applicationType,omitempty"`
	AdpCreateCustodianEngineIdentifier         string      `json:"adp_createCustodian_engineIdentifier,omitempty"`
	AdpLoggingEnabled                          bool        `json:"adp_loggingEnabled,omitempty"`
	AdpCreateCustodianCustodians               []Custodian `json:"adp_createCustodian_custodians,omitempty"`
	AdpCreateCustodianCustodianTaxonomy        string      `json:"adp_createCustodian_custodianTaxonomy,omitempty"`
	AdpTaskTimeout                             int         `json:"adp_taskTimeout,omitempty"`
	AdpCreateCustodianApplicationIdentifier    string      `json:"adp_createCustodian_applicationIdentifier,omitempty"`
	AdpCreateCustodianEngineType               string      `json:"adp_createCustodian_engineType,omitempty"`
	AdpCreateCustodianEngineUserPassword       string      `json:"adp_createCustodian_engineUserPassword,omitempty"`
	AdpCreateCustodianUpdateExistingCustodians bool        `json:"adp_createCustodian_updateExistingCustodians,omitempty"`
}

func (cfg *CreateCustodianConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithCreateCustodianLoggingEnabled(enabled bool) func(*CreateCustodianConfiguration) {
	return func(c *CreateCustodianConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithCreateCustodianApplicationIdentifier(id string) func(*CreateCustodianConfiguration) {
	return func(c *CreateCustodianConfiguration) {
		c.AdpCreateCustodianApplicationIdentifier = id
	}
}

func WithCreateCustodianCustodians(custodians []Custodian) func(*CreateCustodianConfiguration) {
	return func(c *CreateCustodianConfiguration) {
		c.AdpCreateCustodianCustodians = custodians
	}
}

func WithCreateCustodianEngineIdentifier(id string) func(*CreateCustodianConfiguration) {
	return func(c *CreateCustodianConfiguration) {
		c.AdpCreateCustodianEngineIdentifier = id
	}
}

func WithCreateCustodianUpdateExistingCustodians(update bool) func(*CreateCustodianConfiguration) {
	return func(c *CreateCustodianConfiguration) {
		c.AdpCreateCustodianUpdateExistingCustodians = update
	}
}
