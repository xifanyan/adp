package adp

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Service struct {
	ADPClient *Client
}

func (svc *Service) ListEntities(opts ...func(*ListEntitiesConfiguration)) ([]Entity, error) {
	var err error
	var resp *Response

	req := NewRequest().ListEntities(opts...)

	if resp, err = svc.ADPClient.Send(req); err != nil {
		return nil, err
	}

	var meta ListEntitiesExecutionMetaData

	if err = json.Unmarshal(resp.ExecutionMetaData, &meta); err != nil {
		return nil, err
	}

	js := string(meta.AdpEntitiesJSONOutput)
	UnquoteJSONOutput(&js)

	var res = []Entity{}

	// when no entity was found
	if err = json.Unmarshal([]byte(js), &res); err != nil {
		return res, nil
	}

	return res, err
}

func (svc *Service) QueryEngine(opts ...func(*QueryEngineConfiguration)) (QueryEngineResult, error) {
	var err error
	var resp *Response

	req := NewRequest().QueryEngine(opts...)
	if resp, err = svc.ADPClient.Send(req); err != nil {
		return QueryEngineResult{}, err
	}

	var res QueryEngineResult

	if err = json.Unmarshal(resp.ExecutionMetaData, &res); err != nil {
		log.Error().Err(err).Msgf("failed to unmarshal ExecutionMetaData %s", resp.ExecutionMetaData)
		return QueryEngineResult{}, err
	}

	return res, err
}

func (svc *Service) TaxonomyStatistic(opts ...func(*TaxonomyStatisticConfiguration)) ([]Taxonomy, error) {
	var err error
	var resp *Response

	req := NewRequest().TaxonomyStatistic(opts...)
	if resp, err = svc.ADPClient.Send(req); err != nil {
		return TaxonomyStatisticResult{}, err
	}

	var meta TaxonomyStatisticExecutionMetaData

	if err = json.Unmarshal(resp.ExecutionMetaData, &meta); err != nil {
		return nil, err
	}

	js := string(meta.AdpTaxonomyStatisticsJSONOutput)
	UnquoteJSONOutput(&js)

	var res = TaxonomyStatisticJSONOutput{}

	if err = json.Unmarshal([]byte(js), &res); err != nil {
		return nil, err
	}

	return res.Statistics.Taxonomy, err
}

func (svc *Service) CreateCustodian(opts ...func(*CreateCustodianConfiguration)) error {
	req := NewRequest().CreateCustodian(opts...)

	_, err := svc.ADPClient.Send(req)
	return err
}

func (svc *Service) CreateDataSource(opts ...func(*CreateDataSourceConfiguration)) error {
	req := NewRequest().CreateDataSource(opts...)

	_, err := svc.ADPClient.Send(req)
	return err
}

func (svc *Service) ConfigureDataSource(opts ...func(*ConfigureDataSourceConfiguration)) error {
	req := NewRequest().ConfigureDataSource(opts...)

	_, err := svc.ADPClient.Send(req)
	return err
}

func (svc *Service) StartDataSource(opts ...func(*StartDataSourceConfiguration)) error {
	req := NewRequest().StartDataSource(opts...)

	_, err := svc.ADPClient.SendAsync(req)
	return err
}

func (svc *Service) ManageUsersAndGroups(opts ...func(*ManageUsersAndGruopsConfiguration)) (UsersAndGroups, error) {
	var err error
	var resp *Response

	var usersAndGroups UsersAndGroups

	req := NewRequest().ManageUsersAndGroups(opts...)
	if resp, err = svc.ADPClient.Send(req); err != nil {
		return usersAndGroups, err
	}

	var meta ManageUsersAndGroupsExecutionMetaData

	if err = json.Unmarshal(resp.ExecutionMetaData, &meta); err != nil {
		return usersAndGroups, err
	}

	js := string(meta.AdpManageUsersAndGroupsJSONOutput)
	UnquoteJSONOutput(&js)

	err = json.Unmarshal([]byte(js), &usersAndGroups)

	return usersAndGroups, err
}
