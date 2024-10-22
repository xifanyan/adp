package adp

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
