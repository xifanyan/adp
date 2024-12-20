package adp

import (
	"encoding/json"
)

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
