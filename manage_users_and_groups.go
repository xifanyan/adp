package adp

import "encoding/json"

type UserDefinition struct {
	Enabled      bool   `json:"Enabled,omitempty"`
	ExternalUser bool   `json:"External user,omitempty"`
	Password     string `json:"Password,omitempty"`
	Remove       bool   `json:"Remove,omitempty"`
	UserName     string `json:"User name,omitempty"`
}

type GroupDefinition struct {
	Enabled   bool   `json:"Enabled,omitempty"`
	GroupName string `json:"Group name,omitempty"`
	Remove    bool   `json:"Remove,omitempty"`
}

type UserToGroup struct {
	Enabled   bool   `json:"Enabled,omitempty"`
	GroupName string `json:"Group name,omitempty"`
	Remove    bool   `json:"Remove,omitempty"`
	UserName  string `json:"User name,omitempty"`
}

type ManageUsersAndGroupsConfiguration struct {
	AdpProgressTaskTimeout                          int                `json:"adp_progressTaskTimeout,omitempty"`
	AdpManageUsersAndGroupsUserDefinition           []UserDefinition   `json:"adp_manageUsersAndGroups_userDefinition,omitempty"`
	AdpTaskActive                                   bool               `json:"adp_taskActive,omitempty"`
	AdpManageUsersAndGroupsGroupUserIdsToFilterFor  string             `json:"adp_manageUsersAndGroups_GroupUserIdsToFilterFor,omitempty"`
	AdpManageUsersAndGroupsGroupDefinition          []GroupDefinition  `json:"adp_manageUsersAndGroups_groupDefinition,omitempty"`
	AdpManageUsersAndGroupsAppIdsToFilterFor        string             `json:"adp_manageUsersAndGroups_AppIdsToFilterFor,omitempty"`
	AdpExecutionPersistent                          bool               `json:"adp_executionPersistent,omitempty"`
	AdpManageUsersAndGroupsAssignmentUserToGroup    []UserToGroup      `json:"adp_manageUsersAndGroups_assignmentUserToGroup,omitempty"`
	AdpManageUsersAndGroupsAddApplicationRoles      []ApplicationRoles `json:"adp_manageUsersAndGroups_addApplicationRoles,omitempty"`
	AdpAbortWfOnFailure                             bool               `json:"adp_abortWfOnFailure,omitempty"`
	AdpCleanUpHistory                               bool               `json:"adp_cleanUpHistory,omitempty"`
	AdpManageUsersAndGroupsFile                     string             `json:"adp_manageUsersAndGroups_file,omitempty"`
	AdpLoggingEnabled                               bool               `json:"adp_loggingEnabled,omitempty"`
	AdpManageUsersAndGroupsOutputFilename           string             `json:"adp_manageUsersAndGroups_outputFilename,omitempty"`
	AdpManageUsersAndGroupsOutputJSON               string             `json:"adp_manageUsersAndGroups_outputJson,omitempty"`
	AdpTaskTimeout                                  int                `json:"adp_taskTimeout,omitempty"`
	AdpManageUsersAndGroupsReturnAllUsersUnderGroup string             `json:"adp_manageUsersAndGroups_ReturnAllUsersUnderGroup,omitempty"`
}

func (cfg *ManageUsersAndGroupsConfiguration) enforcePersistentExecution() {
	cfg.AdpExecutionPersistent = true
}

func WithManageUsersAndGroupsLoggingEnabled(enabled bool) func(*ManageUsersAndGroupsConfiguration) {
	return func(c *ManageUsersAndGroupsConfiguration) {
		c.AdpLoggingEnabled = enabled
	}
}
func WithManageUsersAndGroupsAppIdsToFilterFor(appIds string) func(*ManageUsersAndGroupsConfiguration) {
	return func(c *ManageUsersAndGroupsConfiguration) {
		c.AdpManageUsersAndGroupsAppIdsToFilterFor = appIds
	}
}

func WithManageUsersAndGroupsUserDefinition(users []UserDefinition) func(*ManageUsersAndGroupsConfiguration) {
	return func(c *ManageUsersAndGroupsConfiguration) {
		c.AdpManageUsersAndGroupsUserDefinition = users
	}
}

func WithManageUsersAndGroupsGroupDefinition(groups []GroupDefinition) func(*ManageUsersAndGroupsConfiguration) {
	return func(c *ManageUsersAndGroupsConfiguration) {
		c.AdpManageUsersAndGroupsGroupDefinition = groups
	}
}

func WithManageUsersAndGroupsAssignmentUserToGroup(userToGroup []UserToGroup) func(*ManageUsersAndGroupsConfiguration) {
	return func(cfg *ManageUsersAndGroupsConfiguration) {
		cfg.AdpManageUsersAndGroupsAssignmentUserToGroup = userToGroup
	}
}

func WithManageUsersAndGroupsAddApplicationRoles(applicationRoles []ApplicationRoles) func(*ManageUsersAndGroupsConfiguration) {
	return func(cfg *ManageUsersAndGroupsConfiguration) {
		cfg.AdpManageUsersAndGroupsAddApplicationRoles = applicationRoles
	}
}

func WithManageUsersAndGroupsReturnAllUsersUnderGroup(returnAllUsersUnderGroup string) func(*ManageUsersAndGroupsConfiguration) {
	return func(c *ManageUsersAndGroupsConfiguration) {
		c.AdpManageUsersAndGroupsReturnAllUsersUnderGroup = returnAllUsersUnderGroup
	}
}

func WithManageUsersAndGroupsGroupUserIdsToFilterFor(groupUserIds string) func(*ManageUsersAndGroupsConfiguration) {
	return func(c *ManageUsersAndGroupsConfiguration) {
		c.AdpManageUsersAndGroupsGroupUserIdsToFilterFor = groupUserIds
	}
}

type ManageUsersAndGroupsExecutionMetaData struct {
	AdpManageUsersAndGroupsJSONOutput json.RawMessage `json:"adp_manageUsersAndGroups_json_output,omitempty"`
}

type UserGroupCommon struct {
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Existent    bool   `json:"existent"`
}

type Group struct {
	UserGroupCommon
	Users []string `json:"users"`
}

type User struct {
	UserGroupCommon
	External     bool   `json:"external"`
	EmailAddress string `json:"emailAddress"`
}

type UsersAndGroups struct {
	Groups map[string]Group `json:"groups"`
	Users  map[string]User  `json:"users"`
}

type ApplicationRoles struct {
	GroupOrUserName       string `json:"Group-/User- name"`
	Enabled               bool   `json:"Enabled"`
	ApplicationIdentifier string `json:"Application identifier"`
	Roles                 string `json:"Role(s)"`
}

type ManageUsersAndGroupsResult struct {
	UsersAndGroups
	Applications map[string]struct {
		UsersAndGroups
	} `json:"applications"`
}
