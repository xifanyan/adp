package adp

type PublishToReviewConfiguration struct {
	AdpProgressTaskTimeout                               int      `json:"adp_progressTaskTimeout,omitempty"`
	AdpTaskActive                                        bool     `json:"adp_taskActive,omitempty"`
	AdpExecutionPersistent                               bool     `json:"adp_executionPersistent,omitempty"`
	AdpAbortWfOnFailure                                  bool     `json:"adp_abortWfOnFailure,omitempty"`
	AdpLoggingEnabled                                    bool     `json:"adp_loggingEnabled,omitempty"`
	AdpCleanUpHistory                                    bool     `json:"adp_cleanUpHistory,omitempty"`
	AdpTaskTimeout                                       int      `json:"adp_taskTimeout,omitempty"`
	AdpPtrMatterSpecificUrlRegEx                         string   `json:"adp_ptr_matterSpecificUrlRegEx,omitempty"`
	AdpPtrMatterId                                       string   `json:"adp_ptr_matterId,omitempty"`
	AdpPtrWebserviceRequestTimeoutSeconds                int      `json:"adp_ptr_webserviceRequestTimeoutSeconds,omitempty"`
	AdpPtrSearchDetails                                  []string `json:"adp_ptr_searchDetails,omitempty"`
	AdpPtrPublishEngineId                                string   `json:"adp_ptr_publishEngineId,omitempty"`
	AdpPtrStartLearner                                   string   `json:"adp_ptr_startLearner,omitempty"`
	AdpPtrSearchString                                   string   `json:"adp_ptr_searchString,omitempty"`
	AdpPtrWaitForMatterCompletion                        bool     `json:"adp_ptr_waitForMatterCompletion,omitempty"`
	AdpPtrUsedWebserviceUrl                              string   `json:"adp_ptr_usedWebserviceUrl,omitempty"`
	AdpPtrSecondsBetweenNextTryToWaitForMatterCompletion string   `json:"adp_ptr_secondsBetweenNextTryToWaitForMatterCompletion,omitempty"`
	AdpPtrEcaEngine                                      string   `json:"adp_ptr_ecaEngine,omitempty"`
	AdpPtrEcaMasterPort                                  string   `json:"adp_ptr_ecaMasterPort,omitempty"`
	AdpPtrHttpsKeystoreFile                              string   `json:"adp_ptr_httpsKeystoreFile,omitempty"`
	AdpPtrMode                                           string   `json:"adp_ptr_mode,omitempty"`
	AdpPtrWebserviceConnectTimeoutSeconds                int      `json:"adp_ptr_webserviceConnectTimeoutSeconds,omitempty"`
	AdpPtrPublishApplicationId                           string   `json:"adp_ptr_publishApplicationId,omitempty"`
	AdpPtrHttpsPassword                                  string   `json:"adp_ptr_httpsPassword,omitempty"`
	AdpPtrWebservicePassword                             string   `json:"adp_ptr_webservicePassword,omitempty"`
	AdpPtrWebserviceUrl                                  string   `json:"adp_ptr_webserviceUrl,omitempty"`
	AdpPtrPublishApplicationUrl                          string   `json:"adp_ptr_publishApplicationUrl,omitempty"`
	AdpPtrEnforceDeduplication                           bool     `json:"adp_ptr_enforceDeduplication,omitempty"`
	AdpPtrMatterSpecificApplication                      string   `json:"adp_ptr_matterSpecificApplication,omitempty"`
	AdpPtrWebserviceUser                                 string   `json:"adp_ptr_webserviceUser,omitempty"`
	AdpPtrDeduplicate                                    bool     `json:"adp_ptr_deduplicate,omitempty"`
	AdpPtrEcaApplication                                 string   `json:"adp_ptr_ecaApplication,omitempty"`
	AdpPtrHttpsAllowUntrustedHosts                       string   `json:"adp_ptr_httpsAllowUntrustedHosts,omitempty"`
	AdpPtrEcaPublish                                     bool     `json:"adp_ptr_ecaPublish,omitempty"`
	AdpPtrEcaMasterHost                                  string   `json:"adp_ptr_ecaMasterHost,omitempty"`
	AdpPtrHttpsTrustCertificate                          string   `json:"adp_ptr_httpsTrustCertificate,omitempty"`
}

func (cfg *PublishToReviewConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithPublishToReviewLoggingEnabled(enabled bool) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}

func WithPublishToReviewExecutionPersistent(persistent bool) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpExecutionPersistent = persistent
	}
}

func WithPublishToReviewMatterId(matterId string) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpPtrMatterId = matterId
	}
}

func WithPublishToReviewPublishEngineId(engineId string) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpPtrPublishEngineId = engineId
	}
}

func WithPublishToReviewPublishApplicationId(appId string) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpPtrPublishApplicationId = appId
	}
}

func WithPublishToReviewWebserviceUrl(url string) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpPtrWebserviceUrl = url
	}
}

func WithPublishToReviewWebserviceUser(user string) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpPtrWebserviceUser = user
	}
}

func WithPublishToReviewWebservicePassword(password string) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpPtrWebservicePassword = password
	}
}

func WithPublishToReviewEcaApplication(application string) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpPtrEcaApplication = application
	}
}

func WithPublishToReviewEcaEngine(engine string) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpPtrEcaEngine = engine
	}
}

func WithPublishToReviewMode(mode string) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpPtrMode = mode
	}
}

func WithPublishToReviewWaitForMatterCompletion(wait bool) func(*PublishToReviewConfiguration) {
	return func(c *PublishToReviewConfiguration) {
		c.AdpPtrWaitForMatterCompletion = wait
	}
}
