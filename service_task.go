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

	log.Debug().Msgf("ExecutionMetaData: %s", string(resp.ExecutionMetaData))

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

	log.Debug().Msgf("ExecutionMetaData: %s", string(resp.ExecutionMetaData))

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

	log.Debug().Msgf("ExecutionMetaData: %s", string(resp.ExecutionMetaData))

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
	resp, err := svc.ADPClient.Send(req)

	log.Debug().Msgf("ExecutionMetaData: %s", string(resp.ExecutionMetaData))

	return err
}

func (svc *Service) CreateDataSource(opts ...func(*CreateDataSourceConfiguration)) error {
	req := NewRequest().CreateDataSource(opts...)
	resp, err := svc.ADPClient.Send(req)

	log.Debug().Msgf("ExecutionMetaData: %s", string(resp.ExecutionMetaData))

	return err
}

func (svc *Service) ConfigureDataSource(opts ...func(*ConfigureDataSourceConfiguration)) error {
	req := NewRequest().ConfigureDataSource(opts...)
	resp, err := svc.ADPClient.Send(req)

	log.Debug().Msgf("ExecutionMetaData: %s", string(resp.ExecutionMetaData))

	return err
}

func (svc *Service) StartDataSource(opts ...func(*StartDataSourceConfiguration)) error {
	req := NewRequest().StartDataSource(opts...)
	_, err := svc.ADPClient.SendAsync(req)

	return err
}

func (svc *Service) ManageUsersAndGroups(opts ...func(*ManageUsersAndGruopsConfiguration)) (ManageUsersAndGroupsResult, error) {
	var err error
	var res ManageUsersAndGroupsResult

	req := NewRequest().ManageUsersAndGroups(opts...)
	resp, err := svc.ADPClient.Send(req)
	if err != nil {
		return res, err
	}

	log.Debug().Msgf("ExecutionMetaData: %s", string(resp.ExecutionMetaData))

	var meta ManageUsersAndGroupsExecutionMetaData
	if err = json.Unmarshal(resp.ExecutionMetaData, &meta); err != nil {
		return res, err
	}

	js := string(meta.AdpManageUsersAndGroupsJSONOutput)
	UnquoteJSONOutput(&js)

	if err = json.Unmarshal([]byte(js), &res); err != nil {
		return res, err
	}

	return res, err
}

func (svc *Service) ConfigureEntity(opts ...func(*ConfigureEntityConfiguration)) error {
	req := NewRequest().ConfigureEntity(opts...)
	resp, err := svc.ADPClient.Send(req)

	log.Debug().Msgf("ExecutionMetaData: %s", string(resp.ExecutionMetaData))

	return err
}

func (svc *Service) GlobalSearches(opts ...func(*GlobalSearchesConfiguration)) ([]GlobalSearch, error) {
	var err error
	var resp *Response
	var res []GlobalSearch = []GlobalSearch{}

	req := NewRequest().GlobalSearches(opts...)
	if resp, err = svc.ADPClient.Send(req); err != nil {
		return res, err
	}

	log.Debug().Msgf("ExecutionMetaData: %s", string(resp.ExecutionMetaData))

	var meta GlobalSearchesExecutionMetaData
	if err = json.Unmarshal(resp.ExecutionMetaData, &meta); err != nil {
		return nil, err
	}

	js := string(meta.AdpGlobalSearchesJSONOutput)
	UnquoteJSONOutput(&js)

	if err = json.Unmarshal([]byte(js), &res); err != nil {
		return nil, err
	}

	return res, err
}

func (svc *Service) ManageTaggers(opts ...func(*ManageTaggersConfiguration)) error {
	var err error
	var resp *Response

	req := NewRequest().ManageTaggers(opts...)
	if resp, err = svc.ADPClient.Send(req); err != nil {
		return err
	}

	log.Debug().Msgf("ExecutionMetaData: %s", string(resp.ExecutionMetaData))

	return nil
}

func (svc *Service) ReadConfiguration(opts ...func(*ReadConfigurationConfiguration)) (ReadConfigurationResult, error) {
	var err error
	var res ReadConfigurationResult

	req := NewRequest().ReadConfiguration(opts...)
	resp, err := svc.ADPClient.Send(req)
	if err != nil {
		return res, err
	}

	log.Debug().Msgf("ExecutionMetaData: %s", string(resp.ExecutionMetaData))

	var meta ReadConfigurationExecutionMetaData
	if err = json.Unmarshal(resp.ExecutionMetaData, &meta); err != nil {
		return res, err
	}

	js := string(meta.AdpEntitiesJSONOutput)
	UnquoteJSONOutput(&js)

	if err = json.Unmarshal([]byte(js), &res); err != nil {
		return nil, err
	}

	return res, err
}
