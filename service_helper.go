package adp

import (
	"encoding/json"
	"fmt"
)

func (svc *Service) ListWorkspaces() ([]Entity, error) {
	return svc.ListEntities(WithListEntitiesType("workspace"))
}

func (svc *Service) ListWorkspacesByUser(user string) ([]Entity, error) {
	return svc.ListEntities(
		WithListEntitiesType("workspace"),
		WithListEntitiesUserHasAccess(user),
	)
}

func (svc *Service) ListDocumentHolds() ([]Entity, error) {
	return svc.ListEntities(WithListEntitiesType("documentHold"))
}

func (svc *Service) ListAxcelerates() ([]Entity, error) {
	return svc.ListEntities(WithListEntitiesType("axcelerate"))
}

func (svc *Service) ListDocumentHoldsByUser(user string) ([]Entity, error) {
	return svc.ListEntities(
		WithListEntitiesType("documentHold"),
		WithListEntitiesUserHasAccess(user),
	)
}

func (svc *Service) ListAxceleratesByUser(user string) ([]Entity, error) {
	return svc.ListEntities(
		WithListEntitiesType("axcelerate"),
		WithListEntitiesUserHasAccess(user),
	)
}

func (svc *Service) ListApplicationsByUser(user string) ([]Entity, error) {
	documentHolds, err := svc.ListDocumentHoldsByUser(user)
	if err != nil {
		return nil, err
	}

	axcelerates, err := svc.ListAxceleratesByUser(user)
	if err != nil {
		return nil, err
	}

	return append(documentHolds, axcelerates...), nil
}

func (svc *Service) ListAvailableTemplates(entityType string, user string) ([]Entity, error) {
	var availableTemplates []Entity

	// list entities with specific type that user has access to
	entities, err := svc.ListEntities(
		WithListEntitiesType(entityType),
		WithListEntitiesUserHasAccess(user),
	)
	if err != nil {
		return nil, err
	}

	// keep track of unique Entity IDs to prevent the situation that the entity user has access
	// to is also a global template.
	var uniqueIDs map[string]struct{} = make(map[string]struct{})

	for _, entity := range entities {
		if _, ok := uniqueIDs[entity.ID]; !ok {
			uniqueIDs[entity.ID] = struct{}{}
			availableTemplates = append(availableTemplates, entity)
		}
	}

	// list all entities with specific type
	entities, err = svc.ListEntities(
		WithListEntitiesType(entityType),
	)
	if err != nil {
		return nil, err
	}

	// filter the entities to find the global templates
	for _, entity := range entities {
		if entity.GlobalTemplateFlag {
			if _, ok := uniqueIDs[entity.ID]; !ok {
				uniqueIDs[entity.ID] = struct{}{}
				availableTemplates = append(availableTemplates, entity)
			}
		}
	}

	return availableTemplates, nil
}

func (svc *Service) ListDatasources() ([]Entity, error) {
	return svc.ListEntities(WithListEntitiesType("dataSource"))
}

func (svc *Service) ListDatasourcesByUser(user string) ([]Entity, error) {
	return svc.ListEntities(
		WithListEntitiesType("dataSource"),
		WithListEntitiesUserHasAccess(user),
	)
}

func (svc *Service) ListHosts() ([]Entity, error) {
	return svc.ListEntities(WithListEntitiesType("host"))
}

func (svc *Service) ListRunningDatasources() ([]Entity, error) {
	return svc.ListEntities(
		WithListEntitiesType("dataSource"),
		WithListEntitiesStatus("Running"),
	)
}

func (svc *Service) ListStartingDatasources() ([]Entity, error) {
	return svc.ListEntities(
		WithListEntitiesType("dataSource"),
		WithListEntitiesStatus("Starting"),
	)
}

func (svc *Service) GetEntityByID(id string) ([]Entity, error) {
	return svc.ListEntities(WithListEntitiesID(id))
}

func (svc *Service) ListEntitiesByRelatedEntity(typ string, id string) ([]Entity, error) {
	return svc.ListEntities(
		WithListEntitiesType(typ),
		WithListEntitiesRelatedEntity(id),
	)
}

func (svc *Service) ListEntitiesByRelatedEntityWithStatus(typ string, id string, status string) ([]Entity, error) {
	return svc.ListEntities(
		WithListEntitiesType(typ),
		WithListEntitiesRelatedEntity(id),
		WithListEntitiesStatus(status),
	)
}

func (svc *Service) ListSingleMindServersByApplication(appID string) ([]Entity, error) {
	return svc.ListEntitiesByRelatedEntity("singleMindServer", appID)
}

func (svc *Service) ListRunningDatasourcesByEngine(engineID string) ([]Entity, error) {
	return svc.ListEntitiesByRelatedEntityWithStatus("dataSource", engineID, "Running")
}

