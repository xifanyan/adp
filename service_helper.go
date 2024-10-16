package adp

func (svc *Service) ListDocumentHolds() ([]Entity, error) {
	return svc.ListEntities(WithListEntitiesType("documentHold"))
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
