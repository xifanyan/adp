package adp

import (
	"encoding/json"
)

/*
provides helper functions for interacting with the ADP API, it parses ExecutionMetaData especially for tasks which doublequotes the JSON output.
*/

func (c *Client) ListEntities(opts ...func(*ListEntitiesConfiguration)) ([]Entity, error) {
	var err error
	var resp *Response

	req := NewRequest().ListEntities(opts...)

	if resp, err = c.Send(req); err != nil {
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

func (c *Client) QueryEngine(opts ...func(*QueryEngineConfiguration)) (QueryEngineResult, error) {
	var err error
	var resp *Response

	req := NewRequest().QueryEngine(opts...)
	if resp, err = c.Send(req); err != nil {
		return QueryEngineResult{}, err
	}

	var res QueryEngineResult

	if err = json.Unmarshal(resp.ExecutionMetaData, &res); err != nil {
		return QueryEngineResult{}, err
	}

	return res, err
}

func (c *Client) TaxonomyStatistic(opts ...func(*TaxonomyStatisticConfiguration)) ([]Taxonomy, error) {
	var err error
	var resp *Response

	req := NewRequest().TaxonomyStatistic(opts...)
	if resp, err = c.Send(req); err != nil {
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

func (c *Client) EngineJobMonitoring(opts ...func(*EngineJobMonitoringConfiguration)) ([]EngineJob, error) {
	req := NewRequest().EngineJobMonitoring(opts...)

	resp, err := c.Send(req)
	if err != nil {
		panic(err)
	}

	meta := EngineJobMonitoringExecutionMetaData{}
	if err = json.Unmarshal(resp.ExecutionMetaData, &meta); err != nil {
		return nil, err
	}

	js := string(meta.AdpEngineJobMonitoringJSONOutput)
	UnquoteJSONOutput(&js)

	var res EngineJobMonitoringResult

	if err = json.Unmarshal([]byte(js), &res); err != nil {
		return nil, err
	}

	return res.EngineJobs, err
}