func (svc *Service) ListStartingDatasourcesByEngine(engineID string) ([]Entity, error) {
	return svc.ListEntitiesByRelatedEntityWithStatus("dataSource", engineID, "Starting")
}

func (svc *Service) GetDocumentCountsByEngine(engineID string) (QueryEngineResult, error) {
	return svc.QueryEngine(
		WithQueryEngineEngineName(engineID),
	)
}

func (svc *Service) GetTaxonomyByApplicationID(taxonomy string, appID string) ([]Taxonomy, error) {
	outputTaxonomies := []OutputTaxonomiesArg{
		{
			Taxonomy:                  taxonomy,
			Mode:                      "Category counts",
			MaximumNumberOfCategories: MAX_NUMBER_OF_CATEGORIES,
		},
	}

	res, err := svc.TaxonomyStatistic(
		WithTaxonomyStatisticApplicationIdentifier(appID),
		WithTaxonomyStatisticOutputTaxonomies(outputTaxonomies),
	)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc *Service) GetCustodiansByApplicationID(appID string) ([]Category, error) {
	custodians, err := svc.GetTaxonomyByApplicationID("rm_custodian", appID)
	if err != nil {
		return nil, err
	}

	if len(custodians) == 0 {
		return []Category{}, err
	}

	return custodians[0].Category, nil
}

func (svc *Service) DropTemplate(identifier string) error {
	action := ActionOnEntityArg{
		Identifier: identifier,
		Action:     "Drop Template",
		Value:      "true",
	}

	req := NewRequest().ConfigureServiceProperties(
		WithConfigureServicePropertiesActionsOnEntity([]ActionOnEntityArg{action}),
	)

	_, err := svc.ADPClient.Send(req)
	return err
}

type FieldProperty struct {
	Active            bool
	TextType          string
	DisplayName       string
	ShortDisplayName  string
	Storage           string
	Style             string
	Simple            bool
	Extended          bool
	Table             string
	StaticTextIfEmpty bool
	FullView          bool
	Link              bool
	Associated        bool
	GroupName         string
	HighLight         bool
	Summarize         bool
	Size              float64
}

func (svc *Service) GetFieldProperties(identifier string) (map[string]FieldProperty, error) {
	fieldProperties := make(map[string]FieldProperty)

	cfg := ConfigurationArg{
		ConfigurationID: identifier,        // ID of the configuration
		NameValueList:   "fieldProperties", // Name ofthe components to be retrieved
	}

	res, err := svc.ReadConfiguration(
		WithReadConfigurationConfigsToRead([]ConfigurationArg{cfg}),
	)
	if err != nil {
		return fieldProperties, err
	}

	cells := res[identifier].Global.Static.Parameters[0].Cells
	for _, cell := range cells {
		key := cell[1].Value.(string)
		fieldProperties[key] = FieldProperty{
			Active:            cell[0].Value.(bool),
			TextType:          key,
			DisplayName:       cell[2].Value.(string),
			ShortDisplayName:  cell[3].Value.(string),
			Storage:           cell[4].Value.(string),
			Style:             cell[5].Value.(string),
			Simple:            cell[6].Value.(bool),
			Extended:          cell[7].Value.(bool),
			Table:             cell[8].Value.(string),
			StaticTextIfEmpty: cell[9].Value.(bool),
			FullView:          cell[10].Value.(bool),
			Link:              cell[11].Value.(bool),
			Associated:        cell[12].Value.(bool),
			GroupName:         cell[13].Value.(string),
			HighLight:         cell[14].Value.(bool),
			Summarize:         cell[15].Value.(bool),
			Size:              cell[16].Value.(float64),
		}
	}

	return fieldProperties, err
}

type IndexConfigurationTable struct {
	TextType           string
	Type               string
	StructuredView     bool
	CategoryDelimiters string
}

func (svc *Service) GetIndexConfigurationTable(identifier string) (map[string]IndexConfigurationTable, error) {
	indexConfigurationTable := make(map[string]IndexConfigurationTable)

	cfg := ConfigurationArg{
		ConfigurationID: identifier,                // ID of the configuration
		NameValueList:   "indexConfigurationTable", // Name ofthe components to be retrieved
	}

	res, err := svc.ReadConfiguration(
		WithReadConfigurationConfigsToRead([]ConfigurationArg{cfg}),
	)

	if err != nil {
		return indexConfigurationTable, err
	}

	cells := res[identifier].Global.Static.Parameters[0].Cells
	for _, cell := range cells {
		key := cell[0].Value.(string)
		indexConfigurationTable[key] = IndexConfigurationTable{
			TextType:           key,
			Type:               cell[2].Value.(string),
			StructuredView:     cell[12].Value.(bool),
			CategoryDelimiters: cell[13].Value.(string),
		}
	}

	return indexConfigurationTable, err
}

func (svc *Service) GetCategories(application string, taxonomy string) ([]Taxonomy, error) {
	return svc.TaxonomyStatistic(
		WithTaxonomyStatisticApplicationIdentifier(application),
		WithTaxonomyStatisticOutputTaxonomies([]OutputTaxonomiesArg{
			{
				Taxonomy:                  taxonomy,
				Mode:                      "Category counts",
				MaximumNumberOfCategories: MAX_NUMBER_OF_CATEGORIES,
			},
		}),
	)
}

func (svc *Service) CreateOrUpdateCategory(application string, taxonomyName string, categoryID string, categoryName string) (ManageTaxonomyResult, error) {
	var res ManageTaxonomyResult
	var err error

	input := []ManageTaxonomyTaxonomyInput{
		{
			TaxonomyName: taxonomyName,
			Categories: []ManageTaxonomyCategoryInput{
				{
					ID:   categoryID,
					Mode: "createOrUpdate",
					Name: categoryName,
				},
			},
		},
	}

	js, err := json.Marshal(input)
	if err != nil {
		return res, err
	}

	res, err = svc.ManageTaxonomy(
		WithManageTaxonomyApplicationName(application),
		WithManageTaxonomyTaxonomiesJSON(string(js)),
	)

	return res, err
}

// StartApplicationAsync starts an application asynchronously using the provided
// application identifier. It sends the request to the ADP client and returns
// the execution ID of the task.
//
// Parameters:
//   - appID: Identifier of the application to start.
//
// Returns:
//   - string: The execution ID of the asynchronous task.
//   - error: An error, if any occurs during the request or task execution.

func (svc *Service) StartApplicationAsync(appID string) (string, error) {
	req := NewRequest().StartApplication(
		WithStartApplicationApplicationIdentifier(appID),
	)

	executionID, err := svc.ADPClient.SendAsync(req)
	if err != nil {
		return "", err
	}

	return executionID, nil
}

// GetAllUsersAndGroups retrieves all users and groups from the ADP system.
//
// Parameters:
//
//   - None
//
// Returns:
//
//   - map[string]User: A map of user identifiers to User objects.
//   - map[string]Group: A map of group identifiers to Group objects.
//   - error: An error, if any occurs during the request or task execution.
func (svc *Service) GetAllUsersAndGroups() (map[string]User, map[string]Group, error) {

	res, err := svc.ManageUsersAndGroups()
	if err != nil {
		return nil, nil, err
	}

	users := make(map[string]User)
	groups := make(map[string]Group)

	for k, v := range res.UsersAndGroups.Users {
		users[k] = v
	}

	for k, v := range res.UsersAndGroups.Groups {
		groups[k] = v
	}

	return users, groups, nil
}

// GetUsersByID retrieves all users with the specified ID from the ADP system.
//
// Parameters:
//
//   - id string: The identifier of the user to retrieve.
//
// Returns:
//
//   - map[string]User: A map of user identifiers to User objects with the specified ID.
//   - error: An error, if any occurs during the request or task execution.
func (svc *Service) GetUsersByID(id string) (map[string]User, error) {
	res, err := svc.ManageUsersAndGroups(
		WithManageUsersAndGroupsGroupUserIdsToFilterFor(id),
	)
	if err != nil {
		return nil, err
	}

	return res.UsersAndGroups.Users, nil
}

// GetGroupsByID retrieves all groups with the specified ID from the ADP system.
//
// Parameters:
//
//   - id string: The identifier of the group to retrieve.
//
// Returns:
//
//   - map[string]Group: A map of group identifiers to Group objects with the specified ID.
//   - error: An error, if any occurs during the request or task execution.
func (svc *Service) GetGroupsByID(id string) (map[string]Group, error) {
	res, err := svc.ManageUsersAndGroups(
		WithManageUsersAndGroupsGroupUserIdsToFilterFor(id),
	)
	if err != nil {
		return nil, err
	}

	return res.UsersAndGroups.Groups, nil
}

// GetUsersByGroupID retrieves all users which are members of the specified group from the ADP system.
//
// Parameters:
//
//   - id string: The identifier of the group to retrieve the users from.
//
// Returns:
//
//   - []string: A slice of user identifiers which are members of the specified group.
//   - error: An error, if any occurs during the request or task execution.
func (svc *Service) GetUsersByGroupID(id string) ([]string, error) {
	res, err := svc.ManageUsersAndGroups(
		WithManageUsersAndGroupsGroupUserIdsToFilterFor(id),
		WithManageUsersAndGroupsReturnAllUsersUnderGroup("true"),
	)

	if err != nil {
		return nil, err
	}

	group, ok := res.UsersAndGroups.Groups[id]
	if !ok {
		return nil, fmt.Errorf("group %s not found", id)
	}

	return group.Users, nil
}
